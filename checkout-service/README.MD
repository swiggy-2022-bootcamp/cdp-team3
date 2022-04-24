# checkout-service

Contains the Checkout Microservice for the final Swiggy I++ E-Commerce Application project created to showcase the skills learnt throughout the programme.

## Relevant APIs

| HTTP Verb  | Endpoint Path                        | Description                                                          |
| ---------- | ------------------------------------ | -------------------------------------------------------------------- |
| GET        | /checkout_service/swagger/index.html | Swagger documentation                                                |
| GET / POST | /checkout_service/confirm            | Get an overview of the order.                                        |
| PUT        | /checkout_service/confirm/success    | Webhook to be hit when order is successfully completed with payment. |

## GRPC Connections

1. Checkout Service -> Cart Service: To send Checkout confirm signal to clear User's Cart.

2. Checkout Service <- Cart Service: To receive the items in the User's Cart to validate and create Order.

3. Checkout Service -> Shipping Service: To receive and validate user's Shipping Address before creating Order.

4. Checkout Service -> Order Service: To create Order if everything is proper in checkout.
