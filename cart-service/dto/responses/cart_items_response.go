package responses

import "github.com/swiggy-ipp/cart-service/models"

// Contains a Response DTO of the current items in the cart.
type CartItemsResponse struct {
	CartItems []models.CartItem `json:"items"` // List of items in the cart.
	Total     float64           `json:"total"` // Total price of the cart.
}
