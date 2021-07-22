package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func RedisClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(client.Context()).Result()
	fmt.Println(pong, err)
	Client = client
}
