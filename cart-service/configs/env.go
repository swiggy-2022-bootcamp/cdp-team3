package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const errorMessage string = "Error loading .env file."

func EnvKafkaBrokerAddress() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}
	return os.Getenv("KAFKA_BROKER_ADDRESS")
}

func EnvKafkaUserCreatedTopic() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}
	return os.Getenv("KAFKA_USER_CREATED_TOPIC")
}

func EnvKafkaUserDeletedTopic() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}
	return os.Getenv("KAFKA_USER_DELETED_TOPIC")
}

func EnvRegion() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1"
	}
	return region
}

func EnvAccessKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func EnvSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func EnvCartHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("CART_HOST")
}

func EnvServicePort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("CART_SERVICE_PORT")
}

func EnvCartServiceGRPCPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(errorMessage)
	}

	return os.Getenv("CART_SERVICE_GRPC_PORT")
}
