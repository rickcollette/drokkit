// ./config/redis.go
package config

import (
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

// InitRedis initializes a Redis client using environment variables
func InitRedis() *redis.Client {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, proceeding with environment variables")
	}

	// Get Redis address and password from environment variables
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatal("REDIS_ADDR is required but not set")
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // No password set if empty
		DB:       0,             // Use default DB
	})

	// Test the connection
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis!")
	}

	return rdb
}
