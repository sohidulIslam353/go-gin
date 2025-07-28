package main

import (
	"gin-bun-cockroach/config"
	"gin-bun-cockroach/internal/pkg/router"
)

func main() {
	config.InitDB()
	config.ConnectRedis()
	r := router.SetupRouter()
	r.Run(":8080") // or your preferred port
}
