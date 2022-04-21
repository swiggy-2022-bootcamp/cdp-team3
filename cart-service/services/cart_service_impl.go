package services

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/cart-service/dto/requests"
	"github.com/swiggy-ipp/cart-service/dto/responses"
	"github.com/swiggy-ipp/cart-service/models"
	"github.com/swiggy-ipp/cart-service/repositories"
)

// Service Layer Business Logic implementation for Cart Service
type cartServiceImpl struct {
	cartRepository repositories.CartRepository
}

// Create a new Cart Service
func NewCartService(cartRepository repositories.CartRepository) CartService {
	return &cartServiceImpl{cartRepository: cartRepository}
}

// GetCartItems fetches the cart items from DB
func (cs *cartServiceImpl) GetCartItems(
	ctx context.Context, 
	cartID string, 
	userID string,
) (*responses.CartItemsResponse, error) {
	// Attempt to fetch the Cart from DB based on request
	var (
		cart *models.Cart
	 	err error
	)
	if cartID != "" {
		cart, err = cs.cartRepository.Read(ctx, cartID)
	} else {
		cart, err = cs.cartRepository.ReadByUserID(ctx, userID)
	}
	if err != nil {
		return nil, err
	}
	// Return the Cart Items
	return &responses.CartItemsResponse{CartItems: cart.Items}, nil
}


// EmptyCart fetches the cart identified by Cart ID or User ID and empties it
func (cs *cartServiceImpl) EmptyCart(ctx context.Context, emptyCartRequest requests.EmptyCartRequest) error {
	// Attempt to fetch the Cart from DB based on request
	var (
		cart *models.Cart
	 	err error
	)
	if emptyCartRequest.CartID != "" {
		cart, err = cs.cartRepository.Read(ctx, emptyCartRequest.CartID)
	} else {
		cart, err = cs.cartRepository.ReadByUserID(ctx, emptyCartRequest.UserID)
	}
	// Check for errors
	if err != nil {
		return err
	}

	// Empty Cart Items
	cart.Items = []models.CartItem{}
	err = cs.cartRepository.UpdateCartItems(ctx, cart)
	if err != nil {
		log.Errorf("Failed to empty cart: %v", err)
		return err
	}
	return nil
}

// DBHealthCheck checks the health of the DB
func (cs *cartServiceImpl) DBHealthCheck(ctx context.Context) error {
	_, err := cs.cartRepository.Read(ctx, "")
	return err
}
