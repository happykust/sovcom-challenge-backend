package main

import (
	"loyality/internal/app/routers"
	"loyality/pkg/core/config"
	"loyality/pkg/core/database"
	"loyality/pkg/core/database/migrations"
	logger "loyality/pkg/logging"
	LoggerTypes "loyality/pkg/logging/types"
)

func init() {
	config.Init()
}

func main() {
	database.OpenPostgresConnect()
	migrationsStatus := migrations.InitMigrations()
	if !migrationsStatus {
		logger.Log(LoggerTypes.ERROR, "[Loyality | Database] Could not make migrations!", nil)
	}

	go routers.ReferralAmqpRouter()
	go routers.PromocodesAmqpRouter()
	select {}
}
