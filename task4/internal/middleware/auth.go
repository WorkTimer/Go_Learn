package middleware

import (
	"blog-system/pkg/auth"
	"blog-system/pkg/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Error("Authorization header is required")
			c.JSON(401, gin.H{"code": 401, "message": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			logger.Error("Invalid authorization header format")
			c.JSON(401, gin.H{"code": 401, "message": "Invalid authorization header format"})
			c.Abort()
			return
		}

		claims, err := auth.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			logger.Error("Invalid token:", err)
			c.JSON(401, gin.H{"code": 401, "message": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

func OptionalAuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.Next()
			return
		}

		claims, err := auth.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			c.Next()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
