package database

import (
	"context"
	"fmt"
	"log"

	"stack-bm/internal/config"

	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
)

func InitRedis() {
	if config.AppConfig.Redis.Host == "" {
		log.Println("Redis: REDIS_HOST not configured, skip")
		return
	}
	addr := fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, config.AppConfig.Redis.Port)
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	})
	ctx := context.Background()
	if err := Redis.Ping(ctx).Err(); err != nil {
		log.Printf("Redis connect failed: %v", err)
		return
	}
	fmt.Printf("Connected to Redis: %s\n", addr)
}
