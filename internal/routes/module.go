package routes

import (
	v1 "gin-bun-cockroach/internal/routes/v1"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	v1.RegisterV1Routes(router.Group("/api/v1"))
}
