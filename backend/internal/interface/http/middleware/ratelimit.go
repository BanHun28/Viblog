package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter implements in-memory token bucket rate limiting
type RateLimiter struct {
	mu       sync.RWMutex
	visitors map[string]*visitor
	rate     int           // requests per window
	window   time.Duration // time window
}

type visitor struct {
	tokens     int
	lastUpdate time.Time
}

// NewRateLimiter creates a new rate limiter
// rate: maximum requests allowed
// window: time window (e.g., 1 minute)
func NewRateLimiter(rate int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		window:   window,
	}

	// Start cleanup goroutine to remove old visitors
	go rl.cleanup()

	return rl
}

// Allow checks if a request from the given identifier is allowed
func (rl *RateLimiter) Allow(identifier string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	v, exists := rl.visitors[identifier]

	if !exists {
		// New visitor
		rl.visitors[identifier] = &visitor{
			tokens:     rl.rate - 1,
			lastUpdate: now,
		}
		return true
	}

	// Calculate token refill
	elapsed := now.Sub(v.lastUpdate)
	tokensToAdd := int(elapsed / rl.window * time.Duration(rl.rate))

	v.tokens += tokensToAdd
	if v.tokens > rl.rate {
		v.tokens = rl.rate
	}
	v.lastUpdate = now

	// Check if request is allowed
	if v.tokens > 0 {
		v.tokens--
		return true
	}

	return false
}

// cleanup removes visitors that haven't been seen in 3x the window duration
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.window)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for id, v := range rl.visitors {
			if now.Sub(v.lastUpdate) > rl.window*3 {
				delete(rl.visitors, id)
			}
		}
		rl.mu.Unlock()
	}
}

// RateLimit creates a rate limiting middleware
func RateLimit(rate int, window time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(rate, window)
	return RateLimitMiddleware(limiter)
}

// RateLimitMiddleware creates a rate limiting middleware
func RateLimitMiddleware(limiter *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use IP address as identifier
		identifier := c.ClientIP()

		if !limiter.Allow(identifier) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	APIRate     int           // API requests per minute (100)
	CommentRate int           // Comment requests per minute (5)
	Window      time.Duration // Time window (1 minute)
}

// DefaultRateLimitConfig returns default rate limiting configuration
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		APIRate:     100,
		CommentRate: 5,
		Window:      time.Minute,
	}
}

// NewAPIRateLimiter creates a rate limiter for general API calls (100 req/min)
func NewAPIRateLimiter() *RateLimiter {
	config := DefaultRateLimitConfig()
	return NewRateLimiter(config.APIRate, config.Window)
}

// NewCommentRateLimiter creates a rate limiter for comments (5 req/min)
func NewCommentRateLimiter() *RateLimiter {
	config := DefaultRateLimitConfig()
	return NewRateLimiter(config.CommentRate, config.Window)
}
