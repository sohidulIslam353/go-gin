package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// if the user is already logged in, they should not be able to access guest-only routes
func GuestOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Already logged in"})
			return
		}
		c.Next()
	}
}
