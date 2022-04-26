# Endpoints

| Route              | service     |
| ------------------ | ----------- |
| /                  | healthcheck |
| /deep/             | healthcheck |
| /swagger/\*        | swagger     |
| /auth/login        | auth        |
| /auth/logout       | auth        |
| /auth/verify-token | auth        |

# Setup

1. Install dependencies using `go mod download`
2. Replicate .env.example to .env
3. For Swagger, run `swag init`
4. Run `go run main.go`
