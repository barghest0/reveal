package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

// CacheService определяет интерфейс для работы с кэшем.
type CacheService interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, dest interface{}) error
	Delete(key string) error
}

// cache реализует интерфейс CacheService.
type cache struct {
	client *redis.Client
	ctx    context.Context
}

// CreateCacheService создает новый экземпляр redisCache.
func CreateCacheService(client *redis.Client) CacheService {
	return &cache{
		client: client,
		ctx:    context.Background(),
	}
}

// Set сохраняет данные в кэше.
func (r *cache) Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(r.ctx, key, data, expiration).Err()
}

// Get получает данные из кэша.
func (r *cache) Get(key string, dest interface{}) error {
	data, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

// Delete удаляет данные из кэша.
func (r *cache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
