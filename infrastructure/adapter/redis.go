package adapter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"rizhua.com/infrastructure/etc"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     etc.C.Redis.Addr,
		Password: etc.C.Redis.Password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 检查是否成功连接到了 redis 服务器
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}
