package errors

import "net/http"

type UserError struct {
	Status       int
	ErrorMessage string
}

func (customerError *UserError) Error() string {
	return customerError.ErrorMessage
}

func NewMalformedIdError() *UserError {
	return &UserError{http.StatusBadRequest, "Malformed customer id"}
}

func NewMarshallError() *UserError {
	return &UserError{http.StatusBadRequest, "Failed to marshal the customer"}
}

func NewUserNotFoundError() *UserError {
	return &UserError{http.StatusNotFound, "User not found"}
}

func NewEmailAlreadyRegisteredError() *UserError {
	return &UserError{http.StatusBadRequest, "User with given email already exists"}
}
