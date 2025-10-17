package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestErrorHandlerMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		errorType      gin.ErrorType
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Bind error",
			errorType:      gin.ErrorTypeBind,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid request",
		},
		{
			name:           "Public error",
			errorType:      gin.ErrorTypePublic,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Private error",
			errorType:      gin.ErrorTypePrivate,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			r.Use(ErrorHandlerMiddleware())
			r.GET("/test", func(c *gin.Context) {
				c.Error(&gin.Error{
					Err:  http.ErrAbortHandler,
					Type: tt.errorType,
				})
			})

			req := httptest.NewRequest("GET", "/test", nil)
			c.Request = req
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response ErrorResponse
			json.Unmarshal(w.Body.Bytes(), &response)

			if tt.expectedError != "" && response.Error != tt.expectedError {
				t.Errorf("Expected error '%s', got '%s'", tt.expectedError, response.Error)
			}
		})
	}
}

func TestNotFoundHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.NoRoute(NotFoundHandler())

	req := httptest.NewRequest("GET", "/notfound", nil)
	c.Request = req
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}

	var response ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Error != "Not found" {
		t.Errorf("Expected error 'Not found', got '%s'", response.Error)
	}

	if response.Code != "NOT_FOUND" {
		t.Errorf("Expected code 'NOT_FOUND', got '%s'", response.Code)
	}
}

func TestMethodNotAllowedHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	handler := MethodNotAllowedHandler()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Manually invoke the handler
	req := httptest.NewRequest("POST", "/test", nil)
	c.Request = req
	handler(c)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}

	var response ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Error != "Method not allowed" {
		t.Errorf("Expected error 'Method not allowed', got '%s'", response.Error)
	}
}

func TestNewErrorResponse(t *testing.T) {
	response := NewErrorResponse("Test Error", "Test message", "TEST_CODE")

	if response.Error != "Test Error" {
		t.Errorf("Expected error 'Test Error', got '%s'", response.Error)
	}

	if response.Message != "Test message" {
		t.Errorf("Expected message 'Test message', got '%s'", response.Message)
	}

	if response.Code != "TEST_CODE" {
		t.Errorf("Expected code 'TEST_CODE', got '%s'", response.Code)
	}
}

func TestNewValidationErrorResponse(t *testing.T) {
	fields := map[string]string{
		"email":    "Invalid email format",
		"password": "Password too short",
	}

	response := NewValidationErrorResponse(fields)

	if response.Error != "Validation failed" {
		t.Errorf("Expected error 'Validation failed', got '%s'", response.Error)
	}

	if len(response.Fields) != 2 {
		t.Errorf("Expected 2 fields, got %d", len(response.Fields))
	}

	if response.Fields["email"] != "Invalid email format" {
		t.Errorf("Expected email error, got '%s'", response.Fields["email"])
	}
}

func TestAbortWithError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/test", func(c *gin.Context) {
		AbortWithError(c, http.StatusBadRequest, "Bad Request", "Invalid input")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	c.Request = req
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Error != "Bad Request" {
		t.Errorf("Expected error 'Bad Request', got '%s'", response.Error)
	}
}

func TestAbortWithValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.POST("/test", func(c *gin.Context) {
		fields := map[string]string{
			"username": "Username is required",
		}
		AbortWithValidationError(c, fields)
	})

	req := httptest.NewRequest("POST", "/test", nil)
	c.Request = req
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response ValidationErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Error != "Validation failed" {
		t.Errorf("Expected error 'Validation failed', got '%s'", response.Error)
	}
}
