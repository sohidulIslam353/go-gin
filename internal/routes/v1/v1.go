package v1

import (
	"gin-bun-cockroach/internal/api/http/controllers"
	"gin-bun-cockroach/internal/api/http/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(rg *gin.RouterGroup) {
	// Authentication routes
	auth := rg.Group("/auth").Use(middleware.GuestOnlyMiddleware())
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	// This group will use the AuthMiddleware to protect routes
	protected := rg.Group("/user").Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.Profile)
		protected.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)
	}

	// Category routes
	categories := rg.Group("/categories")
	{
		categories.GET("/", controllers.GetCategories)
		categories.GET("/:id", controllers.GetSingleCategory)
		categories.POST("/", controllers.CreateCategory)
		categories.PUT("/:id", controllers.UpdateCategory)
		categories.DELETE("/:id", controllers.DeleteCategory)
	}
}
