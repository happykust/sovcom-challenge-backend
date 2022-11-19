package database

import (
	"currency-parser/pkg/core/config"
	logger "currency-parser/pkg/logging"
	LoggerTypes "currency-parser/pkg/logging/types"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func OpenRedisConnect() {
	opt, err := redis.ParseURL(config.GetRedisDSN())
	if err != nil {
		panic(err)
	}

	Redis = redis.NewClient(opt)
	logger.Log(LoggerTypes.INFO, "[Currency-parser | Redis] Redis connection opened", nil)
}
