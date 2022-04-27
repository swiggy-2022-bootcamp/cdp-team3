package models


type ShippingAddress struct {
	Id            string `json:"id" dynamodbav:"id"`
	FirstName     string `json:"firstname" validate:"required" dynamodbav:"firstname"`
	LastName      string `json:"lastname" validate:"required" dynamodbav:"lastname"`
	City          string `json:"city" validate:"required" dynamodbav:"city"`
	Address1      string `json:"address_1" validate:"required" dynamodbav:"address_1"`
	Address2      string `json:"address_2" dynamodbav:"address_2"`
	CountryID     uint32 `json:"country_id" validate:"required" dynamodbav:"country_id"`
	PostCode      uint32 `json:"postcode" validate:"required" dynamodbav:"postcode"`
	UserID        string `json:"user_id" dynamodbav:"user_id"`
	DefaultAddress string `json:"default_address" dynamodbav:"default_address"`
	
}