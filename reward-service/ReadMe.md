## Rewards Service - CDP Team 3
- Part Of Swiggy IPP Final Team Project that includes the implementation of the Rewards Service which handles the rewards Of The User.


## REST Endpoints (:3009)

| Http Method |        Enpoint         |                    Description                    |
| ----------- | :--------------------: | :-----------------------------------------------: |
| GET         |        /rewards        |         Get List of all the added Rewards         |
| POST        |        /rewards        |                  Add New Reward                   |
| GET         |  /rewards/{rewardId}   |          Get Reward details by Reward ID          |
| GET         | /rewards/user/{userId} | Get all the Rewards for a customer by Customer ID |

## Technologies Used
- Golang
- Gin Gonic
- gRPC
- DynamoDB
- Kafka
- Docker & Docker Compose
- Jenkins
- Swagger UI

## Checklist
- [x] Completed Microservice Architecture
- [x] Written test cases for service layer. 
- [x] Dockerized the application.
- [x] Swagger Documentation - http://localhost:3008/swagger/index.html
- [x] Effective Use of GRPC for inter-service communication.
- [x] Sonarqube to calculate code coverage.
- [x] Implemented REST endpoints using DynamoDB.
- [x] Maintained the code repository
- [x] Builds the application using CI/CD pipeline
- [x] Data models properly showcased
- [x] Kafka Implementation
- [x] Mock Testing using mockgen and testify
- [x] Zap Logger - Logging in a separate file

## Microservice Architecture
![Untitled Diagram drawio (2)](https://user-images.githubusercontent.com/19664740/165355916-3410be1c-c62d-4056-85cf-1ffa09f625ce.png)

## File Structure
![reward service file](https://user-images.githubusercontent.com/19664740/165356787-12c6ae3f-d92f-402b-960b-46d6f5cf3821.PNG)

## gRPC Communication
- Rewards Service(Client) -> Auth Service (Server) - To get the token verified and only allow the admin to perform the actions
- Rewards Service(Client) -> Admin Service (Server) - To update the reward points in the Customers Table. Once the Sucess message comes in, Reward Service Appends the Reward Details to Reward Table on DynamoDB

## Running The Application 

Run locally 
```sh
go run main.go
```

Generate swagger docs
```sh
swag init
```

## Swagger UI
> http://localhost:3008/swagger/index.html
![image](https://user-images.githubusercontent.com/19664740/165358229-f8bead13-e55e-4e96-b477-6c6b9007039b.png)

