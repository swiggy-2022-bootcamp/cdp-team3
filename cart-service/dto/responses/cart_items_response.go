package responses

import "github.com/swiggy-ipp/cart-service/models"

// Contains a Response DTO of the current items in the cart.
type CartItemsResponse struct {
	CartItems []models.CartItem `json:"cart_items"` // List of items in the cart.
}
