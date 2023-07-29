package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	AccessTokenExpiresIn = 30
	SecretKey            = "eGluaGRlcHR1eWV"
	AccessTokenMaxAge    = 24
	TOPIC                = "order_service_topic"
)

func GetMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file cause:", err.Error())
	}
	return os.Getenv("MONGODB_URL")
}
