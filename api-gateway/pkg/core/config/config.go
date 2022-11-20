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

func GetAMQPUri() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", os.Getenv("AMQP_USER"), os.Getenv("AMQP_PASS"),
		os.Getenv("AMQP_HOST"), os.Getenv("AMQP_PORT"))
}
