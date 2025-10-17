package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// AccessTokenDuration is the duration for access tokens (15 minutes)
	AccessTokenDuration = 15 * time.Minute
	// RefreshTokenDuration is the duration for refresh tokens (7 days)
	RefreshTokenDuration = 7 * 24 * time.Hour
)

// TokenClaims represents the JWT claims
type TokenClaims struct {
	UserID  uint   `json:"user_id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

// JWTService handles JWT token operations
type JWTService struct {
	secretKey       []byte
	refreshSecretKey []byte
}

// NewJWTService creates a new JWT service
func NewJWTService(secretKey, refreshSecretKey string) *JWTService {
	return &JWTService{
		secretKey:       []byte(secretKey),
		refreshSecretKey: []byte(refreshSecretKey),
	}
}

// GenerateAccessToken generates an access token for a user
func (s *JWTService) GenerateAccessToken(userID uint, email string, isAdmin bool) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID:  userID,
		Email:   email,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(AccessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

// GenerateRefreshToken generates a refresh token for a user
func (s *JWTService) GenerateRefreshToken(userID uint, email string) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(RefreshTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.refreshSecretKey)
}

// ValidateAccessToken validates an access token and returns the claims
func (s *JWTService) ValidateAccessToken(tokenString string) (*TokenClaims, error) {
	return s.validateToken(tokenString, s.secretKey)
}

// ValidateRefreshToken validates a refresh token and returns the claims
func (s *JWTService) ValidateRefreshToken(tokenString string) (*TokenClaims, error) {
	return s.validateToken(tokenString, s.refreshSecretKey)
}

// validateToken validates a token and returns the claims
func (s *JWTService) validateToken(tokenString string, secret []byte) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
