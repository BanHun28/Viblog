package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

// CORS handles Cross-Origin Resource Sharing
func CORS(config CORSConfig) gin.HandlerFunc {
	return CORSMiddleware(config)
}

// CORSMiddleware handles Cross-Origin Resource Sharing
func CORSMiddleware(config CORSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if origin is allowed
		allowedOrigin := ""
		for _, allowed := range config.AllowedOrigins {
			if allowed == "*" || allowed == origin {
				allowedOrigin = allowed
				break
			}
		}

		// If no match found and wildcard not present, skip CORS headers
		if allowedOrigin == "" && len(config.AllowedOrigins) > 0 {
			c.Next()
			return
		}

		// Set CORS headers
		if allowedOrigin == "*" {
			c.Header("Access-Control-Allow-Origin", "*")
		} else if allowedOrigin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		if config.AllowCredentials {
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if len(config.AllowedMethods) > 0 {
			methods := ""
			for i, method := range config.AllowedMethods {
				if i > 0 {
					methods += ", "
				}
				methods += method
			}
			c.Header("Access-Control-Allow-Methods", methods)
		}

		if len(config.AllowedHeaders) > 0 {
			headers := ""
			for i, header := range config.AllowedHeaders {
				if i > 0 {
					headers += ", "
				}
				headers += header
			}
			c.Header("Access-Control-Allow-Headers", headers)
		}


		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
