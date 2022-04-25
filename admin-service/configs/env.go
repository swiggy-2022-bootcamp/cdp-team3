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

func EnvGrpcRewardServerPORT() string {
	loadEnvFile()
	return os.Getenv("GRPC_REWARD_SERVER_PORT")
}

func EnvGrpcTransactionServerPORT() string {
	loadEnvFile()
	return os.Getenv("GRPC_TRANSACTION_SERVER_PORT")
}

func EnvGrpcAuthClientPORT() string {
	loadEnvFile()
	return os.Getenv("GRPC_AUTH_CLIENT_PORT")
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
	return os.Getenv("ADMIN_SERVICE_PORT")
}

func EnvAddUserAdditionTopic() string {
	loadEnvFile()
	return os.Getenv("USER_ADDITION_TOPIC")
}

func EnvAddUserDeletionTopic() string {
	loadEnvFile()
	return os.Getenv("USER_DELETION_TOPIC")
}

func EnvAddUserBrokerAddress() string {
	loadEnvFile()
	return os.Getenv("KAFKA_BROKER_ADDRESS")
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		zap.L().Fatal("Error loading .env file.")
	}
}
