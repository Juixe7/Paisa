package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv            string
	LogLevel          string
	DatabaseURL       string
	RedisURL          string
	RabbitMQURL       string
	JWTSecret         string
	JWTRefreshSecret  string
	S3Bucket          string
	S3Endpoint        string
	AWSAccessKeyID    string
	AWSSecretKey      string
	AIServiceURL      string
	RazorpayKeyID     string
	RazorpaySecret    string
}

func Load() (*Config, error) {
	cfg := &Config{
		AppEnv:            os.Getenv("APP_ENV"),
		LogLevel:          os.Getenv("LOG_LEVEL"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		RedisURL:          os.Getenv("REDIS_URL"),
		RabbitMQURL:       os.Getenv("RABBITMQ_URL"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTRefreshSecret:  os.Getenv("JWT_REFRESH_SECRET"),
		S3Bucket:          os.Getenv("S3_BUCKET"),
		S3Endpoint:        os.Getenv("S3_ENDPOINT"),
		AWSAccessKeyID:    os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretKey:      os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AIServiceURL:      os.Getenv("AI_SERVICE_URL"),
		RazorpayKeyID:     os.Getenv("RAZORPAY_KEY_ID"),
		RazorpaySecret:    os.Getenv("RAZORPAY_KEY_SECRET"),
	}

	// Verify required variables
	required := map[string]string{
		"DATABASE_URL":      cfg.DatabaseURL,
		"REDIS_URL":         cfg.RedisURL,
		"RABBITMQ_URL":      cfg.RabbitMQURL,
		"JWT_SECRET":        cfg.JWTSecret,
		"JWT_REFRESH_SECRET": cfg.JWTRefreshSecret,
	}

	for key, val := range required {
		if val == "" {
			return nil, fmt.Errorf("missing required environment variable: %s", key)
		}
	}

	if cfg.AppEnv == "" {
		cfg.AppEnv = "local"
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	return cfg, nil
}
