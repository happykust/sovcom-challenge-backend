package config

import (
	logger "currency-parser/pkg/logging"
	LoggerTypes "currency-parser/pkg/logging/types"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logger.Log(LoggerTypes.INFO, "[Currency-parser | Config] Config initialized", nil)
}

func GetPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_PORT"))
}

func GetRedisDSN() string {
	return fmt.Sprintf(
		"redis://%s:%s@%s:%s/%s",
		os.Getenv("REDIS_USER"), os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"), os.Getenv("REDIS_DB"))
}
