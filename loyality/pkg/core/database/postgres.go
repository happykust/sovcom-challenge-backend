package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"loyality/pkg/core/config"
	logger "loyality/pkg/logging"
	LoggerTypes "loyality/pkg/logging/types"
)

var PG gorm.DB

func OpenPostgresConnect() {
	db, err := gorm.Open(postgres.Open(config.GetPostgresDSN()), &gorm.Config{})
	if err != nil {
		logger.Log(LoggerTypes.ERROR, "[Loyality | Database] Database connection error!", nil)
	}
	PG = *db
}
