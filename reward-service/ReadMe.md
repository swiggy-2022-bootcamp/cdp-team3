### Endoints

| Http Method        | Enpoint   | Description |
| ------------- |:-------------:|:-------------:|
| GET     |/orders| Get List of Orders of all the customer from the front store 
| GET     |/orders/status/{order_status}| Get Order Details By Status
| GET     |/orders/{orderId}| Get Order details by Order ID 
| PUT     |/orders/{orderId}| Update Order Status by ID 
| DELETE  |/orders/{orderId}| Delete Order  by ID 
| GET     |/orders/user/{userId}| Get Order details by Customer ID 
| POST    |/orders/invoice/{orderId}| Generate Invoice number for the order
| GET     |/orders/{orderId}/order_status| (Front Store) Get Order Status by Order ID placed by a Customer



Run locally 
```sh
go run main.go
```

Generate swagger docs
```sh
swag init
```

Swagger page
> http://localhost:3005/swagger/index.html
