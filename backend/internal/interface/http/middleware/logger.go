package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger logs HTTP requests using Zap logger
func Logger(logger *zap.Logger) gin.HandlerFunc {
	return LoggerMiddleware(logger)
}

// Recovery recovers from panics and logs them
func Recovery(logger *zap.Logger) gin.HandlerFunc {
	return RecoveryMiddleware(logger)
}

// LoggerMiddleware logs HTTP requests using Zap logger
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Prepare log fields
		fields := []zap.Field{
			zap.Int("status", statusCode),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", clientIP),
			zap.String("user_agent", userAgent),
			zap.Duration("latency", latency),
		}

		if errorMessage != "" {
			fields = append(fields, zap.String("error", errorMessage))
		}

		// Log based on status code
		switch {
		case statusCode >= 500:
			logger.Error("Server error", fields...)
		case statusCode >= 400:
			logger.Warn("Client error", fields...)
		case statusCode >= 300:
			logger.Info("Redirection", fields...)
		default:
			logger.Info("Request completed", fields...)
		}
	}
}

// StructuredLoggerMiddleware logs with additional context
func StructuredLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Get user context if available
		userID, _ := c.Get("userID")
		userRole, _ := c.Get("userRole")

		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", time.Since(start)),
			zap.Int("body_size", c.Writer.Size()),
		}

		// Add user context if authenticated
		if userID != nil {
			fields = append(fields, zap.Any("user_id", userID))
		}
		if userRole != nil {
			fields = append(fields, zap.Any("user_role", userRole))
		}

		// Log errors if any
		if len(c.Errors) > 0 {
			fields = append(fields, zap.String("errors", c.Errors.String()))
		}

		// Log based on status
		statusCode := c.Writer.Status()
		switch {
		case statusCode >= 500:
			logger.Error("Server error", fields...)
		case statusCode >= 400:
			logger.Warn("Client error", fields...)
		default:
			logger.Info("Request processed", fields...)
		}
	}
}

// RecoveryMiddleware recovers from panics and logs them
func RecoveryMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
					zap.String("ip", c.ClientIP()),
				)

				c.AbortWithStatusJSON(500, gin.H{
					"error": "Internal server error",
				})
			}
		}()

		c.Next()
	}
}
