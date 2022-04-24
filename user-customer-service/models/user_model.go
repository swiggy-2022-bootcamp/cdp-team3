package models

import "time"

type User struct {
	CustomerId        string    `json:"customerId" dynamodbav:"customerId"`
	Firstname         string    `json:"firstname" dynamodbav:"firstname"`
	Lastname          string    `json:"lastname" dynamodbav:"lastname"`
	Username          string    `json:"username" dynamodbav:"lastname"`
	Password          string    `json:"password" dynamodbav:"password"`
	ConfirmPassword   string    `json:"confirmpassword" dynamodbav:"confirmpassword"`
	Email             string    `json:"email" dynamodbav:"email"`
	Telephone         string    `json:"telephone" dynamodbav:"telephone"`
	AddressId         []string  `json:"addressId" dynamodbav:"addressId"`
	Status            string    `json:"status" dynamodbav:"status"`
	Approved          string    `json:"approved" dynamodbav:"approved"`
	Cart              []string  `json:"cart" dynamodbav:"cart"`
	DateAdded         time.Time `json:"dateAdded" dynamodbav:"dateAdded"`
	Rewards           string    `json:"rewards" dynamodbav:"rewards"`
	TransactionPoints int32     `json:"transactionpoints" dynamodbav:"transactionpoints"`
}
