package errors

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) Error() error {
	return errors.New(e.Message)
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewConflictRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusConflict,
	}
}

func NewExpectationFailed(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusExpectationFailed,
	}
}

func UserAlreadyExists(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusConflict,
	}
}

func ParseFail(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusFailedDependency,
	}
}

func UserNotFound(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}
