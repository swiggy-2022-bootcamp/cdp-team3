package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reward struct {
	ID      primitive.ObjectID `json:"_id"				bson:"_id"`
	Message string             `json:"message"			bson:"message"			validate:"required"`
	Rewards string             `json:"rewards"			bson:"rewards"			validate:"required"`
}
