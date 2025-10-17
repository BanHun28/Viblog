package user

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
	"github.com/yourusername/viblog/pkg/errors"
	"github.com/yourusername/viblog/pkg/password"
	"github.com/yourusername/viblog/pkg/validator"
)

// LoginUseCase handles user login
type LoginUseCase struct {
	userRepo repository.UserRepository
}

// NewLoginUseCase creates a new LoginUseCase
func NewLoginUseCase(userRepo repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepo: userRepo,
	}
}

// LoginInput represents the input for user login
type LoginInput struct {
	Email    string
	Password string
}

// Execute authenticates a user and returns the user entity
func (uc *LoginUseCase) Execute(ctx context.Context, input LoginInput) (*entity.User, error) {
	// Validate email format
	if !validator.IsValidEmail(input.Email) {
		return nil, errors.ErrInvalidEmail
	}

	// Find user by email
	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.ErrDatabaseError.WithError(err)
	}
	if user == nil {
		return nil, errors.New(
			errors.ErrCodeUnauthorized,
			"Invalid email or password",
			401,
		)
	}

	// Verify password
	if !password.Verify(input.Password, user.Password) {
		return nil, errors.New(
			errors.ErrCodeUnauthorized,
			"Invalid email or password",
			401,
		)
	}

	// Update last login timestamp
	if err := uc.userRepo.UpdateLastLoginAt(ctx, user.ID); err != nil {
		// Log error but don't fail login
		// In production, use proper logging
	}

	return user, nil
}
