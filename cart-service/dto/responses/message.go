package responses

// Basic response DTO object to be returned to the client.
type MessageResponse struct {
	Message string `json:"message" example:"Sample Message"` // Message being sent through the DTO
}
