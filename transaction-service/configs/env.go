package configs

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
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

func EnvPORT() string {
	loadEnvFile()
	return os.Getenv("TRANSACTION_SERVICE_PORT")
}

func EnvGrpcAuthClientPORT() string {
	loadEnvFile()
	return os.Getenv("AUTH_CLIENT_GRPC_PORT")
}

func EnvGrpcAdminClientPORT() string {
	loadEnvFile()
	return os.Getenv("ADMIN_CLIENT_GRPC_PORT")
}

func EnvAddTransactionAmountBrokerAddress() string {
	loadEnvFile()
	return os.Getenv("KAFKA_BROKER_ADDRESS")
}

func EnvAddTransactionAmountTopic() string {
	loadEnvFile()
	return os.Getenv("ADD_TRANSACTION_AMOUNT_TOPIC")
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		zap.L().Fatal("Error loading .env file.")
	}
}