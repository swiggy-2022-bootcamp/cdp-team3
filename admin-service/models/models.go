package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
}

type Customer struct {
	ID              primitive.ObjectID `json:"_id"				bson:"_id"`
	UserGroup       string             `json:"userGroup"		bson:"userGroup"		validate:"required"`
	FirstName       string             `json:"firstname"		bson:"Email"			validate:"required"`
	LastName        string             `json:"lastname"			bson:"lastname"			validate:"required"`
	Password        string             `json:"password"			bson:"password"			validate:"required"`
	ConfirmPassword string             `json:"confirmpassword"	bson:"confirmpassword"	validate:"required"`
	Email           string             `json:"email"			bson:"email"			validate:"required"`
	Telephone       string             `json:"telephone"		bson:"telephone"		validate:"required"`
	Status          string             `json:"status"			bson:"status"`
	Approved        string             `json:"approved"			bson:"approved"`
	Cart            []string           `json:"cart"				bson:"cart"`
	DateAdded       time.Time          `json:"dateAdded"		bson:"dateAdded"`
	Rewards         string             `json:"rewards"			bson:"rewards"`
}
