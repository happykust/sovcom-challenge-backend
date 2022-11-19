package migrations

import (
	"os"
	"payments/internal/app/domain/balance"
	"payments/internal/app/domain/transactions"
	"payments/pkg/core/database"
	logger "payments/pkg/logging"
	LoggerTypes "payments/pkg/logging/types"
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
	err := dbConnect.AutoMigrate(&balance.Balance{}, &balance.Wallet{}, &transactions.Transaction{})
	if err != nil {
		return err
	}
	return nil
}
