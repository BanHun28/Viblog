package utils

import (
	"testing"
	"time"
)

func TestRateLimiter_Allow(t *testing.T) {
	rl := NewRateLimiter(3, 1*time.Second)
	identifier := "user1"
	
	// Should allow first 3 requests
	for i := 0; i < 3; i++ {
		if !rl.Allow(identifier) {
			t.Errorf("request %d should be allowed", i+1)
		}
	}
	
	// Fourth request should be blocked
	if rl.Allow(identifier) {
		t.Error("fourth request should be blocked")
	}
	
	// Wait for window to reset
	time.Sleep(1100 * time.Millisecond)
	
	// Should allow requests again
	if !rl.Allow(identifier) {
		t.Error("request after window reset should be allowed")
	}
}

func TestRateLimiter_MultipleIdentifiers(t *testing.T) {
	rl := NewRateLimiter(2, 1*time.Second)
	
	user1 := "user1"
	user2 := "user2"
	
	// User1 uses their quota
	if !rl.Allow(user1) {
		t.Error("user1 first request should be allowed")
	}
	if !rl.Allow(user1) {
		t.Error("user1 second request should be allowed")
	}
	if rl.Allow(user1) {
		t.Error("user1 third request should be blocked")
	}
	
	// User2 should still have quota
	if !rl.Allow(user2) {
		t.Error("user2 first request should be allowed")
	}
	if !rl.Allow(user2) {
		t.Error("user2 second request should be allowed")
	}
}

func TestRateLimiter_Reset(t *testing.T) {
	rl := NewRateLimiter(2, 1*time.Second)
	identifier := "user1"
	
	// Use up quota
	rl.Allow(identifier)
	rl.Allow(identifier)
	
	if rl.Allow(identifier) {
		t.Error("request should be blocked before reset")
	}
	
	// Reset
	rl.Reset(identifier)
	
	// Should allow again
	if !rl.Allow(identifier) {
		t.Error("request should be allowed after reset")
	}
}

func TestRateLimiter_GetRemaining(t *testing.T) {
	rl := NewRateLimiter(5, 1*time.Second)
	identifier := "user1"
	
	// Initial state
	remaining := rl.GetRemaining(identifier)
	if remaining != 5 {
		t.Errorf("expected 5 remaining, got %d", remaining)
	}
	
	// After one request
	rl.Allow(identifier)
	remaining = rl.GetRemaining(identifier)
	if remaining != 4 {
		t.Errorf("expected 4 remaining, got %d", remaining)
	}
	
	// After three more requests
	rl.Allow(identifier)
	rl.Allow(identifier)
	rl.Allow(identifier)
	remaining = rl.GetRemaining(identifier)
	if remaining != 1 {
		t.Errorf("expected 1 remaining, got %d", remaining)
	}
}

func TestRateLimiter_ConcurrentAccess(t *testing.T) {
	rl := NewRateLimiter(100, 1*time.Second)
	identifier := "user1"
	
	done := make(chan bool)
	allowed := make(chan bool, 150)
	
	// Spawn 150 concurrent goroutines
	for i := 0; i < 150; i++ {
		go func() {
			allowed <- rl.Allow(identifier)
			done <- true
		}()
	}
	
	// Wait for all goroutines
	for i := 0; i < 150; i++ {
		<-done
	}
	close(allowed)
	
	// Count allowed requests
	allowedCount := 0
	for a := range allowed {
		if a {
			allowedCount++
		}
	}
	
	// Should allow exactly 100 (or close to it due to race conditions)
	if allowedCount > 100 {
		t.Errorf("expected max 100 allowed, got %d", allowedCount)
	}
	
	// Should have some blocked requests
	if allowedCount == 150 {
		t.Error("expected some requests to be blocked")
	}
}

func TestRateLimiter_WindowRefill(t *testing.T) {
	rl := NewRateLimiter(3, 500*time.Millisecond)
	identifier := "user1"
	
	// Use up quota
	rl.Allow(identifier)
	rl.Allow(identifier)
	rl.Allow(identifier)
	
	// Should be blocked
	if rl.Allow(identifier) {
		t.Error("should be blocked")
	}
	
	// Wait half the window
	time.Sleep(250 * time.Millisecond)
	
	// Should still be blocked
	if rl.Allow(identifier) {
		t.Error("should still be blocked")
	}
	
	// Wait for full window
	time.Sleep(300 * time.Millisecond)
	
	// Should be allowed again (tokens refilled)
	if !rl.Allow(identifier) {
		t.Error("should be allowed after window")
	}
}
