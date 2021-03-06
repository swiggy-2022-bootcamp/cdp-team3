package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
}

func EnvAccessKey() string {
	loadEnv()
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func EnvAccessSecretKey() string {
	loadEnv()
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func EnvRegion() string {
	loadEnv()
	return os.Getenv("REGION")
}

func EnvJWTSecretKey() string {
	loadEnv()
	return os.Getenv("SECRET_KEY")
}

func EnvPORT() string {
	loadEnv()
	return os.Getenv("PAYMENT_PORT")
}

func EnvAuthHost() string {
	loadEnv()
	return os.Getenv("AUTH_HOST")
}

func EnvPaymentHost() string {
	loadEnv()
	return os.Getenv("PAYMENT_HOST")
}

func EnvAuthGRPCPort() string {
	loadEnv()
	return os.Getenv("AUTH_GRPC_PORT")
}

func EnvOrderHost() string {
	loadEnv()
	return os.Getenv("ORDER_HOST")
}

func EnvOrderServiceGRPCPort() string {
	loadEnv()
	return os.Getenv("ORDER_SERVICE_GRPC_PORT")
}

func EnvBrokerAddress() string {
	loadEnv()
	return os.Getenv("KAFKA_BROKER_ADDRESS")
}
