package main

import (
	"currency/internal/app/routers"
	"currency/pkg/core/config"
	"currency/pkg/core/database"
	"currency/pkg/core/database/migrations"
	logger "currency/pkg/logging"
	LoggerTypes "currency/pkg/logging/types"
	"currency/server"
)

func init() {
	config.Init()
	database.OpenPostgresConnect()
	migrations.InitMigrations()
	database.OpenRedisConnect()
}

func main() {
	go routers.MainAmqpRouter()
	err := server.App().Run(":7070")
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "[Currency | API Server] Error run server! Shutdown.", err)
		panic(err)
	}
	select {}
}
