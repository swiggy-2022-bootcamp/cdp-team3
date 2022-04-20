package main

import (
	"fmt"

	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
)

func GenerateUser() {
	user := models.User{
		Id:       "1",
		Name:     "John Doe",
		Email:    "john@gmail.com",
		Password: utils.HashPassword("123456"),
		IsAdmin:  false,
	}
	dbUser, err := models.CreateUser(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbUser)
}
