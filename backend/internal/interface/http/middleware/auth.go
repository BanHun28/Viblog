package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/viblog/internal/infrastructure/auth"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	UserIDKey           = "userID"
	UserEmailKey        = "userEmail"
	IsAdminKey          = "isAdmin"
)

// AuthMiddleware validates JWT tokens from Authorization header
func AuthMiddleware(jwtService *auth.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		// Extract token from "Bearer <token>"
		if !strings.HasPrefix(authHeader, BearerPrefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header format",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)

		// Validate token using JWT service
		claims, err := jwtService.ValidateAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		// Set user context
		c.Set(UserIDKey, claims.UserID)
		c.Set(UserEmailKey, claims.Email)
		c.Set(IsAdminKey, claims.IsAdmin)

		c.Next()
	}
}

// RequireAdmin requires admin role
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get(IsAdminKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "user authentication required",
			})
			return
		}

		if !isAdmin.(bool) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "admin access required",
			})
			return
		}

		c.Next()
	}
}

// OptionalAuth validates token if present but doesn't require it
func OptionalAuth(jwtService *auth.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			c.Next()
			return
		}

		// Try to validate token
		if strings.HasPrefix(authHeader, BearerPrefix) {
			tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
			claims, err := jwtService.ValidateAccessToken(tokenString)

			if err == nil && claims != nil {
				c.Set(UserIDKey, claims.UserID)
				c.Set(UserEmailKey, claims.Email)
				c.Set(IsAdminKey, claims.IsAdmin)
			}
		}

		c.Next()
	}
}
