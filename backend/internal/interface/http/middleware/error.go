package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Error   string                 `json:"error"`
	Message string                 `json:"message,omitempty"`
	Code    string                 `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ErrorHandler handles errors in a standardized way
func ErrorHandler() gin.HandlerFunc {
	return ErrorHandlerMiddleware()
}

// ErrorHandlerMiddleware handles errors in a standardized way
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Default error response
			response := ErrorResponse{
				Error:   "Internal server error",
				Message: err.Error(),
			}

			// Set status code based on error type
			statusCode := http.StatusInternalServerError

			// Check for custom error types (you can expand this)
			switch err.Type {
			case gin.ErrorTypeBind:
				statusCode = http.StatusBadRequest
				response.Error = "Invalid request"
			case gin.ErrorTypePublic:
				statusCode = http.StatusBadRequest
			case gin.ErrorTypePrivate:
				statusCode = http.StatusInternalServerError
			}

			// If status hasn't been written yet, write it
			if c.Writer.Status() == http.StatusOK {
				c.JSON(statusCode, response)
			}
		}
	}
}

// ValidationErrorResponse represents validation error details
type ValidationErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// NotFoundHandler returns a 404 error response
func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "Not found",
			Message: "The requested resource was not found",
			Code:    "NOT_FOUND",
		})
	}
}

// MethodNotAllowedHandler returns a 405 error response
func MethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, ErrorResponse{
			Error:   "Method not allowed",
			Message: "The HTTP method is not allowed for this endpoint",
			Code:    "METHOD_NOT_ALLOWED",
		})
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(error string, message string, code string) ErrorResponse {
	return ErrorResponse{
		Error:   error,
		Message: message,
		Code:    code,
	}
}

// NewValidationErrorResponse creates a validation error response
func NewValidationErrorResponse(fields map[string]string) ValidationErrorResponse {
	return ValidationErrorResponse{
		Error:  "Validation failed",
		Fields: fields,
	}
}

// AbortWithError is a helper to abort with a standardized error
func AbortWithError(c *gin.Context, statusCode int, error string, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Error:   error,
		Message: message,
	})
}

// AbortWithErrorCode is a helper to abort with error code
func AbortWithErrorCode(c *gin.Context, statusCode int, error string, message string, code string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Error:   error,
		Message: message,
		Code:    code,
	})
}

// AbortWithValidationError is a helper to abort with validation errors
func AbortWithValidationError(c *gin.Context, fields map[string]string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, NewValidationErrorResponse(fields))
}
