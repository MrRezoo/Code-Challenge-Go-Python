package db

import (
	"context"
	"github.com/MrRezoo/code-challenge/config"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()
var RedisClient *redis.Client

func InitRedis(config *config.RedisConfig) {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	pingResult, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	} else {
		log.Printf("Connected to Redis successfully: %s", pingResult)
	}
}
