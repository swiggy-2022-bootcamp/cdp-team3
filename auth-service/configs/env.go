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
	return os.Getenv("AUTH_PORT")
}

func EnvGRPCPORT() string {
	loadEnv()
	return os.Getenv("AUTH_GRPC_PORT")
}