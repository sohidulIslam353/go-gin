package controllers

import (
	"gin-bun-cockroach/internal/api/http/middleware"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userID, userName, userEmail := middleware.AuthUser(c)
	c.JSON(200, gin.H{
		"message":   "User profile",
		"userID":    userID,
		"userName":  userName,
		"userEmail": userEmail,
	})
}
