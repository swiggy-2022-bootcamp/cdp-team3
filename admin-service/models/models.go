package models

import "time"

type Address struct {
	ShippingAddressId string `json:"shippingAddressId"`
	Address1          string `json:"address1"`
	Address2          string `json:"address2"`
	City              string `json:"city"`
	CountryID         uint32 `json:"countryid"`
	PostCode          uint32 `json:"postcode"`
	Default           string `json:"default"`
}

type ShippingAddress struct {
	ShippingAddressId string `json:"shippingAddressId"`
	Address1          string `json:"house_number"`
	Address2          string `json:"street"`
	City              string `json:"city"`
	Countryid         int    `json:"countryId"`
	Postcode          int    `json:"postcode"`
	Default           int    `json:"default"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	UserId            string `json:"userId"`
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
