package migrations

import (
	"deals/internal/app/domain/currencyDeals"
	"deals/internal/app/domain/simpleDeals"
	"deals/pkg/core/database"
	logger "deals/pkg/logging"
	LoggerTypes "deals/pkg/logging/types"
	"os"
)

func InitMigrations() bool {
	migrateConfig := os.Getenv("POSTGRES_MIGRATIONS")
	if migrateConfig == "" {
		logger.Log(LoggerTypes.CRITICAL, "Could not find POSTGRES_MIGRATIONS env variable", nil)
		return false
	}

	if migrateConfig == "true" {
		logger.Log(LoggerTypes.INFO, "Migrations are enabled", nil)
		err := MakeMigrations()
		if err != nil {
			logger.Log(LoggerTypes.CRITICAL, "Could not make migrations", err)
			return false
		}
		return true
	}

	logger.Log(LoggerTypes.INFO, "Migrations are disabled or the key is unknown", nil)

	return false
}

func MakeMigrations() error {
	dbConnect := &database.PG
	err := dbConnect.AutoMigrate(&simpleDeals.SimpleDeal{}, &currencyDeals.CurrencyDeal{})
	if err != nil {
		return err
	}
	return nil
}
