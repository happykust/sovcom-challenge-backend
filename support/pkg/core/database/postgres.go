package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"support/pkg/core/config"
)

var PG gorm.DB

func OpenPostgresConnect() {
	db, err := gorm.Open(postgres.Open(config.GetPostgresDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	PG = *db
}
