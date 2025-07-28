// internal/config/redis.go

package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // If using KeyDB, port is the same
		Password: "",               // Set if any
		DB:       0,                // Default DB
	})
	log.Println("✅ Connected to Redis")
}
