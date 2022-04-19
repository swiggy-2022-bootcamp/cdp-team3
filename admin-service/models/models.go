package models

import (
	"time"
)

type Address struct {
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Pincode     string `json:"pincode"`
	Default     int    `json:"default"`
}

type Admin struct {
	UserId      string    `json:"userId"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Telephone   string    `json:"telephone"`
	UserGroupId int       `json:"user_group_id"`
	UserGroup   string    `json:"user_group"`
	Status      int       `json:"status"`
	DateAdded   time.Time `json:"date_added"`
}

type Customer struct {
	CustomerId      string    `json:"customerId"		bson:"customerId"`
	UserGroup       string    `json:"userGroup"		bson:"userGroup"`
	Firstname       string    `json:"firstname"		bson:"Email""`
	Lastname        string    `json:"lastname"			bson:"lastname""`
	Username        string    `json:"username"			bson:"lastname`
	Password        string    `json:"password"			bson:"password""`
	ConfirmPassword string    `json:"confirmpassword"	bson:"confirmpassword""`
	Email           string    `json:"email"			bson:"email""`
	Telephone       string    `json:"telephone"		bson:"telephone"`
	Address         []Address `json:"address"			bson:"address"`
	Status          string    `json:"status"			bson:"status"`
	Approved        string    `json:"approved"			bson:"approved"`
	Cart            []string  `json:"cart"				bson:"cart"`
	DateAdded       time.Time `json:"dateAdded"		bson:"dateAdded"`
	Rewards         string    `json:"rewards"			bson:"rewards"`
}
