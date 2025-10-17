package utils

import (
	"fmt"
	"sync"
	"time"
)

// CacheItem represents a cached item with expiration
type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

// Cache is a simple in-memory cache with expiration
type Cache struct {
	mu    sync.RWMutex
	items map[string]*CacheItem
}

// NewCache creates a new cache instance
func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]*CacheItem),
	}
	
	// Start cleanup goroutine
	go c.cleanupExpired()
	
	return c
}

// Set adds an item to the cache with expiration
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	expiration := time.Now().Add(duration)
	c.items[key] = &CacheItem{
		Value:      value,
		Expiration: expiration,
	}
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	item, exists := c.items[key]
	if !exists {
		return nil, false
	}
	
	// Check if expired
	if time.Now().After(item.Expiration) {
		return nil, false
	}
	
	return item.Value, true
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// Exists checks if a key exists and is not expired
func (c *Cache) Exists(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	item, exists := c.items[key]
	if !exists {
		return false
	}
	
	return time.Now().Before(item.Expiration)
}

// Clear removes all items from the cache
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]*CacheItem)
}

// cleanupExpired removes expired items periodically
func (c *Cache) cleanupExpired() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, item := range c.items {
			if now.After(item.Expiration) {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}

// ViewCountCache is a specialized cache for view counts (24-hour IP-based)
type ViewCountCache struct {
	cache *Cache
}

// NewViewCountCache creates a new view count cache
func NewViewCountCache() *ViewCountCache {
	return &ViewCountCache{
		cache: NewCache(),
	}
}

// HasViewed checks if an IP has viewed a post within 24 hours
func (v *ViewCountCache) HasViewed(postID uint, ip string) bool {
	key := GenerateViewKey(postID, ip)
	return v.cache.Exists(key)
}

// MarkViewed marks a post as viewed by an IP for 24 hours
func (v *ViewCountCache) MarkViewed(postID uint, ip string) {
	key := GenerateViewKey(postID, ip)
	v.cache.Set(key, true, 24*time.Hour)
}

// GenerateViewKey generates a cache key for view count
func GenerateViewKey(postID uint, ip string) string {
	return fmt.Sprintf("view:%d:%s", postID, ip)
}
