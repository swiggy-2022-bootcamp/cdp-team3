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

func EnvServicePort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("CHECKOUT_SERVICE_PORT")
}

func CartServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("CART_SERVICE_GRPC_PORT")
}

func ShippingServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("SHIPPING_SERVICE_GRPC_PORT")
}
