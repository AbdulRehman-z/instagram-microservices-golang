package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(Host, Port, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     Host + ":" + Port,
		Password: password,
		DB:       db,
	})

	return &RedisClient{client}
}

// Set sets the given key to the string value. If expiration is 0, the key is set without expiration.
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

// Get gets the value of the given key.
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// Del deletes the given key.
func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

// AddToStream adds the given message to the stream stored at key.
func (r *RedisClient) AddToStream(ctx context.Context, key string, message interface{}) error {
	return r.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: key,
		Values: map[string]interface{}{"message": message},
	}).Err()
}
