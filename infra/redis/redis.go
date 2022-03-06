package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"go-vote/config"
	"time"
)

func Init(cfg config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDb,
	})
	return rdb
}

func Get(client *redis.Client, key string, value interface{}) error {
	ctx := context.Background()
	cache, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Errorf("failed to get redis cache: %v", err)
		return err
	}
	err = json.Unmarshal([]byte(cache), &value)
	if err != nil {
		log.Errorf("error unmarshal json: %v", err)
		return err
	}
	return nil
}

func Set(client *redis.Client, key string, value interface{}, expirationSecond int) error {
	ctx := context.Background()
	cache, err := json.Marshal(value)
	if err != nil {
		log.Errorf("error marshal json: %v", err)
		return err
	}
	err = client.Set(ctx, key, cache, time.Second*time.Duration(expirationSecond)).Err()
	if err != nil {
		log.Errorf("failed to set redis cache: %v", err)
	}
	return err
}
