package services

import (
	"context"

	"github.com/swiggy-ipp/cart-service/dto/requests"
	"github.com/swiggy-ipp/cart-service/dto/responses"
)

// The CartService Layer abstracting the business logic of Carts Microservice
type CartService interface{
	// GetCartItems fetches the cart items from DB
	GetCartItems(ctx context.Context, cartID string, userID string) (*responses.CartItemsResponse, error)
	// EmptyCart empties a cart
	EmptyCart(ctx context.Context, emptyCartRequest requests.EmptyCartRequest) error
	// DBHealthCheck checks the health of the DB
	DBHealthCheck(ctx context.Context) error
}
