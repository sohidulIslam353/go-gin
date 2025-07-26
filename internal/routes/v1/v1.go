package v1

import (
	"gin-bun-cockroach/internal/api/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(rg *gin.RouterGroup) {
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
