package main

import (
	redis "github.com/leonardogregoriocs/rate_limiter/database"
)

func main() {
	_, err := redis.NewRedisClient()
	if err != nil {
		panic(err)
	}
}
