package main

import (
	"context"
	"currency-parser/internal/app"
	"currency-parser/pkg/core/config"
	"currency-parser/pkg/core/database"
	logger "currency-parser/pkg/logging"
	LoggerTypes "currency-parser/pkg/logging/types"
)

func init() {
	config.Init()
}

func main() {
	logger.Log(LoggerTypes.INFO, "[Currency-parser | Main] Start currency-parser microservice", nil)
	database.OpenPostgresConnect()
	database.InitMigrations()
	database.OpenRedisConnect()

	// Erase all tickers groups
	allTickersGroups := database.Redis.SMembers(context.Background(), config.RedisTickersGroupsSet).Val()
	for _ = range allTickersGroups {
		database.Redis.SPop(context.Background(), config.RedisTickersGroupsSet)
	}

	app.StartParsers()
}
