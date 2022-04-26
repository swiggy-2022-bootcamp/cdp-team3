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

func EnvGrpcRewardClientPORT() string {
	loadEnvFile()
	return os.Getenv("GRPC_REWARD_CLIENT_PORT")
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

func EnvRewardServicePORT() string {
	loadEnvFile()
	return os.Getenv("REWARD_SERVICE_PORT")
}

func EnvAdminHost() string {
	loadEnvFile()
	return os.Getenv("ADMIN_HOST")
}

func EnvAuthHost() string {
	loadEnvFile()
	return os.Getenv("AUTH_HOST")
}

func EnvRewardsHost() string {
	loadEnvFile()
	return os.Getenv("REWARDS_HOST")
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		zap.L().Fatal("Error loading .env file.")
	}
}
