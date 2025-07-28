package middleware

import (
	"gin-bun-cockroach/config"
	"gin-bun-cockroach/internal/api/http/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if token is revoked
		val, err := config.RedisClient.Get(config.Ctx, tokenStr).Result()
		if err == nil && val == "revoked" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is revoked"})
			return
		}

		// Token is valid
		c.Set("userID", claims.UserID)
		c.Set("userName", claims.Name)
		c.Set("userEmail", claims.Email)
		c.Next()
	}
}

func AuthUser(c *gin.Context) (uint, string, string) {
	userID, _ := c.Get("userID")
	name, _ := c.Get("userName")
	email, _ := c.Get("userEmail")
	return userID.(uint), name.(string), email.(string)
}
