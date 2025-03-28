package redis_client

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Re *redis.Client

func InitRedis() error {
	Re = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	err := Re.Ping(context.Background()).Err()
	return err
}
