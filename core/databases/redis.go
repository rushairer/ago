package databases

import (
	"github.com/go-redis/redis/v8"
)

func NewRedis(dsn string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})
}
