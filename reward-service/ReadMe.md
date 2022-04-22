### Endoints

| Http Method |        Enpoint         |                    Description                    |
| ----------- | :--------------------: | :-----------------------------------------------: |
| GET         |        /rewards        |         Get List of all the added Rewards         |
| POST        |        /rewards        |                  Add New Reward                   |
| GET         |  /rewards/{rewardId}   |          Get Reward details by Reward ID          |
| GET         | /rewards/user/{userId} | Get all the Rewards for a customer by Customer ID |




Run locally 
```sh
go run main.go
```

Generate swagger docs
```sh
swag init
```

Swagger page
> http://localhost:3008/swagger/index.html
