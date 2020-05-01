package services

import (
	"os"

	"github.com/go-redis/redis/v7"
)

// GetRedisClient returns a Redis client.
func GetRedisClient() *redis.Client {
	url := os.Getenv("REDIS_URL")
	password := os.Getenv("REDIS_PASSWORD")

	return redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       0,
	})
}
