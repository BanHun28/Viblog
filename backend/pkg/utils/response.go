package utils

import "time"

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success   bool                   `json:"success"`
	Error     string                 `json:"error"`
	Code      string                 `json:"code"`
	Details   map[string]interface{} `json:"details,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Success:   true,
		Data:      data,
		Message:   message,
		Timestamp: time.Now(),
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(error, code string, details map[string]interface{}) *ErrorResponse {
	return &ErrorResponse{
		Success:   false,
		Error:     error,
		Code:      code,
		Details:   details,
		Timestamp: time.Now(),
	}
}
