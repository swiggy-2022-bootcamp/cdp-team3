package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMonogoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
	return os.Getenv("MONGO_URI")
}

func EnvServiceAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("CART_SERVICE_ADDRESS")
}

func EnvServicePort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("CART_SERVICE_PORT")
}
