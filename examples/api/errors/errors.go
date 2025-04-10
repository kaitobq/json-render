package errors

import "net/http"

// APIError represents an error response from the API
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// NewAPIError creates a new APIError
func NewAPIError(code int, message string, details ...string) *APIError {
	err := &APIError{
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

// Common API errors
var (
	ErrInvalidRequest = NewAPIError(http.StatusBadRequest, "Invalid request")
	ErrUnauthorized   = NewAPIError(http.StatusUnauthorized, "Unauthorized")
	ErrForbidden      = NewAPIError(http.StatusForbidden, "Forbidden")
	ErrNotFound       = NewAPIError(http.StatusNotFound, "Resource not found")
	ErrInternal       = NewAPIError(http.StatusInternalServerError, "Internal server error")
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents multiple validation errors
type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

// Error implements the error interface
func (e *ValidationErrors) Error() string {
	return "Validation failed"
}
