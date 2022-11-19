package main

import (
	"support/pkg/core/config"
	"support/pkg/core/database"
	"support/pkg/core/database/migrations"
	logger "support/pkg/logging"
	LoggerTypes "support/pkg/logging/types"
	"support/server"
)

func init() {
	config.Init()
	database.OpenPostgresConnect()
	migrations.InitMigrations()
}

func main() {
	err := server.App().Run(":8080")
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not start server", err)
	}
}
