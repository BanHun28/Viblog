package utils

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	secret := "test-secret-key-123"
	userID := uint(1)
	email := "test@example.com"
	role := "admin"
	
	tests := []struct {
		name      string
		tokenType TokenType
	}{
		{"access token", AccessToken},
		{"refresh token", RefreshToken},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GenerateToken(userID, email, role, tt.tokenType, secret)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			
			if token == "" {
				t.Error("expected non-empty token")
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	secret := "test-secret-key-123"
	userID := uint(1)
	email := "test@example.com"
	role := "admin"
	
	token, _ := GenerateToken(userID, email, role, AccessToken, secret)
	
	tests := []struct {
		name    string
		token   string
		secret  string
		wantErr bool
	}{
		{"valid token", token, secret, false},
		{"wrong secret", token, "wrong-secret", true},
		{"invalid token", "invalid.token.string", secret, true},
		{"empty token", "", secret, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			claims, err := ParseToken(tt.token, tt.secret)
			
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				
				if claims.UserID != userID {
					t.Errorf("expected userID %d, got %d", userID, claims.UserID)
				}
				
				if claims.Email != email {
					t.Errorf("expected email %s, got %s", email, claims.Email)
				}
				
				if claims.Role != role {
					t.Errorf("expected role %s, got %s", role, claims.Role)
				}
			}
		})
	}
}

func TestValidateTokenType(t *testing.T) {
	secret := "test-secret-key-123"
	
	accessToken, _ := GenerateToken(1, "test@example.com", "user", AccessToken, secret)
	accessClaims, _ := ParseToken(accessToken, secret)
	
	refreshToken, _ := GenerateToken(1, "test@example.com", "user", RefreshToken, secret)
	refreshClaims, _ := ParseToken(refreshToken, secret)
	
	tests := []struct {
		name         string
		claims       *Claims
		expectedType TokenType
		wantErr      bool
	}{
		{"valid access token", accessClaims, AccessToken, false},
		{"valid refresh token", refreshClaims, RefreshToken, false},
		{"invalid type - access as refresh", accessClaims, RefreshToken, true},
		{"invalid type - refresh as access", refreshClaims, AccessToken, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTokenType(tt.claims, tt.expectedType)
			
			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			
			if !tt.wantErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

func TestTokenExpiration(t *testing.T) {
	secret := "test-secret-key-123"
	
	// Test that access token has correct expiration
	accessToken, _ := GenerateToken(1, "test@example.com", "user", AccessToken, secret)
	accessClaims, _ := ParseToken(accessToken, secret)
	
	expectedExpiry := time.Now().Add(AccessTokenDuration)
	actualExpiry := accessClaims.ExpiresAt.Time
	
	// Allow 5 second tolerance for test execution time
	diff := actualExpiry.Sub(expectedExpiry)
	if diff < 0 {
		diff = -diff
	}
	if diff > 5*time.Second {
		t.Errorf("access token expiration mismatch: expected ~%v, got %v (diff: %v)", 
			expectedExpiry, actualExpiry, diff)
	}
	
	// Test that refresh token has correct expiration
	refreshToken, _ := GenerateToken(1, "test@example.com", "user", RefreshToken, secret)
	refreshClaims, _ := ParseToken(refreshToken, secret)
	
	expectedExpiry = time.Now().Add(RefreshTokenDuration)
	actualExpiry = refreshClaims.ExpiresAt.Time
	
	diff = actualExpiry.Sub(expectedExpiry)
	if diff < 0 {
		diff = -diff
	}
	if diff > 5*time.Second {
		t.Errorf("refresh token expiration mismatch: expected ~%v, got %v (diff: %v)", 
			expectedExpiry, actualExpiry, diff)
	}
}
