package redis

import (
	"bluebell/setting"
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func Init(cfg *setting.RedisConfig) (err error) {
	option := &redis.Options{
		Addr:               fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:           cfg.Password,
		DB:                 cfg.DB,
		PoolSize:           cfg.PoolSize,
		MinIdleConns:       cfg.MinIdleConns,
	}
	client = redis.NewClient(option)
	_, err = client.Ping().Result()
	return
}

func Close() {
	_ = client.Close()
}
