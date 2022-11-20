package main

import (
	"deals/internal/app/routers"
	"deals/pkg/core/config"
	"deals/pkg/core/database"
	"deals/pkg/core/database/migrations"
)

func init() {
	config.Init()
	database.OpenPostgresConnect()
	migrations.InitMigrations()
	database.OpenRedisConnect()
}

func main() {
	go routers.SimpleDealsAmqpRouter()
	go routers.CurrencyDealsAmqpRouter()
	select {}
}
