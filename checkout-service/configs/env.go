package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const errorMessage string = "Error loading .env file."

func EnvAuthHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("AUTH_HOST")
}

func EnvCartHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("CART_HOST")
}

func EnvCheckoutHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("CHECKOUT_HOST")
}

func EnvOrderHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("ORDER_HOST")
}

func EnvShippingHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("SHIPPING_HOST")
}

func EnvKafkaBrokerPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}
	return os.Getenv("KAFKA_BROKER_PORT")
}

func EnvKafkaUserDeletedTopic() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}
	return os.Getenv("KAFKA_TOPIC")
}

func EnvServicePort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("CHECKOUT_SERVICE_PORT")
}

func EnvAuthServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("AUTH_SERVICE_GRPC_PORT")
}

func EnvCartServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("CART_SERVICE_GRPC_PORT")
}

func EnvOrderServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("ORDER_SERVICE_GRPC_PORT")
}

func EnvShippingServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("SHIPPING_SERVICE_GRPC_PORT")
}
