### Admin Service

### Endoints

| Http Method |          Enpoint          |              Description              |
| ----------- | :-----------------------: | :-----------------------------------: |
| POST        |          /admin           |          To Create an Admin           |
| GET         |        /admin/user        | To get the details of logged in Admin |
| POST        |        /customers         |         To create a customer          |
| GET         |  /customers/:customerId   |  To Get the customer by customer ID   |
| GET         | /customers/email/:emailId |    To get the customer by email ID    |
| PUT         |  /customers/:customerId   | To update the customer by customer Id |
| DELETE      |  /customers/:customerId   | To delete the customer by customer Id |





Run locally 
```sh
go run main.go
```

Generate swagger docs
```sh
swag init
```

Swagger page
> http://localhost:3009/swagger/index.html