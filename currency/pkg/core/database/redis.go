package database

import (
	"currency/pkg/core/config"
	logger "currency/pkg/logging"
	LoggerTypes "currency/pkg/logging/types"
	"github.com/go-redis/redis/v7"
)

var Redis *redis.Client

func OpenRedisConnect() {
	opt, err := redis.ParseURL(config.GetRedisDSN())
	if err != nil {
		panic(err)
	}

	Redis = redis.NewClient(opt)
	logger.Log(LoggerTypes.INFO, "[Currency | Redis] Redis connection opened", nil)
}
