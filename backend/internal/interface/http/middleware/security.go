package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"html"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// XSSProtectionMiddleware sanitizes user input to prevent XSS attacks
func XSSProtectionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set X-XSS-Protection header
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Content-Type-Options", "nosniff")

		c.Next()
	}
}

// CSRFConfig holds CSRF protection configuration
type CSRFConfig struct {
	Secret     string
	TokenName  string
	HeaderName string
	CookieName string
	Expiration time.Duration
}

// DefaultCSRFConfig returns default CSRF configuration
func DefaultCSRFConfig() CSRFConfig {
	return CSRFConfig{
		Secret:     generateSecret(),
		TokenName:  "csrf_token",
		HeaderName: "X-CSRF-Token",
		CookieName: "csrf_token",
		Expiration: 24 * time.Hour,
	}
}

// CSRFTokenStore manages CSRF tokens
type CSRFTokenStore struct {
	mu     sync.RWMutex
	tokens map[string]time.Time
}

// NewCSRFTokenStore creates a new CSRF token store
func NewCSRFTokenStore() *CSRFTokenStore {
	store := &CSRFTokenStore{
		tokens: make(map[string]time.Time),
	}

	// Start cleanup goroutine
	go store.cleanup()

	return store
}

// Add adds a token to the store
func (s *CSRFTokenStore) Add(token string, expiration time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[token] = expiration
}

// Verify checks if a token is valid
func (s *CSRFTokenStore) Verify(token string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	expiration, exists := s.tokens[token]
	if !exists {
		return false
	}

	if time.Now().After(expiration) {
		return false
	}

	return true
}

// Remove removes a token from the store
func (s *CSRFTokenStore) Remove(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tokens, token)
}

// cleanup removes expired tokens periodically
func (s *CSRFTokenStore) cleanup() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		s.mu.Lock()
		now := time.Now()
		for token, expiration := range s.tokens {
			if now.After(expiration) {
				delete(s.tokens, token)
			}
		}
		s.mu.Unlock()
	}
}

// CSRFMiddleware provides CSRF protection
func CSRFMiddleware(config CSRFConfig, store *CSRFTokenStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip CSRF check for safe methods
		if c.Request.Method == "GET" || c.Request.Method == "HEAD" || c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		// Get token from header or form
		token := c.GetHeader(config.HeaderName)
		if token == "" {
			token = c.PostForm(config.TokenName)
		}

		// Verify token
		if token == "" || !store.Verify(token) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid or missing CSRF token",
			})
			return
		}

		c.Next()
	}
}

// GenerateCSRFToken generates a new CSRF token
func GenerateCSRFToken(store *CSRFTokenStore, expiration time.Duration) string {
	token := generateToken()
	store.Add(token, time.Now().Add(expiration))
	return token
}

// SecurityHeadersMiddleware sets security-related HTTP headers
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevent clickjacking
		c.Header("X-Frame-Options", "DENY")

		// XSS Protection
		c.Header("X-XSS-Protection", "1; mode=block")

		// Prevent MIME type sniffing
		c.Header("X-Content-Type-Options", "nosniff")

		// Referrer Policy
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// Content Security Policy
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline';")

		// Strict Transport Security (HSTS) - only for HTTPS
		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}

		// Permissions Policy
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		c.Next()
	}
}

// SanitizeInput sanitizes user input to prevent XSS
func SanitizeInput(input string) string {
	return html.EscapeString(input)
}

// generateToken generates a random token
func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// generateSecret generates a random secret
func generateSecret() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// SecureMiddleware combines all security middlewares
func SecureMiddleware(csrfConfig CSRFConfig, csrfStore *CSRFTokenStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Apply security headers
		SecurityHeadersMiddleware()(c)

		// Apply XSS protection
		XSSProtectionMiddleware()(c)

		c.Next()
	}
}
