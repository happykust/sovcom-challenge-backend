package database

import (
	"deals/pkg/core/config"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"github.com/go-redis/redis/v7"
)

var Redis *redis.Client

func OpenRedisConnect() {
	opt, err := redis.ParseURL(config.GetRedisDSN())
	if err != nil {
		panic(err)
	}

	Redis = redis.NewClient(opt)
	logger.Log(LoggerTypes.INFO, "[Deals | Redis] Redis connection opened", nil)
}
