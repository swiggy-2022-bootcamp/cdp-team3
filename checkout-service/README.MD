# checkout-service

Contains the Checkout Microservice for the final Swiggy I++ E-Commerce Application project created to showcase the skills learnt throughout the programme.

## Project Structure

| ![Folder Structure](https://github.com/swiggy-2022-bootcamp/cdp-team3/blob/main/checkout-service/screenshots/folder_structure.png?raw=true) |
| --- |

## Relevant REST APIs (:3011)

| HTTP Verb  | Endpoint Path                        | Description                                                          |
| ---------- | ------------------------------------ | -------------------------------------------------------------------- |
| GET        | /checkout_service/swagger/index.html | Swagger documentation                                                |
| GET / POST | /checkout_service/confirm            | Get an overview of the order.                                        |
| PUT        | /checkout_service/confirm/success    | Webhook to be hit when order is successfully completed with payment. |

| ![Checkout Service REST Swagger](https://github.com/swiggy-2022-bootcamp/cdp-team3/blob/main/checkout-service/screenshots/swagger_overview.png?raw=true) |
| ---------- |

## GRPC Connections

1. Checkout Service -> Cart Service (:8010): To send Checkout confirm signal to clear User's Cart.

2. Checkout Service <- Cart Service (:8010): To receive the items in the User's Cart to validate and create Order.

3. Checkout Service -> Shipping Service (:8003): To receive and validate user's Shipping Address before creating Order.

4. Checkout Service -> Order Service (:8004): To create Order if everything is proper in checkout.

## Sonarqube Quality Check

| ![Sonarqube Quality Dashboard](https://github.com/swiggy-2022-bootcamp/cdp-team3/blob/main/checkout-service/screenshots/sonarqube_dashboard.png?raw=true) |
| ---------- |
