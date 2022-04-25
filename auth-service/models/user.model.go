package models

import "time"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

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
	Status    string    `json:"status"`
	DateAdded time.Time `json:"date_added"`
	IsAdmin   bool      `json:"isAdmin"`
}

type Customer struct {
	CustomerId      string    `json:"customerId"`
	Firstname       string    `json:"firstname"`
	Lastname        string    `json:"lastname"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirmpassword"`
	Email           string    `json:"email"`
	Telephone       string    `json:"telephone"`
	Address         Address   `json:"address"`
	Status          string    `json:"status"`
	Approved        string    `json:"approved"`
	DateAdded       time.Time `json:"dateAdded"`
	Rewards         string    `json:"rewards"`
	IsAdmin         bool      `json:"isAdmin"`
}

var UserTableName = "users"
var AdminTableName = "Admins"
var CustomerTableName = "Customers"
