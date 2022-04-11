package responses

import "github.com/swiggy-ipp/cart-service/models"

type CartItemsResponse struct {
	CartItems []models.CartItem `json:"cart_items"` // List of items in the cart.
}
