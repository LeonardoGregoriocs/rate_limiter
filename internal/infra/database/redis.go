package database

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	"github.com/leonardogregoriocs/rate_limiter/config"
)

type RedisDatabaseInterface interface{}

type RedisDatabase struct {
	Client *redis.Client
}

func NewConnection(cfg config.Config, logger zerolog.Logger) (*RedisDatabase, error) {
	addr := fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort)

	log.Printf("Connecting to Redis on [%s]", addr)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected")

	return &RedisDatabase{
		Client: client,
	}, nil
}
