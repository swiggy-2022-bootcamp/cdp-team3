package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvAccessKey() string {
	loadEnvFile()
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func EnvSecretKey() string {
	loadEnvFile()
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func EnvRegion() string {
	loadEnvFile()
	return os.Getenv("REGION")
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
}
func EnvShippingServiceGRPCPort() string {
	loadEnvFile()

	return os.Getenv("SHIPPING_SERVICE_GRPC_PORT")
}
func EnvShippingPort() string {
	loadEnvFile()
	return os.Getenv("SHIPPING_SERVICE_PORT")
}
func EnvShippingHost() string {
	loadEnvFile()

	return os.Getenv("SHIPPING_HOST")
}
func EnvAuthServiceGRPCPort() string {
	loadEnvFile()

	return os.Getenv("AUTH_SERVICE_GRPC_PORT")
}
func EnvAuthHost() string {
	loadEnvFile()

	return os.Getenv("AUTH_HOST")
}
func EnvProductsHost() string {
	loadEnvFile()

	return os.Getenv("PRODUCTS_HOST")
}
func EnvProductsServiceGRPCPort() string {
	loadEnvFile()

	return os.Getenv("PRODUCT_SERVICE_GRPC_PORT")
}