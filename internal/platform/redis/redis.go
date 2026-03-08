package redis

import (
	"github.com/rdl/core/internal/platform/config"
	"github.com/redis/go-redis/v9"
)

type Client = redis.Client

func Connect(cfg config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return rdb, nil
}
