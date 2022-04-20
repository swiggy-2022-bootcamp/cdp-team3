package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cart Data Model
type Cart struct {
	Id     string     `json:"id" validate:"required"`      // DB ID of the cart.
	UserID string     `json:"user_id" validate:"required"` // ID of the user that the cart belongs to.
	Items  []CartItem `json:"items" validate:"required"`   // Items in the User's cart
}

// Helper method to generate a key for the Data Model
func (c *Cart) GenerateKey() {
	c.Id = primitive.NewObjectID().Hex()
}

// Serialize Cart for DynamoDB and return a map[string]types.AttributeValue
// Returns an error if the serialization fails.
func (c *Cart) Marshal() (map[string]types.AttributeValue, error) {
	return attributevalue.MarshalMapWithOptions(c, encoderJSON)
}

// Unserialize data from DynamoDB into the models.Cart object calling this.
// Returns an error if the deserialization fails.
func (c *Cart) UnMarshal(data map[string]types.AttributeValue) error {
	return attributevalue.UnmarshalMapWithOptions(data, &c, decoderJSON)
}

// Cart Item Data Model
type CartItem struct {
	ProductID string `json:"product_id" validate:"required"` // ID of the product
	Quantity  string `json:"quantity" validate:"required"`   // Quantity of the product in the cart.
}
