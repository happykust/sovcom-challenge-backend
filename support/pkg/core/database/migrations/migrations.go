package migrations

import (
	"os"
	"support/internal/app/api/domain/support"
	"support/pkg/core/database"
	logger "support/pkg/logging"
	LoggerTypes "support/pkg/logging/types"
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
	err := dbConnect.AutoMigrate(&support.Ticket{}, &support.TicketMessage{})
	if err != nil {
		return err
	}
	return nil
}
