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

func CartServiceGRPCAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("CART_SERVICE_GRPC_ADDRESS")
}

func ShippingServiceGRPCAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("SHIPPING_SERVICE_GRPC_ADDRESS")
}
