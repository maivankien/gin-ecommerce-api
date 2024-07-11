package initialize

import (
	"context"
	"fmt"

	"github.com/maivankien/go-ecommerce-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("InitRedis initialization failed", zap.Error(err))
		panic(err)
	}

	fmt.Println("Redis connected")
	global.Rdb = rdb
}
