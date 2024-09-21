package errors

import (
	"net/http"
)

// AppError represents an application-specific error with a message and status code
type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Err        error  `json:"-"` // Underlying error, if any (optional)
}

// Error implements the error interface for AppError
func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

// GetStatusCode returns the HTTP status code for the error
func (e *AppError) GetStatusCode() int {
	return e.StatusCode
}

// Unwrap returns the underlying error, if any, to allow Go's error unwrapping
func (e *AppError) Unwrap() error {
	return e.Err
}

// Constructors for different error types

// NewBadRequestError creates a new AppError with 400 Bad Request status
func NewBadRequestError(message string, err error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Err:        err,
	}
}

// NewNotFoundError creates a new AppError with 404 Not Found status
func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Err:        err,
	}
}

// NewInternalServerError creates a new AppError with 500 Internal Server Error status
func NewInternalServerError(message string, err error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
}

// NewUnauthorizedError creates a new AppError with 401 Unauthorized status
func NewUnauthorizedError(message string, err error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Err:        err,
	}
}
