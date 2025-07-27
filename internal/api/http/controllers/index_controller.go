package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome To Gin Ecommerce API.",
	})
}
