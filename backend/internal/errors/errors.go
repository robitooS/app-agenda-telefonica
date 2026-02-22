package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrAlreadyExists = errors.New("already exists")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
	ErrInternal      = errors.New("internal server error")
)

type APIError struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func NewAPIError(code, message string, details ...string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func WrapErrorf(err error, format string, args ...interface{}) error {
	return fmt.Errorf(format+": %w", append(args, err)...)
}
