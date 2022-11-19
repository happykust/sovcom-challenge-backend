package main

import (
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
	app.StartParsers()
}
