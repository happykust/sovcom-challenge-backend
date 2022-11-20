package database

import (
	"deals/pkg/core/config"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PG gorm.DB

func OpenPostgresConnect() {
	db, err := gorm.Open(postgres.Open(config.GetPostgresDSN()), &gorm.Config{})
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not connect to postgres:", err)
		panic(err)
	}
	PG = *db
}
