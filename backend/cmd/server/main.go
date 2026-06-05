package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avadhut/paisa/backend/internal/api"
	"github.com/avadhut/paisa/backend/internal/auth"
	"github.com/avadhut/paisa/backend/internal/categorisation"
	"github.com/avadhut/paisa/backend/internal/analytics"
	"github.com/avadhut/paisa/backend/internal/notification"
	"github.com/avadhut/paisa/backend/internal/goals"
	"github.com/avadhut/paisa/backend/internal/transaction"
	
	"github.com/avadhut/paisa/backend/internal/shared/config"
	"github.com/avadhut/paisa/backend/internal/shared/db"
	"github.com/avadhut/paisa/backend/internal/shared/logger"
	"github.com/avadhut/paisa/backend/internal/shared/redis"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	// 1. Initialize Logger
	logger.Init("info")
	log := logger.Log

	// 2. Load Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load configuration")
	}
	logger.Init(cfg.LogLevel)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 3. Connect to Database (Postgres)
	dbClient, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	defer dbClient.Close()

	// 4. Connect to Redis
	redisClient, err := redis.Connect(ctx, cfg.RedisURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to redis")
	}
	defer redisClient.Close()

	// 5. Connect to RabbitMQ (Optional local fail-safe for scaffolding)
	var rabbitConn *amqp091.Connection
	var rabbitChan *amqp091.Channel
	
	rabbitConn, err = amqp091.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Warn().Err(err).Msg("failed to connect to RabbitMQ, running in degraded mode")
	} else {
		defer rabbitConn.Close()
		rabbitChan, err = rabbitConn.Channel()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to open RabbitMQ channel")
		}
		defer rabbitChan.Close()

		// Declare exchanges (Rule 6.1)
		err = rabbitChan.ExchangeDeclare("transactions.exchange", "topic", true, false, false, false, nil)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to declare transactions.exchange")
		}
	}

	// 6. Initialize Repositories and Services
	authRepo := auth.NewRepository(dbClient.Pool)
	authSvc := auth.NewService(authRepo, cfg.JWTSecret, cfg.JWTRefreshSecret)
	authHandler := auth.NewHandler(authSvc)

	// Mock publisher if rabbitmq isn't available
	var txPublisher transaction.Publisher
	if rabbitChan != nil {
		txPublisher = transaction.NewPublisher(rabbitChan)
	} else {
		// Mock publisher
		txPublisher = &mockPublisher{}
	}

	txRepo := transaction.NewRepository(dbClient.Pool)
	txSvc := transaction.NewService(txRepo, txPublisher)
	txHandler := transaction.NewHandler(txSvc)

	catRepo := categorisation.NewRepository(dbClient.Pool)
	catSvc := categorisation.NewService(catRepo, redisClient.RDB)
	catHandler := categorisation.NewHandler(catSvc)

	analyticsRepo := analytics.NewRepository(dbClient.Pool)
	analyticsSvc := analytics.NewService(analyticsRepo)
	analyticsHandler := analytics.NewHandler(analyticsSvc)

	notifRepo := notification.NewRepository(dbClient.Pool)
	notifSvc := notification.NewService(notifRepo)
	notifHandler := notification.NewHandler(notifSvc)

	goalsRepo := goals.NewRepository(dbClient.Pool)
	goalsSvc := goals.NewService(goalsRepo)
	goalsHandler := goals.NewHandler(goalsSvc)

	// 7. Setup Router
	router := api.SetupRouter(api.RouterConfig{
		JWTSecret:        []byte(cfg.JWTSecret),
		RedisClient:      redisClient.RDB,
		AuthHandler:      authHandler,
		TxHandler:        txHandler,
		CatHandler:       catHandler,
		AnalyticsHandler: analyticsHandler,
		NotifHandler:     notifHandler,
		GoalsHandler:     goalsHandler,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 8. Graceful Shutdown
	go func() {
		log.Info().Msgf("Server listening on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to listen and serve")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting gracefully")
}

// mockPublisher used when RabbitMQ is not available locally
type mockPublisher struct{}

func (m *mockPublisher) PublishTransactionCreated(ctx context.Context, tx *transaction.Transaction) error {
	logger.Log.Info().Msgf("[Mock RabbitMQ] Published transaction.created for ID: %s", tx.ID)
	return nil
}
