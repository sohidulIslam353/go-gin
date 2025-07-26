package router

import (
	"gin-bun-cockroach/internal/api/http/controllers"
	"gin-bun-cockroach/internal/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// ✅ Root route
	router.GET("/", controllers.IndexPage)
	// ✅ Register API routes
	routes.RegisterRoutes(router)
	return router
}
