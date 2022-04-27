# Auth-service

Contains the Auth Microservice for the final Swiggy I++ E-Commerce Application project created to showcase the skills learnt throughout the programme.

# Responsibilities

- It is responsible for validating the user's credentials and generating the JWT token.
- It also handles the user's logout.
- It also provides REST and GRPC endpoints for validating the user's details.

# Architecture

![Auth Flow](https://github.com/swiggy-2022-bootcamp/cdp-team3/blob/main/auth-service/screenshots/Auth-Flow.png)

## REST Endpoints

![Auth Service](https://github.com/swiggy-2022-bootcamp/cdp-team3/blob/main/auth-service/screenshots/Auth-Service.png)

## GRPC Endpoints

- VerifyToken: To validate the user's credentials.

## Middleware

- auth.midddleware.go : Contains the middleware to validate the user's credentials using GRPC call, its design like plug and play.

## Endpoints Design

| Route              | service     | Description                                                               |
| ------------------ | ----------- | ------------------------------------------------------------------------- |
| /                  | healthcheck | App level Health Check                                                    |
| /deep/             | healthcheck | Service level Health Check                                                |
| /swagger/\*        | swagger     | Swagger Documentation                                                     |
| /auth/login        | auth        | Login Route essentially set cookie(jwt-token) after validating user creds |
| /auth/logout       | auth        | Clear Cookie                                                              |
| /auth/verify-token | auth        | With valid jwt token in cookie, it returns claims stored in that          |

## Technologies Used

- Golang
- Gin Gonic
- gRPC
- DynamoDB
- Kafka
- Docker & Docker Compose
- Jenkins
- Swagger UI

# Checklist

- [x] Dockerized the services.
- [x] Swagger Documentation for all the services
- [x] Effective Use of GRPC for inter-service communication.
- [x] Sonarqube to calculate code coverage.
- [x] Implemented REST endpoints using DynamoDB.
- [x] Maintained the code repository
- [x] Built the services using CI/CD pipeline
- [x] Data models properly showcased
- [x] Kafka Implementation
- [x] Mock Testing using mockgen and testify
- [x] Logging in a separate file

# Setup

1. Install dependencies using `go mod download`
2. Replicate .env.example to .env
3. For Swagger, run `swag init`
4. Run `go run main.go`
