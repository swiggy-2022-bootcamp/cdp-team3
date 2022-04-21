package repositories

import (
	"context"

	"github.com/swiggy-ipp/cart-service/models"
)

// CRUD Repository for Cart Collection in DynamoDB
type CartRepository interface {
	// Create creates a new cart
	Create(ctx context.Context, cart *models.Cart) error
	// Read reads a cart by its ID
	Read(ctx context.Context, id string) (*models.Cart, error)
	// Read by userID reads a cart by its associated userID
	ReadByUserID(ctx context.Context, userID string) (*models.Cart, error)
	// UpdateCartItems updates a cart
	UpdateCartItems(ctx context.Context, cart *models.Cart) error
	// Delete deletes a cart
	Delete(ctx context.Context, id string) error
}
