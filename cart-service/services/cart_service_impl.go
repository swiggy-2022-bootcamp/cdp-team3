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

// CreateCartItem creates a new Cart Item
func (cs *cartServiceImpl) CreateCartItem(
	ctx context.Context,
	cartItemRequest *requests.CartItemRequest,
	userID string,
) error {
	// Attempt to fetch the Cart from DB
	cart, err := cs.cartRepository.ReadByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// Check if the item already exists in the cart
	for _, item := range cart.Items {
		if item.ProductID == cartItemRequest.ProductID {
			item.Quantity += cartItemRequest.Quantity
			err = cs.cartRepository.UpdateCartItems(ctx, cart)
			if err != nil {	
				return err
			}
			return nil
		}
	}

	// Create Cart Item
	cartItem := models.CartItem{
		ProductID: cartItemRequest.ProductID,
		Quantity:  cartItemRequest.Quantity,
	}
	cart.Items = append(cart.Items, cartItem)
	err = cs.cartRepository.UpdateCartItems(ctx, cart)
	if err != nil {
		log.Errorf("Failed to create cart item: %v", err)
		return err
	}
	return nil
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
		err  error
	)
	if cartID != "" {
		cart, err = cs.cartRepository.Read(ctx, cartID)
	} else {
		cart, err = cs.cartRepository.ReadByUserID(ctx, userID)
	}
	if err != nil {
		return nil, err
	}
	// Calculate the total price of the cart
	totalPrice := 0.0
	for _, item := range cart.Items {
		totalPrice += float64(item.Quantity) * item.Price
	}
	// Return the Cart Items
	return &responses.CartItemsResponse{
		Total: totalPrice, CartItems: cart.Items,
	}, nil
}

// UpdateCartItem updates a Cart Item
func (cs *cartServiceImpl) UpdateCartItem(
	ctx context.Context,
	cartItemRequest *requests.CartItemRequest,
	userID string,
) error {
	// Attempt to fetch the Cart from DB
	cart, err := cs.cartRepository.ReadByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// Check if the item already exists in the cart
	for i, item := range cart.Items {
		if item.ProductID == cartItemRequest.ProductID {
			cart.Items[i].Quantity = cartItemRequest.Quantity
			err = cs.cartRepository.UpdateCartItems(ctx, cart)
			if err != nil {
				return err
			}
			return nil
		}
	}

	// Create Cart Item instead as it doesn't exist
	err = cs.CreateCartItem(ctx, cartItemRequest, userID)
	return err
}

// DeleteCartItem deletes a Cart Item
func (cs *cartServiceImpl) DeleteCartItem(
	ctx context.Context,
	productID string,
	userID string,
) error {
	// Attempt to fetch the Cart from DB
	cart, err := cs.cartRepository.ReadByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// Delete the Cart Item if present
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			err = cs.cartRepository.UpdateCartItems(ctx, cart)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// EmptyCart fetches the cart identified by Cart ID or User ID and empties it
func (cs *cartServiceImpl) EmptyCart(ctx context.Context, emptyCartRequest requests.EmptyCartRequest) error {
	// Attempt to fetch the Cart from DB based on request
	var (
		cart *models.Cart
		err  error
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
