package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCORSMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name                 string
		config               CORSConfig
		origin               string
		method               string
		expectedStatus       int
		expectedAllowOrigin  string
		expectedAllowMethods string
	}{
		{
			name: "Allowed origin",
			config: CORSConfig{
				AllowedOrigins:   []string{"http://localhost:30001"},
				AllowedMethods:   []string{"GET", "POST"},
				AllowedHeaders:   []string{"Content-Type"},
				AllowCredentials: true,
			},
			origin:               "http://localhost:30001",
			method:               "GET",
			expectedStatus:       http.StatusOK,
			expectedAllowOrigin:  "http://localhost:30001",
			expectedAllowMethods: "GET, POST",
		},
		{
			name: "Wildcard origin",
			config: CORSConfig{
				AllowedOrigins: []string{"*"},
				AllowedMethods: []string{"GET"},
			},
			origin:               "http://example.com",
			method:               "GET",
			expectedStatus:       http.StatusOK,
			expectedAllowOrigin:  "*",
			expectedAllowMethods: "GET",
		},
		{
			name: "Preflight request",
			config: CORSConfig{
				AllowedOrigins: []string{"http://localhost:30001"},
				AllowedMethods: []string{"GET", "POST"},
			},
			origin:         "http://localhost:30001",
			method:         "OPTIONS",
			expectedStatus: http.StatusNoContent,
		},
		{
			name: "Disallowed origin",
			config: CORSConfig{
				AllowedOrigins: []string{"http://localhost:30001"},
			},
			origin:         "http://evil.com",
			method:         "GET",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			r.Use(CORSMiddleware(tt.config))
			r.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			})
			r.OPTIONS("/test", func(c *gin.Context) {
				c.Status(http.StatusNoContent)
			})

			req := httptest.NewRequest(tt.method, "/test", nil)
			req.Header.Set("Origin", tt.origin)

			c.Request = req
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedAllowOrigin != "" {
				allowOrigin := w.Header().Get("Access-Control-Allow-Origin")
				if allowOrigin != tt.expectedAllowOrigin {
					t.Errorf("Expected Allow-Origin %s, got %s", tt.expectedAllowOrigin, allowOrigin)
				}
			}

			if tt.expectedAllowMethods != "" {
				allowMethods := w.Header().Get("Access-Control-Allow-Methods")
				if allowMethods != tt.expectedAllowMethods {
					t.Errorf("Expected Allow-Methods %s, got %s", tt.expectedAllowMethods, allowMethods)
				}
			}
		})
	}
}

func TestDefaultCORSConfig(t *testing.T) {
	config := DefaultCORSConfig()

	if len(config.AllowedOrigins) == 0 {
		t.Error("Expected default allowed origins")
	}

	if config.AllowedOrigins[0] != "http://localhost:30001" {
		t.Errorf("Expected localhost:30001, got %s", config.AllowedOrigins[0])
	}

	if !config.AllowCredentials {
		t.Error("Expected credentials to be allowed by default")
	}
}

func TestProductionCORSConfig(t *testing.T) {
	origins := []string{"https://example.com", "https://api.example.com"}
	config := ProductionCORSConfig(origins)

	if len(config.AllowedOrigins) != 2 {
		t.Errorf("Expected 2 origins, got %d", len(config.AllowedOrigins))
	}

	if config.AllowedOrigins[0] != origins[0] {
		t.Errorf("Expected %s, got %s", origins[0], config.AllowedOrigins[0])
	}
}
