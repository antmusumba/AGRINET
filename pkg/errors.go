package pkg

import (
	"fmt"
)

// ErrorType represents the type of error that occurred
type ErrorType string

const (
	ErrorTypeValidation ErrorType = "VALIDATION"
	ErrorTypeNotFound   ErrorType = "NOT_FOUND"
	ErrorTypeDatabase   ErrorType = "DATABASE"
	ErrorTypeInternal   ErrorType = "INTERNAL"
)

// Error represents a structured error in the user package
type Error struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Field   string    `json:"field,omitempty"`
	Err     error     `json:"-"`
}

// Error implements the error interface
func (e *Error) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s", e.Field, e.Message)
	}
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap implements the errors.Unwrap interface
func (e *Error) Unwrap() error {
	return e.Err
}

// NewError creates a new error with the type and message, optionally wrapping another error
func NewError(errType ErrorType, message string) *Error {
	return &Error{
		Type:    errType,
		Message: message,
	}
}

// NewValidationError creates a new VALIDATION error
func NewValidationError(field, message string) *Error {
	return &Error{
		Type:    ErrorTypeValidation,
		Message: message,
		Field:   field,
	}
}

// WrapError wraps an error with a custom error type and message
func WrapError(err error, errType ErrorType, message string) *Error {
	// If the error is nil, return a generic internal error
	if err == nil {
		return NewError(errType, message)
	}

	// If the error is not nil, wrap the original error with additional context
	return &Error{
		Type:    errType,
		Message: message,
		Err:     err,
	}
}
