package utils

import (
	"testing"
	"time"
)

func TestCache_SetAndGet(t *testing.T) {
	cache := NewCache()
	
	key := "test_key"
	value := "test_value"
	
	cache.Set(key, value, 1*time.Second)
	
	result, ok := cache.Get(key)
	if !ok {
		t.Error("expected key to exist")
	}
	
	if result != value {
		t.Errorf("expected %v, got %v", value, result)
	}
}

func TestCache_Expiration(t *testing.T) {
	cache := NewCache()
	
	key := "test_key"
	value := "test_value"
	
	cache.Set(key, value, 100*time.Millisecond)
	
	// Should exist immediately
	if _, ok := cache.Get(key); !ok {
		t.Error("key should exist immediately after set")
	}
	
	// Wait for expiration
	time.Sleep(150 * time.Millisecond)
	
	// Should be expired
	if _, ok := cache.Get(key); ok {
		t.Error("key should be expired")
	}
}

func TestCache_Delete(t *testing.T) {
	cache := NewCache()
	
	key := "test_key"
	value := "test_value"
	
	cache.Set(key, value, 1*time.Second)
	
	// Verify it exists
	if _, ok := cache.Get(key); !ok {
		t.Error("key should exist")
	}
	
	// Delete
	cache.Delete(key)
	
	// Should not exist
	if _, ok := cache.Get(key); ok {
		t.Error("key should not exist after delete")
	}
}

func TestCache_Exists(t *testing.T) {
	cache := NewCache()
	
	key := "test_key"
	value := "test_value"
	
	// Should not exist initially
	if cache.Exists(key) {
		t.Error("key should not exist initially")
	}
	
	cache.Set(key, value, 1*time.Second)
	
	// Should exist
	if !cache.Exists(key) {
		t.Error("key should exist after set")
	}
	
	// Set with short expiration
	cache.Set("expire_key", "value", 50*time.Millisecond)
	time.Sleep(100 * time.Millisecond)
	
	// Should not exist after expiration
	if cache.Exists("expire_key") {
		t.Error("key should not exist after expiration")
	}
}

func TestCache_Clear(t *testing.T) {
	cache := NewCache()
	
	cache.Set("key1", "value1", 1*time.Second)
	cache.Set("key2", "value2", 1*time.Second)
	cache.Set("key3", "value3", 1*time.Second)
	
	// Verify they exist
	if !cache.Exists("key1") || !cache.Exists("key2") || !cache.Exists("key3") {
		t.Error("keys should exist")
	}
	
	// Clear all
	cache.Clear()
	
	// Should not exist
	if cache.Exists("key1") || cache.Exists("key2") || cache.Exists("key3") {
		t.Error("no keys should exist after clear")
	}
}

func TestCache_MultipleTypes(t *testing.T) {
	cache := NewCache()
	
	// Test different value types
	cache.Set("string", "test", 1*time.Second)
	cache.Set("int", 42, 1*time.Second)
	cache.Set("bool", true, 1*time.Second)
	cache.Set("struct", struct{ Name string }{"test"}, 1*time.Second)
	
	// Retrieve and verify
	strVal, _ := cache.Get("string")
	if strVal != "test" {
		t.Error("string value mismatch")
	}
	
	intVal, _ := cache.Get("int")
	if intVal != 42 {
		t.Error("int value mismatch")
	}
	
	boolVal, _ := cache.Get("bool")
	if boolVal != true {
		t.Error("bool value mismatch")
	}
}

func TestViewCountCache(t *testing.T) {
	viewCache := NewViewCountCache()
	
	postID := uint(123)
	ip1 := "192.168.1.1"
	ip2 := "192.168.1.2"
	
	// Initially not viewed
	if viewCache.HasViewed(postID, ip1) {
		t.Error("should not have viewed initially")
	}
	
	// Mark as viewed
	viewCache.MarkViewed(postID, ip1)
	
	// Should have viewed
	if !viewCache.HasViewed(postID, ip1) {
		t.Error("should have viewed after marking")
	}
	
	// Different IP should not have viewed
	if viewCache.HasViewed(postID, ip2) {
		t.Error("different IP should not have viewed")
	}
}

func TestGenerateViewKey(t *testing.T) {
	postID := uint(123)
	ip := "192.168.1.1"
	
	key := GenerateViewKey(postID, ip)
	
	expected := "view:123:192.168.1.1"
	if key != expected {
		t.Errorf("expected %s, got %s", expected, key)
	}
}

func TestCache_ConcurrentAccess(t *testing.T) {
	cache := NewCache()
	done := make(chan bool)
	
	// Concurrent writes
	for i := 0; i < 100; i++ {
		go func(n int) {
			cache.Set(string(rune(n)), n, 1*time.Second)
			done <- true
		}(i)
	}
	
	// Wait for all writes
	for i := 0; i < 100; i++ {
		<-done
	}
	
	// Concurrent reads
	for i := 0; i < 100; i++ {
		go func(n int) {
			cache.Get(string(rune(n)))
			done <- true
		}(i)
	}
	
	// Wait for all reads
	for i := 0; i < 100; i++ {
		<-done
	}
}
