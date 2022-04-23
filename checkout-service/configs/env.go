package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvKafkaBrokerAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
	return os.Getenv("KAFKA_BROKER_ADDRESS")
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

func EnvCartServiceGRPCAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("CART_SERVICE_GRPC_ADDRESS")
}

func EnvShippingServiceGRPCAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("SHIPPING_SERVICE_GRPC_ADDRESS")
}

func EnvOrderGRPCAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("ORDER_SERVICE_GRPC_ADDRESS")
}
