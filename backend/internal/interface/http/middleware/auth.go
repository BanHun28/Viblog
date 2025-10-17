package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	UserIDKey           = "userID"
	UserRoleKey         = "userRole"
)

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey string
}

// AuthMiddleware validates JWT tokens from Authorization header
func AuthMiddleware(config JWTConfig) gin.HandlerFunc {
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

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.SecretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Set user context
			if userID, ok := claims["user_id"].(float64); ok {
				c.Set(UserIDKey, uint(userID))
			}
			if role, ok := claims["role"].(string); ok {
				c.Set(UserRoleKey, role)
			}
		}

		c.Next()
	}
}

// AdminMiddleware requires admin role
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(UserRoleKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "user role not found",
			})
			return
		}

		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "admin access required",
			})
			return
		}

		c.Next()
	}
}

// OptionalAuthMiddleware validates token if present but doesn't require it
func OptionalAuthMiddleware(config JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			c.Next()
			return
		}

		// Try to validate token
		if strings.HasPrefix(authHeader, BearerPrefix) {
			tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(config.SecretKey), nil
			})

			if err == nil && token.Valid {
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					if userID, ok := claims["user_id"].(float64); ok {
						c.Set(UserIDKey, uint(userID))
					}
					if role, ok := claims["role"].(string); ok {
						c.Set(UserRoleKey, role)
					}
				}
			}
		}

		c.Next()
	}
}
