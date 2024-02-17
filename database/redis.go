package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return nil, fmt.Errorf("failed to connect to Redis: %s", err)
	}
	fmt.Println("connection to Redis established")

	return redisClient, nil
}
