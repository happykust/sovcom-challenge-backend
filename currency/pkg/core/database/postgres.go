package database

import (
	"currency/pkg/core/config"
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
}
