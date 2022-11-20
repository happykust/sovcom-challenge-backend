package migrations

import (
	"loyality/internal/app/domain/promocodes"
	"loyality/pkg/core/database"
	logger "loyality/pkg/logging"
	LoggerTypes "loyality/pkg/logging/types"
	"os"
)

func InitMigrations() bool {
	migrateConfig := os.Getenv("POSTGRES_MIGRATIONS")
	if migrateConfig == "" {
		logger.Log(LoggerTypes.CRITICAL,
			"[Loyality | Database] Could not find POSTGRES_MIGRATIONS env variable", nil)
		return false
	}

	if migrateConfig == "true" {
		logger.Log(LoggerTypes.INFO, "[Loyality | Database] Migrations are enabled", nil)
		err := MakeMigrations()
		if err != nil {
			logger.Log(LoggerTypes.CRITICAL, "[Loyality | Database] Could not make migrations", err)
			return false
		}
		return true
	}

	logger.Log(LoggerTypes.INFO,
		"[Loyality | Database] Migrations are disabled or the key is unknown", nil)

	return false
}

func MakeMigrations() error {
	dbConnect := &database.PG
	err := dbConnect.AutoMigrate(&promocodes.Promocode{})
	if err != nil {
		return err
	}
	return nil
}
