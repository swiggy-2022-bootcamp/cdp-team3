package models

import (
	//"time"
)

type ShippingAddress struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	City      string `json:"city" validate:"required"`
	Address1  string `json:"address_1" validate:"required"`
	Address2  string `json:"address_2"`
	CountryID int    `json:"country_id" validate:"required"`
	PostCode  int    `json:"postcode" validate:"required"`
	
}