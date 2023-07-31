package repository

import (
	"authentication-service/config"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client *redis.Client
}

type IRedisRepo interface {
	Add(ctx context.Context, key string, value interface{}, exp time.Time) error
	Get(ctx context.Context, key string) (interface{}, error)
}

func NewRedisRepo(redisConfig *config.RedisConfig) IRedisRepo {
	return &RedisRepo{client: redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       0,                    // use default DB
	})}
}

func (r *RedisRepo) Add(ctx context.Context, key string, value interface{}, exp time.Time) error {
	return r.client.SetArgs(ctx, key, value, redis.SetArgs{
		ExpireAt: exp,
	}).Err()
}

func (r *RedisRepo) Get(ctx context.Context, key string) (interface{}, error) {
	return r.client.Get(ctx, key).Result()
}
