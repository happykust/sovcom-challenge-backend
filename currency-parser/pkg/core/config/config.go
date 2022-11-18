package config

import (
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
}

func GetPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_PORT"))
}
