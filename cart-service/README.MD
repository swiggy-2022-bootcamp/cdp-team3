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

## Kafka Topics

1. `user.created.topic` -> To receive Created Users' IDs to create empty cart automatically for the created user whenever they are created.

2. `user.deleted.topic` -> To receive Deleted Users' IDs to delete their cart automatically whenever they are deleted.

## Data Models

1. Cart Data Model: cart Collection in DynamoDB

    ```json
    {
        "id": string, // Unique ID of the Cart
        "user_id": string, // ID of the User who the Cart belongs to
        "items": [
            {
                "product_id": string, // ID of the Product
                "price": float, // Price of the Product
                "quantity": int, // Quantity of the Product
            }
        ]
    }
    ```

## Database Configuration

Amazon DynamoDB is used as the database for the Cart Service. Region is set to 'us-east-1' and the table name is 'cart'.
