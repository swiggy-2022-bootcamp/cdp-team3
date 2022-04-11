package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	Id        primitive.ObjectID `json:"id" validate:"required"` // MongoDB ID of the cart item
	ProductID string             `json:"product_id" validate:"required"` // ID of the product
	Quantity  string             `json:"quantity" validate:"required"` // Quantity of the product in the cart.
}
