## Admin Service
- Part Of Swiggy IPP Final Team Project that includes the implementation of the Admin Service which handles the CRUD operations on Customer using REST, gRPC and Kafka for relevant synchronous and asynchronous communications with other services.
 
### REST Endpoints (:3008)

| Http Method |          Enpoint          |              Description              |
| ----------- | :-----------------------: | :-----------------------------------: |
| POST        |          /admin           |          To Create an Admin           |
| GET         |        /admin/user        | To get the details of logged in Admin |
| POST        |        /customers         |         To create a customer          |
| GET         |  /customers/:customerId   |  To Get the customer by customer ID   |
| GET         | /customers/email/:emailId |    To get the customer by email ID    |
| PUT         |  /customers/:customerId   | To update the customer by customer Id |
| DELETE      |  /customers/:customerId   | To delete the customer by customer Id |


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
- [x] Swagger Documentation - http://localhost:3009/swagger/index.html
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
![Untitled Diagram drawio (6)](https://user-images.githubusercontent.com/19664740/165362861-27ee7f90-f45c-40db-a7a3-13a4cf0b85c6.png)


## File Structure
![adminfileservice](https://user-images.githubusercontent.com/19664740/165359544-e1095e43-9f30-4f47-8bd7-fe8178b603b0.PNG)

## gRPC Communication
- Admin Service(Client) -> Auth Service (Server) - To get the token verified and only allow the admin to perform the actions
- Rewards Service(Client) -> Admin Service (Server) - To update the reward points in the Customers Table. Once the Sucess message goes out to Rewards Service, Reward Service Appends the Reward Details to Reward Table on DynamoDB
- Transaction Service(Client) -> Admin Service (Server) - To update the transaction points after every successful order completion, in the Customers Table. Once the Sucess message goes out to Transaction Service, Transaction Service Appends the Details to Transactions Table on DynamoDB
- Admin Service(Client) -> Shipping Service (Server)
  * To send the address object to Shipping Service during the creation of user so that Shipping Service can parse that and add it to Shipping Address Table and return a ShippingAddressID which will be stored in the Customers Table along with Address. This ShippingAddressID will be used for communication later with other services for payment flow and checkout flow.
  * To send the ShippingAddressID of the Address during the deletion of the user so that Shipping Service can delete the saved Address in Shipping Address Table.

## Kafka
Admin Service(Producer) -> Cart Service (Consumer) - Kafka is used to produce messages in the User Creation & Deletion Topics which will be consumed by the Cart Service to create an empty cart when the user is created and to delete it when the user is deleted. Because there's no dependency on the deletion or creation, asynchronous communication was thought to be the best way to achieve it.

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
> http://localhost:3009/swagger/index.html
![admin_swagger](https://user-images.githubusercontent.com/19664740/165361453-669d8a6d-f495-4d4f-bcee-37b7312cf500.PNG)
