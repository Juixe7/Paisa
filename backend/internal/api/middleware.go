package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	redisclient "github.com/redis/go-redis/v9"
)

func AuthMiddleware(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required", "code": "UNAUTHORIZED"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}", "code": "INVALID_TOKEN"})
			c.Abort()
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token", "code": "INVALID_TOKEN"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims", "code": "INVALID_TOKEN"})
			c.Abort()
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID in token", "code": "INVALID_TOKEN"})
			c.Abort()
			return
		}

		// Store userID in context
		c.Set("userID", userID)
		c.Next()
	}
}

func RateLimiter(rdb *redisclient.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Identify by client IP or User ID if authenticated
		identifier := c.ClientIP()
		if userID, exists := c.Get("userID"); exists {
			identifier = userID.(string)
		}

		endpoint := c.FullPath()
		key := fmt.Sprintf("rate:%s:%s", identifier, endpoint)

		ctx := context.Background()
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			// Fail-safe to allow request in case of Redis failure
			c.Next()
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded", "code": "RATE_LIMIT_EXCEEDED"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Log and return the first error
			err := c.Errors.Last()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"code":  "INTERNAL_SERVER_ERROR",
			})
		}
	}
}
