package controllers

import (
	"net/http"
	"strings"
	"time"

	"gin-bun-cockroach/config"
	"gin-bun-cockroach/internal/api/http/auth"
	"gin-bun-cockroach/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Name     string `form:"name" binding:"required"`
		Email    string `form:"email" binding:"required,email"`
		Phone    string `form:"phone" binding:"required"`
		Password string `form:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}

	user := models.User{
		Name:      body.Name,
		Email:     body.Email,
		Phone:     body.Phone,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = config.DB.NewInsert().Model(&user).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful!"})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := config.DB.NewSelect().Model(&user).Where("email = ?", body.Email).Scan(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	token, err := auth.GenerateJWT(uint(user.ID), user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Missing Authorization header"})
		return
	}

	token := strings.Replace(authHeader, "Bearer ", "", 1)

	// Set token in Redis blacklist
	err := config.RedisClient.Set(config.Ctx, token, "blacklisted", time.Hour*24).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to blacklist token"})
		return
	}

	c.JSON(200, gin.H{"message": "Logged out successfully"})
}
