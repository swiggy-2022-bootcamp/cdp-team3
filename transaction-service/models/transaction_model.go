package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Amount     string             `json:"amount,omitempty" validate:"required"`
	Decription string             `json:"description,omitempty"`
}
