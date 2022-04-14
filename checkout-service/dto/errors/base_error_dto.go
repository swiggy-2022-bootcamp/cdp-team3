package errors

// Create a new Error.
func NewHTTPErrorDTO(statusCode int, err error, details ...string) HTTPErrorDTO {
	var errMessage string
	if err != nil {
		errMessage = err.Error()
	}
	return HTTPErrorDTO{
		Code:    statusCode,
		Message: errMessage,
		Details: details,
	}
}

// HTTPError example
type HTTPErrorDTO struct {
	Code    int    `json:"code" example:"400"` // Error code
	Message string `json:"message,omitempty" example:"Status bad request."` // Brief Error Message
	Details []string `json:"details,omitempty" example:"Invalid email."` // Detailed Error Messages
}
