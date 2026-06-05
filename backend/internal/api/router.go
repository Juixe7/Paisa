package api

import (
	"time"

	"github.com/Juixe7/Paisa/backend/internal/auth"
	"github.com/Juixe7/Paisa/backend/internal/categorisation"
	"github.com/Juixe7/Paisa/backend/internal/analytics"
	"github.com/Juixe7/Paisa/backend/internal/notification"
	"github.com/Juixe7/Paisa/backend/internal/goals"
	"github.com/Juixe7/Paisa/backend/internal/transaction"
	
	"github.com/gin-gonic/gin"
	redisclient "github.com/redis/go-redis/v9"
)

type RouterConfig struct {
	JWTSecret        []byte
	RedisClient      *redisclient.Client
	AuthHandler      *auth.Handler
	TxHandler        *transaction.Handler
	CatHandler       *categorisation.Handler
	AnalyticsHandler *analytics.Handler
	NotifHandler     *notification.Handler
	GoalsHandler     *goals.Handler
}

func SetupRouter(cfg RouterConfig) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ErrorHandler())

	// Standard health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	// Public Auth endpoints
	authLimiter := RateLimiter(cfg.RedisClient, 10, 1*time.Minute)
	authGroup := r.Group("/auth")
	authGroup.Use(authLimiter)
	{
		authGroup.POST("/register", cfg.AuthHandler.Register)
		authGroup.POST("/login", cfg.AuthHandler.Login)
		authGroup.POST("/refresh", cfg.AuthHandler.Refresh)
		authGroup.POST("/logout", cfg.AuthHandler.Logout)
	}

	// Protected Endpoints
	apiLimiter := RateLimiter(cfg.RedisClient, 100, 1*time.Minute)
	authMiddleware := AuthMiddleware(cfg.JWTSecret)

	apiGroup := r.Group("/api/v1")
	apiGroup.Use(authMiddleware)
	apiGroup.Use(apiLimiter)
	{
		// User profile
		apiGroup.GET("/users/me", cfg.AuthHandler.GetMe)
		apiGroup.PATCH("/users/me", cfg.AuthHandler.UpdateMe)

		// Transactions
		apiGroup.POST("/transactions", cfg.TxHandler.CreateManual)
		apiGroup.GET("/transactions", cfg.TxHandler.List)
		apiGroup.GET("/transactions/:id", cfg.TxHandler.Detail)
		apiGroup.POST("/transactions/sms", cfg.TxHandler.IngestSMS)
		apiGroup.POST("/transactions/intent", cfg.TxHandler.CreateIntent)
		apiGroup.PATCH("/transactions/intent/:id/confirm", cfg.TxHandler.ConfirmIntent)

		// Categorisation
		apiGroup.PATCH("/transactions/:id/category", cfg.CatHandler.CorrectCategory)
		apiGroup.GET("/categories", cfg.CatHandler.GetCategories)

		// Analytics
		apiGroup.GET("/analytics/summary", cfg.AnalyticsHandler.GetSummary)
		apiGroup.GET("/analytics/budgets", cfg.AnalyticsHandler.GetBudgets)
		apiGroup.PATCH("/analytics/budgets/:category_id", cfg.AnalyticsHandler.UpdateBudget)

		// Notifications
		apiGroup.GET("/notifications", cfg.NotifHandler.GetNotifications)
		apiGroup.PATCH("/notifications/:id/read", cfg.NotifHandler.MarkRead)

		// Goals
		apiGroup.POST("/goals", cfg.GoalsHandler.Create)
		apiGroup.GET("/goals", cfg.GoalsHandler.List)
		apiGroup.POST("/goals/:id/contribute", cfg.GoalsHandler.Contribute)
	}

	return r
}
