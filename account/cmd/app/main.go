package main

import (
	"account/internal/router"
	"account/pkg/core/config"
	"account/pkg/core/database"
	"account/pkg/core/database/migrations"
)

func init() {
	config.Init()
}

func main() {
	database.OpenPostgresConnect()
	migrations.InitMigrations()
	go router.AmqpMainRouter()
	select {}

	//select {}
}
