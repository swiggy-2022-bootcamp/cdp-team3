package responses

// Contains a Response DTO of the current items in the cart.
type CartItemsResponse struct {
	CartItems []CartItem `json:"items"` // List of items in the cart.
}

// Cart Item Data Model
type CartItem struct {
	ProductID string `json:"product_id" validate:"required"` // ID of the product
	Quantity  int64  `json:"quantity" validate:"required"`   // Quantity of the product in the cart.
}
