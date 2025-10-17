package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestXSSProtectionMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.Use(XSSProtectionMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req := httptest.NewRequest("GET", "/test", nil)
	c.Request = req
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	xssHeader := w.Header().Get("X-XSS-Protection")
	if xssHeader != "1; mode=block" {
		t.Errorf("Expected X-XSS-Protection header '1; mode=block', got '%s'", xssHeader)
	}

	contentTypeHeader := w.Header().Get("X-Content-Type-Options")
	if contentTypeHeader != "nosniff" {
		t.Errorf("Expected X-Content-Type-Options header 'nosniff', got '%s'", contentTypeHeader)
	}
}

func TestCSRFTokenStore(t *testing.T) {
	store := NewCSRFTokenStore()

	token := "test-token-123"
	expiration := time.Now().Add(1 * time.Hour)

	// Add token
	store.Add(token, expiration)

	// Verify valid token
	if !store.Verify(token) {
		t.Error("Valid token should be verified")
	}

	// Verify invalid token
	if store.Verify("invalid-token") {
		t.Error("Invalid token should not be verified")
	}

	// Remove token
	store.Remove(token)

	// Verify removed token
	if store.Verify(token) {
		t.Error("Removed token should not be verified")
	}
}

func TestCSRFTokenStore_ExpiredToken(t *testing.T) {
	store := NewCSRFTokenStore()

	token := "expired-token"
	expiration := time.Now().Add(-1 * time.Hour) // Already expired

	store.Add(token, expiration)

	if store.Verify(token) {
		t.Error("Expired token should not be verified")
	}
}

func TestCSRFMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := DefaultCSRFConfig()
	store := NewCSRFTokenStore()

	tests := []struct {
		name           string
		method         string
		token          string
		addToken       bool
		expectedStatus int
	}{
		{
			name:           "GET request (safe method)",
			method:         "GET",
			token:          "",
			addToken:       false,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST with valid token",
			method:         "POST",
			token:          "valid-token",
			addToken:       true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST without token",
			method:         "POST",
			token:          "",
			addToken:       false,
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "POST with invalid token",
			method:         "POST",
			token:          "invalid-token",
			addToken:       false,
			expectedStatus: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			if tt.addToken {
				store.Add(tt.token, time.Now().Add(1*time.Hour))
			}

			r.Use(CSRFMiddleware(config, store))
			r.POST("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			})
			r.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			})

			req := httptest.NewRequest(tt.method, "/test", nil)
			if tt.token != "" {
				req.Header.Set(config.HeaderName, tt.token)
			}

			c.Request = req
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestGenerateCSRFToken(t *testing.T) {
	store := NewCSRFTokenStore()

	token := GenerateCSRFToken(store, 1*time.Hour)

	if token == "" {
		t.Error("Token should not be empty")
	}

	if !store.Verify(token) {
		t.Error("Generated token should be valid")
	}
}

func TestSecurityHeadersMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.Use(SecurityHeadersMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req := httptest.NewRequest("GET", "/test", nil)
	c.Request = req
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Check security headers
	expectedHeaders := map[string]string{
		"X-Frame-Options":         "DENY",
		"X-XSS-Protection":        "1; mode=block",
		"X-Content-Type-Options":  "nosniff",
		"Referrer-Policy":         "strict-origin-when-cross-origin",
		"Content-Security-Policy": "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline';",
		"Permissions-Policy":      "geolocation=(), microphone=(), camera=()",
	}

	for header, expected := range expectedHeaders {
		actual := w.Header().Get(header)
		if actual != expected {
			t.Errorf("Header %s: expected '%s', got '%s'", header, expected, actual)
		}
	}
}

func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Script tag",
			input:    "<script>alert('xss')</script>",
			expected: "&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;",
		},
		{
			name:     "HTML entities",
			input:    "<div>Hello & goodbye</div>",
			expected: "&lt;div&gt;Hello &amp; goodbye&lt;/div&gt;",
		},
		{
			name:     "Normal text",
			input:    "Hello World",
			expected: "Hello World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeInput(tt.input)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestDefaultCSRFConfig(t *testing.T) {
	config := DefaultCSRFConfig()

	if config.TokenName != "csrf_token" {
		t.Errorf("Expected token name 'csrf_token', got '%s'", config.TokenName)
	}

	if config.HeaderName != "X-CSRF-Token" {
		t.Errorf("Expected header name 'X-CSRF-Token', got '%s'", config.HeaderName)
	}

	if config.Expiration != 24*time.Hour {
		t.Errorf("Expected expiration 24 hours, got %v", config.Expiration)
	}

	if config.Secret == "" {
		t.Error("Secret should not be empty")
	}
}

func TestSecureMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := DefaultCSRFConfig()
	store := NewCSRFTokenStore()

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.Use(SecureMiddleware(config, store))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req := httptest.NewRequest("GET", "/test", nil)
	c.Request = req
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify security headers are set
	if w.Header().Get("X-Frame-Options") == "" {
		t.Error("Security headers should be set")
	}
}
