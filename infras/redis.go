package cache

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Cache *RedisCache
)

func init() {
	redis := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "password",
		DB:       0,
	})
	ctx := context.Background()
	_, err := redis.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	Cache = &RedisCache{
		redis: redis,
	}
}

type RedisCache struct {
	RWMutex sync.RWMutex
	redis   *redis.Client
}

func (r *RedisCache) Ping(ctx context.Context) (string, error) {
	return r.redis.Ping(ctx).Result()
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	r.RWMutex.Lock()
	defer r.RWMutex.Unlock()
	byteVal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, key, byteVal, ttl).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, bool, error) {
	byteVal, err := r.redis.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return byteVal, true, nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	r.RWMutex.Lock()
	defer r.RWMutex.Unlock()

	return r.redis.Del(ctx, key).Err()
}

func (r *RedisCache) Close() {
	r.redis.Close()
}
