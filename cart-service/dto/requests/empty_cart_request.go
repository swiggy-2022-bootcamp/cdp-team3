package requests

// Request DTO for emptying a Cart. Must either provide Cart ID or User ID, but not both.
type EmptyCartRequest struct {
	CartID string `json:"cart_id"` // ID of the cart
	UserID string `json:"user_id"` // ID of the user
}