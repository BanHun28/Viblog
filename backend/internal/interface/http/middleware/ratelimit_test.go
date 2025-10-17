package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRateLimiter_Allow(t *testing.T) {
	limiter := NewRateLimiter(3, time.Minute)

	// First 3 requests should be allowed
	for i := 0; i < 3; i++ {
		if !limiter.Allow("test-ip") {
			t.Errorf("Request %d should be allowed", i+1)
		}
	}

	// 4th request should be denied
	if limiter.Allow("test-ip") {
		t.Error("4th request should be denied")
	}
}

func TestRateLimiter_MultipleIPs(t *testing.T) {
	limiter := NewRateLimiter(2, time.Minute)

	// IP 1
	if !limiter.Allow("192.168.1.1") {
		t.Error("First request from IP1 should be allowed")
	}
	if !limiter.Allow("192.168.1.1") {
		t.Error("Second request from IP1 should be allowed")
	}

	// IP 2
	if !limiter.Allow("192.168.1.2") {
		t.Error("First request from IP2 should be allowed")
	}
	if !limiter.Allow("192.168.1.2") {
		t.Error("Second request from IP2 should be allowed")
	}

	// Both IPs should be rate limited now
	if limiter.Allow("192.168.1.1") {
		t.Error("IP1 should be rate limited")
	}
	if limiter.Allow("192.168.1.2") {
		t.Error("IP2 should be rate limited")
	}
}

func TestRateLimitMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestCount   int
		expectedStatus []int
	}{
		{
			name:         "Within limit",
			requestCount: 2,
			expectedStatus: []int{
				http.StatusOK,
				http.StatusOK,
			},
		},
		{
			name:         "Exceed limit",
			requestCount: 3,
			expectedStatus: []int{
				http.StatusOK,
				http.StatusOK,
				http.StatusTooManyRequests,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create new limiter for each test
			testLimiter := NewRateLimiter(2, time.Minute)

			for i := 0; i < tt.requestCount; i++ {
				w := httptest.NewRecorder()
				c, r := gin.CreateTestContext(w)

				r.Use(RateLimitMiddleware(testLimiter))
				r.GET("/test", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "success"})
				})

				req := httptest.NewRequest("GET", "/test", nil)
				req.RemoteAddr = "192.168.1.100:1234"

				c.Request = req
				r.ServeHTTP(w, req)

				if w.Code != tt.expectedStatus[i] {
					t.Errorf("Request %d: Expected status %d, got %d", i+1, tt.expectedStatus[i], w.Code)
				}
			}
		})
	}
}

func TestDefaultRateLimitConfig(t *testing.T) {
	config := DefaultRateLimitConfig()

	if config.APIRate != 100 {
		t.Errorf("Expected API rate 100, got %d", config.APIRate)
	}

	if config.CommentRate != 5 {
		t.Errorf("Expected comment rate 5, got %d", config.CommentRate)
	}

	if config.Window != time.Minute {
		t.Errorf("Expected window 1 minute, got %v", config.Window)
	}
}

func TestNewAPIRateLimiter(t *testing.T) {
	limiter := NewAPIRateLimiter()

	if limiter == nil {
		t.Error("Expected limiter to be created")
	}

	if limiter.rate != 100 {
		t.Errorf("Expected rate 100, got %d", limiter.rate)
	}
}

func TestNewCommentRateLimiter(t *testing.T) {
	limiter := NewCommentRateLimiter()

	if limiter == nil {
		t.Error("Expected limiter to be created")
	}

	if limiter.rate != 5 {
		t.Errorf("Expected rate 5, got %d", limiter.rate)
	}
}

func TestRateLimiter_TokenRefill(t *testing.T) {
	// Create a limiter with very short window for testing
	limiter := NewRateLimiter(2, 100*time.Millisecond)

	// Use up all tokens
	limiter.Allow("test-ip")
	limiter.Allow("test-ip")

	// Should be denied
	if limiter.Allow("test-ip") {
		t.Error("Should be rate limited")
	}

	// Wait for token refill
	time.Sleep(150 * time.Millisecond)

	// Should be allowed again
	if !limiter.Allow("test-ip") {
		t.Error("Should be allowed after token refill")
	}
}
