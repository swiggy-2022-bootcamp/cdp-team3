# cart-service

Contains the Cart Microservice for the final Swiggy I++ E-Commerce Application project created to showcase the skills learnt throughout the programme.

## Relevant APIs

| HTTP Verb | Endpoint Path                    | Description            |
| --------- | -------------------------------- | ---------------------- |
| GET       | /cart_service/swagger/index.html | Swagger documentation  |
| POST      | /cart_service/cart               | Add Cart Item          |
| GET       | /cart_service/cart               | Get Cart Items         |
| PUT       | /cart_service/cart               | Update Cart Item       |
| DELETE    | /cart_service/cart/{key}         | Delete Cart Item by ID |
| DELETE    | /cart_service/cart/empty         | Empty entire Cart      |

## GRPC Connections

1. Cart Service <- Checkout Service: To receive Checkout confirm signal to clear user's Cart.
