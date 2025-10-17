package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(ErrCodeNotFound, "test error", http.StatusNotFound)
	
	if err.Code != ErrCodeNotFound {
		t.Errorf("expected code %s, got %s", ErrCodeNotFound, err.Code)
	}
	
	if err.Message != "test error" {
		t.Errorf("expected message 'test error', got '%s'", err.Message)
	}
	
	if err.StatusCode != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, err.StatusCode)
	}
}

func TestAppError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *AppError
		expected string
	}{
		{
			name:     "without wrapped error",
			err:      New(ErrCodeNotFound, "resource not found", http.StatusNotFound),
			expected: "resource not found",
		},
		{
			name: "with wrapped error",
			err: New(ErrCodeDatabaseError, "database error", http.StatusInternalServerError).
				WithError(errors.New("connection failed")),
			expected: "database error: connection failed",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, got)
			}
		})
	}
}

func TestAppError_WithDetails(t *testing.T) {
	err := New(ErrCodeValidation, "validation error", http.StatusBadRequest).
		WithDetails(map[string]interface{}{
			"field": "email",
			"issue": "invalid format",
		})
	
	if err.Details == nil {
		t.Error("expected details to be set")
	}
	
	if err.Details["field"] != "email" {
		t.Errorf("expected field 'email', got '%v'", err.Details["field"])
	}
}

func TestAppError_WithError(t *testing.T) {
	wrappedErr := errors.New("wrapped error")
	err := New(ErrCodeInternal, "internal error", http.StatusInternalServerError).
		WithError(wrappedErr)
	
	if err.Err != wrappedErr {
		t.Error("expected wrapped error to be set")
	}
	
	if !errors.Is(err, wrappedErr) {
		t.Error("expected errors.Is to return true")
	}
}

func TestIs(t *testing.T) {
	err := ErrNotFound.WithError(errors.New("test"))
	
	if !Is(err, ErrNotFound) {
		t.Error("expected Is to return true for matching error")
	}
	
	if Is(err, ErrUnauthorized) {
		t.Error("expected Is to return false for non-matching error")
	}
	
	stdErr := errors.New("standard error")
	if Is(stdErr, ErrNotFound) {
		t.Error("expected Is to return false for standard error")
	}
}

func TestGetStatusCode(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{
			name:     "AppError",
			err:      ErrNotFound,
			expected: http.StatusNotFound,
		},
		{
			name:     "wrapped AppError",
			err:      ErrUnauthorized.WithError(errors.New("test")),
			expected: http.StatusUnauthorized,
		},
		{
			name:     "standard error",
			err:      errors.New("test"),
			expected: http.StatusInternalServerError,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStatusCode(tt.err); got != tt.expected {
				t.Errorf("expected status %d, got %d", tt.expected, got)
			}
		})
	}
}

func TestPredefinedErrors(t *testing.T) {
	tests := []struct {
		name       string
		err        *AppError
		statusCode int
		code       ErrorCode
	}{
		{"Unauthorized", ErrUnauthorized, http.StatusUnauthorized, ErrCodeUnauthorized},
		{"Forbidden", ErrForbidden, http.StatusForbidden, ErrCodeForbidden},
		{"NotFound", ErrNotFound, http.StatusNotFound, ErrCodeNotFound},
		{"Validation", ErrValidation, http.StatusBadRequest, ErrCodeValidation},
		{"RateLimitExceeded", ErrRateLimitExceeded, http.StatusTooManyRequests, ErrCodeRateLimitExceeded},
		{"Internal", ErrInternal, http.StatusInternalServerError, ErrCodeInternal},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.StatusCode != tt.statusCode {
				t.Errorf("expected status %d, got %d", tt.statusCode, tt.err.StatusCode)
			}
			if tt.err.Code != tt.code {
				t.Errorf("expected code %s, got %s", tt.code, tt.err.Code)
			}
		})
	}
}
