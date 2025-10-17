package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := JWTConfig{
		SecretKey: "test-secret-key",
	}

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		expectedError  bool
	}{
		{
			name:           "Valid token",
			authHeader:     generateValidToken(config.SecretKey),
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "Missing authorization header",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectedError:  true,
		},
		{
			name:           "Invalid format - missing Bearer prefix",
			authHeader:     "InvalidToken",
			expectedStatus: http.StatusUnauthorized,
			expectedError:  true,
		},
		{
			name:           "Invalid token",
			authHeader:     "Bearer invalid.token.here",
			expectedStatus: http.StatusUnauthorized,
			expectedError:  true,
		},
		{
			name:           "Expired token",
			authHeader:     generateExpiredToken(config.SecretKey),
			expectedStatus: http.StatusUnauthorized,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			r.Use(AuthMiddleware(config))
			r.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			})

			req := httptest.NewRequest("GET", "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			c.Request = req
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestAdminMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		userRole       string
		setRole        bool
		expectedStatus int
	}{
		{
			name:           "Admin role",
			userRole:       "admin",
			setRole:        true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "User role",
			userRole:       "user",
			setRole:        true,
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "No role set",
			setRole:        false,
			expectedStatus: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			r.Use(func(c *gin.Context) {
				if tt.setRole {
					c.Set(UserRoleKey, tt.userRole)
				}
				c.Next()
			})
			r.Use(AdminMiddleware())
			r.GET("/admin", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "admin access granted"})
			})

			req := httptest.NewRequest("GET", "/admin", nil)
			c.Request = req
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestOptionalAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := JWTConfig{
		SecretKey: "test-secret-key",
	}

	tests := []struct {
		name       string
		authHeader string
		shouldAuth bool
	}{
		{
			name:       "Valid token",
			authHeader: generateValidToken(config.SecretKey),
			shouldAuth: true,
		},
		{
			name:       "No token",
			authHeader: "",
			shouldAuth: false,
		},
		{
			name:       "Invalid token",
			authHeader: "Bearer invalid.token",
			shouldAuth: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			r.Use(OptionalAuthMiddleware(config))
			r.GET("/test", func(c *gin.Context) {
				_, exists := c.Get(UserIDKey)
				c.JSON(http.StatusOK, gin.H{"authenticated": exists})
			})

			req := httptest.NewRequest("GET", "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			c.Request = req
			r.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", w.Code)
			}
		})
	}
}

// Helper functions
func generateValidToken(secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(1),
		"role":    "admin",
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(secret))
	return "Bearer " + tokenString
}

func generateExpiredToken(secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(1),
		"role":    "admin",
		"exp":     time.Now().Add(-time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(secret))
	return "Bearer " + tokenString
}
