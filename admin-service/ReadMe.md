### Admin Service

### Endoints

| Http Method |         Enpoint         |
| ----------- | :---------------------: |
| GET         |          /user          |
| POST        |       /customers        |
| GET         |     /customers/:id      |
| GET         | /customers/email/:email |
| PUT         |     /customers/:id      |
| DELETE      |     /customers/:id      |






Run locally 
```sh
go run main.go
```

Generate swagger docs
```sh
swag init
```

Docker Image Pushed To
```sh
immadiuttej/admin-service
```

Swagger page
> http://localhost:3009/swagger/index.html
