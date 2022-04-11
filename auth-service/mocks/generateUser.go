package main

import (
	"fmt"

	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUser() {
	user := models.User{
		Id:       primitive.NewObjectID(),
		Name:     "John Doe",
		Email:    "john@gmail.com",
		Password: utils.HashPassword("123456"),
		IsAdmin:  false,
	}
	result, err := models.UserCollection.InsertOne(nil, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func main() {
	GenerateUser()
}
