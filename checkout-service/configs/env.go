package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvKafkaBrokerPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
	return os.Getenv("KAFKA_BROKER_PORT")
}

func EnvKafkaUserDeletedTopic() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
	return os.Getenv("KAFKA_TOPIC")
}

func EnvServicePort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("CHECKOUT_SERVICE_PORT")
}

func EnvCartServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("CART_SERVICE_GRPC_PORT")
}

func EnvShippingServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("SHIPPING_SERVICE_GRPC_PORT")
}

func EnvOrderGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("ORDER_SERVICE_GRPC_PORT")
}
