package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"www.github.com/ygxiaobai111/qiniu/server/config"

	logging "github.com/sirupsen/logrus"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client
var RedisContext = context.Background()

// Init 在中间件中初始化redis链接
func Init() {

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPw,
		DB:       int(config.RedisDbName),
	})
	_, err := client.Ping(RedisContext).Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
