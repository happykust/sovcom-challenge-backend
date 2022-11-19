package migrations

import (
	"account/internal/domain/auth"
	"account/internal/domain/user"
	"account/pkg/core/database"
	logger "account/pkg/logging"
	LoggerTypes "account/pkg/logging/types"
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
	err := dbConnect.AutoMigrate(&auth.UnverifiedUsers{}, &user.User{})
	if err != nil {
		return err
	}
	return nil
}
