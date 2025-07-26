package main

import (
	"gin-bun-cockroach/config"
	"gin-bun-cockroach/internal/pkg/router"
)

func main() {
	config.InitDB()
	r := router.SetupRouter()
	r.Run(":8080") // or your preferred port
}
