package models

import "time"

type Address struct {
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Pincode     string `json:"pincode"`
	Default     int    `json:"default"`
}

type Admin struct {
	AdminId   string    `json:"adminId"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Telephone string    `json:"telephone"`
	IsAdmin   bool      `json:"isAdmin"`
	Status    string    `json:"status"`
	DateAdded time.Time `json:"date_added"`
}

type SwaggerAdmin struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
}

type Customer struct {
	CustomerId        string    `json:"customerId"		bson:"customerId"`
	IsAdmin           bool      `json:"isAdmin"			bson:"isAdmin"`
	Firstname         string    `json:"firstname"			bson:"firstname""`
	Lastname          string    `json:"lastname"			bson:"lastname""`
	Username          string    `json:"username"			bson:"username`
	Password          string    `json:"password"			bson:"password""`
	ConfirmPassword   string    `json:"confirmpassword"	bson:"confirmpassword""`
	Email             string    `json:"email"				bson:"email""`
	Telephone         string    `json:"telephone"			bson:"telephone"`
	Address           Address   `json:"address"			bson:"address"`
	Status            string    `json:"status"			bson:"status"`
	Approved          string    `json:"approved"			bson:"approved"`
	DateAdded         time.Time `json:"dateAdded"			bson:"dateAdded"`
	Rewards           int32     `json:"rewards"			bson:"rewards"`
	TransactionPoints float32   `json:"transaction_points"			bson:"transaction_points"`
}

type SwaggerCustomer struct {
	Firstname       string  `json:"firstname"		bson:"firstname""`
	Lastname        string  `json:"lastname"			bson:"lastname""`
	Username        string  `json:"username"			bson:"username`
	Password        string  `json:"password"			bson:"password""`
	ConfirmPassword string  `json:"confirmpassword"	bson:"confirmpassword""`
	Email           string  `json:"email"			bson:"email""`
	Telephone       string  `json:"telephone"		bson:"telephone"`
	Address         Address `json:"address"			bson:"address"`
}
