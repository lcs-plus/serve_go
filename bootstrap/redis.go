package bootstrap

import (
	"0121_1/global"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Addr,
		Password: global.App.Config.Redis.Password,
		DB:       global.App.Config.Redis.DB,
	})
}

func SetKeyValue(key string, value interface{}, expiresAt time.Duration) {
	ctx := context.Background()
	global.App.Redis.Set(ctx, key, value, expiresAt)
}
