package redis

import (
	"github.com/go-redis/redis/v8"

	"go-vote/config"
)

func Init(cfg config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDb,
	})
	return rdb
}