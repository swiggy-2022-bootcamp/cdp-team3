# checkout-service

Contains the Checkout Microservice for the final Swiggy I++ E-Commerce Application project created to showcase the skills learnt throughout the programme.

## Relevant APIs

| HTTP Verb  | Endpoint Path                        | Description                                                          |
| ---------- | ------------------------------------ | -------------------------------------------------------------------- |
| GET        | /checkout_service/swagger/index.html | Swagger documentation                                                |
| GET / POST | /checkout_service/confirm            | Get an overview of the order.                                        |
| PUT        | /checkout_service/confirm/success    | Webhook to be hit when order is successfully completed with payment. |

## GRPC Connections

1. Checkout Service -> Checkout Service: To send Checkout confirm signal to clear user's Cart.
2. Checkout Service -> Shipping Service: To send request to receive user's Shipping Address.
