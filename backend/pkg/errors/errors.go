package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// AppError represents a structured application error
type AppError struct {
	Code       ErrorCode              `json:"code"`
	Message    string                 `json:"message"`
	StatusCode int                    `json:"-"`
	Details    map[string]interface{} `json:"details,omitempty"`
	Err        error                  `json:"-"`
}

// Error implements the error interface
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

// WithDetails adds details to the error
func (e *AppError) WithDetails(details map[string]interface{}) *AppError {
	e.Details = details
	return e
}

// WithError wraps an underlying error
func (e *AppError) WithError(err error) *AppError {
	e.Err = err
	return e
}

// New creates a new AppError
func New(code ErrorCode, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

// Predefined errors
var (
	// Authentication & Authorization
	ErrUnauthorized = New(ErrCodeUnauthorized, "Unauthorized", http.StatusUnauthorized)
	ErrForbidden    = New(ErrCodeForbidden, "Forbidden", http.StatusForbidden)
	ErrInvalidToken = New(ErrCodeInvalidToken, "Invalid token", http.StatusUnauthorized)
	ErrExpiredToken = New(ErrCodeExpiredToken, "Token expired", http.StatusUnauthorized)
	ErrInvalidPassword = New(ErrCodeInvalidPassword, "Invalid password", http.StatusUnauthorized)

	// Validation
	ErrValidation      = New(ErrCodeValidation, "Validation failed", http.StatusBadRequest)
	ErrInvalidEmail    = New(ErrCodeInvalidEmail, "Invalid email format", http.StatusBadRequest)
	ErrInvalidURL      = New(ErrCodeInvalidURL, "Invalid URL format", http.StatusBadRequest)
	ErrPasswordTooWeak = New(ErrCodePasswordTooWeak, "Password does not meet requirements (min 8 chars, letters+numbers+special chars)", http.StatusBadRequest)
	ErrInvalidInput    = New(ErrCodeInvalidInput, "Invalid input", http.StatusBadRequest)

	// Resource
	ErrNotFound      = New(ErrCodeNotFound, "Resource not found", http.StatusNotFound)
	ErrAlreadyExists = New(ErrCodeAlreadyExists, "Resource already exists", http.StatusConflict)
	ErrConflict      = New(ErrCodeConflict, "Resource conflict", http.StatusConflict)

	// Rate Limiting
	ErrRateLimitExceeded = New(ErrCodeRateLimitExceeded, "Rate limit exceeded", http.StatusTooManyRequests)
	ErrTooManyRequests   = New(ErrCodeTooManyRequests, "Too many requests", http.StatusTooManyRequests)

	// Internal
	ErrInternal      = New(ErrCodeInternal, "Internal server error", http.StatusInternalServerError)
	ErrDatabaseError = New(ErrCodeDatabaseError, "Database error", http.StatusInternalServerError)
	ErrCacheError    = New(ErrCodeCacheError, "Cache error", http.StatusInternalServerError)

	// Business Logic
	ErrInvalidOperation       = New(ErrCodeInvalidOperation, "Invalid operation", http.StatusBadRequest)
	ErrInsufficientPermission = New(ErrCodeInsufficientPermission, "Insufficient permission", http.StatusForbidden)
)

// Is checks if the error is of a specific type
func Is(err error, target *AppError) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == target.Code
	}
	return false
}

// GetStatusCode returns the HTTP status code for an error
func GetStatusCode(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.StatusCode
	}
	return http.StatusInternalServerError
}
