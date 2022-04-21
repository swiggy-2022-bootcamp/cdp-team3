package configs

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func EnvAccessKey() string {
	loadEnvFile()
	return "fakeAccessKey"
	// return os.Getenv("AWS_ACCESS_KEY_ID")
}

func EnvSecretKey() string {
	loadEnvFile()
	return "fakeSecretAccessKey"
	// return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func EnvRegion() string {
	loadEnvFile()
	return "ap-southeast-2"
	// return os.Getenv("REGION")
}

func loadEnvFile() {
	return
	err := godotenv.Load()
	if err != nil {
		zap.L().Fatal("Error loading .env file.")
	}
}
