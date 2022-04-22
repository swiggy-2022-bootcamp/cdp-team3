package requests

type CartItemRequest struct {
	ProductID string `json:"product_id" validate:"required"` // ID of the product
	Quantity  int64 `json:"quantity" validate:"required"`   // Quantity of the product in the cart.
}
