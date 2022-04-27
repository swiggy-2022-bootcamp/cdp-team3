# Payment-service

Contains the Payment Microservice for the final Swiggy I++ E-Commerce Application project created to showcase the skills learnt throughout the programme.

# Responsibilities

- It will handle the payment related operations.
- It will make a GRPC call to Order Service to validate the order details.
- It will make a GRPC call to Mode of Payment Service to validate the payment details.
- It will update the order status to "COMPLETED/FAILED" after successful/unsuccessful payment respectively using kafka.

# Architecture

![Payment Flow](https://github.com/swiggy-2022-bootcamp/cdp-team3/blob/main/payment-service/screenshots/Payment-Flow.png)

## REST Endpoints

![Payment Service](https://github.com/swiggy-2022-bootcamp/cdp-team3/blob/main/payment-service/screenshots/Payment-Service.png)

## Endpoints Design

| Route       | service     | Description                                                  |
| ----------- | ----------- | ------------------------------------------------------------ |
| /           | healthcheck | App level Health Check                                       |
| /deep/      | healthcheck | Service level Health Check                                   |
| /swagger/\* | swagger     | Swagger Documentation                                        |
| /pay        | payment     | Handling payment after validating order and user credentials |

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
