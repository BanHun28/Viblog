package utils

import (
	"sync"
	"time"
)

// RateLimiter manages rate limiting using token bucket algorithm
type RateLimiter struct {
	mu       sync.RWMutex
	visitors map[string]*visitor
	rate     int           // requests per window
	window   time.Duration // time window
	cleanup  time.Duration // cleanup interval
}

// visitor represents a single visitor's rate limit state
type visitor struct {
	tokens     int
	lastRefill time.Time
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(rate int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		window:   window,
		cleanup:  window * 2,
	}
	
	// Start cleanup goroutine
	go rl.cleanupStale()
	
	return rl
}

// Allow checks if a request from the given identifier is allowed
func (rl *RateLimiter) Allow(identifier string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	v, exists := rl.visitors[identifier]
	now := time.Now()
	
	if !exists {
		rl.visitors[identifier] = &visitor{
			tokens:     rl.rate - 1,
			lastRefill: now,
		}
		return true
	}
	
	// Refill tokens if window has passed
	if now.Sub(v.lastRefill) >= rl.window {
		v.tokens = rl.rate
		v.lastRefill = now
	}
	
	if v.tokens > 0 {
		v.tokens--
		return true
	}
	
	return false
}

// Reset resets the rate limit for a specific identifier
func (rl *RateLimiter) Reset(identifier string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.visitors, identifier)
}

// cleanupStale removes stale visitors
func (rl *RateLimiter) cleanupStale() {
	ticker := time.NewTicker(rl.cleanup)
	defer ticker.Stop()
	
	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for id, v := range rl.visitors {
			if now.Sub(v.lastRefill) > rl.cleanup {
				delete(rl.visitors, id)
			}
		}
		rl.mu.Unlock()
	}
}

// GetRemaining returns remaining tokens for an identifier
func (rl *RateLimiter) GetRemaining(identifier string) int {
	rl.mu.RLock()
	defer rl.mu.RUnlock()
	
	v, exists := rl.visitors[identifier]
	if !exists {
		return rl.rate
	}
	
	// Check if tokens should be refilled
	if time.Now().Sub(v.lastRefill) >= rl.window {
		return rl.rate
	}
	
	return v.tokens
}
