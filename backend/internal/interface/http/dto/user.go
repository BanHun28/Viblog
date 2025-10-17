package dto

import "time"

// RegisterRequest represents a user registration request
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Nickname string `json:"nickname" binding:"required,min=2,max=20"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest represents a profile update request
type UpdateProfileRequest struct {
	Nickname  *string `json:"nickname,omitempty" binding:"omitempty,min=2,max=20"`
	AvatarURL *string `json:"avatar_url,omitempty" binding:"omitempty,url"`
	Bio       *string `json:"bio,omitempty" binding:"omitempty,max=500"`
}

// RefreshTokenRequest represents a token refresh request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UserResponse represents a user in responses
type UserResponse struct {
	ID          uint       `json:"id"`
	Email       string     `json:"email"`
	Nickname    string     `json:"nickname"`
	AvatarURL   *string    `json:"avatar_url,omitempty"`
	Bio         *string    `json:"bio,omitempty"`
	IsAdmin     bool       `json:"is_admin"`
	CreatedAt   time.Time  `json:"created_at"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
}

// AuthResponse represents authentication response with tokens
type AuthResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

// TokenResponse represents a token refresh response
type TokenResponse struct {
	AccessToken string `json:"access_token"`
}
