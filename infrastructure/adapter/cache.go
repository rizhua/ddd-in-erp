package adapter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Rdb *redis.Client
	ctx context.Context
}

func NewCache() *Cache {
	return &Cache{
		Rdb: NewRedis(),
		ctx: context.Background(),
	}
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	return c.Rdb.Set(c.ctx, key, value, expiration).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.Rdb.Get(c.ctx, key).Result()
}

func (c *Cache) Del(key string) error {
	return c.Rdb.Del(c.ctx, key).Err()
}

func (c *Cache) Exists(key string) bool {
	return c.Rdb.Exists(c.ctx, key).Val() > 0
}

func (c *Cache) Keys(pattern string) []string {
	return c.Rdb.Keys(c.ctx, pattern).Val()
}

func (c *Cache) Expire(key string, expiration time.Duration) bool {
	return c.Rdb.Expire(c.ctx, key, expiration).Val()
}

func (c *Cache) TTL(key string) time.Duration {
	return c.Rdb.TTL(c.ctx, key).Val()
}

func (c *Cache) Incr(key string) int64 {
	return c.Rdb.Incr(c.ctx, key).Val()
}

func (c *Cache) Decr(key string) int64 {
	return c.Rdb.Decr(c.ctx, key).Val()
}

func (c *Cache) HSet(key string, field string, value interface{}) error {
	return c.Rdb.HSet(c.ctx, key, field, value).Err()
}

func (c *Cache) HGet(key string, field string) string {
	return c.Rdb.HGet(c.ctx, key, field).Val()
}

func (c *Cache) HDel(key string, fields ...string) error {
	return c.Rdb.HDel(c.ctx, key, fields...).Err()
}
