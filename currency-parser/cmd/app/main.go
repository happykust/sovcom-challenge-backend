package main

import (
	"currency-parser/pkg/core/config"
	"currency-parser/pkg/core/database"
)

func init() {
	config.Init()
}

func main() {
	database.OpenPostgresConnect()
}
