package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// AccessTokenDuration is 15 minutes
	AccessTokenDuration = 15 * time.Minute
	// RefreshTokenDuration is 7 days
	RefreshTokenDuration = 7 * 24 * time.Hour
)

// TokenType represents the type of JWT token
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// Claims represents JWT claims
type Claims struct {
	UserID   uint      `json:"user_id"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token
func GenerateToken(userID uint, email, role string, tokenType TokenType, secret string) (string, error) {
	var duration time.Duration
	if tokenType == AccessToken {
		duration = AccessTokenDuration
	} else {
		duration = RefreshTokenDuration
	}

	claims := Claims{
		UserID:    userID,
		Email:     email,
		Role:      role,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken parses and validates a JWT token
func ParseToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ValidateTokenType checks if the token is of the expected type
func ValidateTokenType(claims *Claims, expectedType TokenType) error {
	if claims.TokenType != expectedType {
		return fmt.Errorf("invalid token type: expected %s, got %s", expectedType, claims.TokenType)
	}
	return nil
}
