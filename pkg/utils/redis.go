package utils

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Rdb *redis.Client
}

func NewRedis() (*Redis, error) {
	ctx := context.Background()
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	addr := cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Index,
	})
	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return &Redis{Rdb: rdb}, nil
}

func (r *Redis) Set(key string, value interface{}) error {
	return r.Rdb.Set(context.Background(), key, value, 0).Err()
}
func (r *Redis) Get(key string) (interface{}, error) {
	return r.Rdb.Get(context.Background(), key).Result()
}
func (r *Redis) Del(key string) (int64, error) {
	return r.Rdb.Del(context.Background(), key).Result()
}
