package utils

import (
	"github.com/gin-gonic/gin"
)

func AuthUserID(c *gin.Context) uint {
	if userID, exists := c.Get("userID"); exists {
		return userID.(uint)
	}
	return 0
}

func AuthUserName(c *gin.Context) string {
	if name, exists := c.Get("userName"); exists {
		return name.(string)
	}
	return ""
}
