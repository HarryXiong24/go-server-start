package errors

import (
	"fmt"
	"net/http"
)

// AppError represents an application error
type AppError struct {
	Code       int    // HTTP status code
	Message    string // User-friendly error message
	Err        error  // Original error
	StatusCode int    // HTTP status code (for compatibility)
}

// Error returns the error message
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError
func New(code int, message string, err error) *AppError {
	return &AppError{
		Code:       code,
		StatusCode: code, // Set both for compatibility
		Message:    message,
		Err:        err,
	}
}

// NewBadRequest creates a new bad request error
func NewBadRequest(message string, err error) *AppError {
	return New(http.StatusBadRequest, message, err)
}

// NewUnauthorized creates a new unauthorized error
func NewUnauthorized(message string, err error) *AppError {
	return New(http.StatusUnauthorized, message, err)
}

// NewForbidden creates a new forbidden error
func NewForbidden(message string, err error) *AppError {
	return New(http.StatusForbidden, message, err)
}

// NewNotFound creates a new not found error
func NewNotFound(message string, err error) *AppError {
	return New(http.StatusNotFound, message, err)
}

// NewInternalServer creates a new internal server error
func NewInternalServer(message string, err error) *AppError {
	return New(http.StatusInternalServerError, message, err)
}

// Common errors
var (
	ErrBadRequest     = &AppError{Code: http.StatusBadRequest, StatusCode: http.StatusBadRequest, Message: "Bad Request"}
	ErrUnauthorized   = &AppError{Code: http.StatusUnauthorized, StatusCode: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden      = &AppError{Code: http.StatusForbidden, StatusCode: http.StatusForbidden, Message: "Forbidden"}
	ErrNotFound       = &AppError{Code: http.StatusNotFound, StatusCode: http.StatusNotFound, Message: "Not Found"}
	ErrInternalServer = &AppError{Code: http.StatusInternalServerError, StatusCode: http.StatusInternalServerError, Message: "Internal Server Error"}
)
