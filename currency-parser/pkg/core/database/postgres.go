package database

import (
	"currency-parser/pkg/core/config"
	logger "currency-parser/pkg/logging"
	LoggerTypes "currency-parser/pkg/logging/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var PG gorm.DB

func OpenPostgresConnect() {
	db, err := gorm.Open(postgres.Open(config.GetPostgresDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	PG = *db
	logger.Log(LoggerTypes.INFO, "[Currency-parser | Postgres] Postgres connection opened", nil)
}
