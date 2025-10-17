package user

import (
	"context"
	"fmt"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
	"github.com/yourusername/viblog/pkg/errors"
	"github.com/yourusername/viblog/pkg/password"
	"github.com/yourusername/viblog/pkg/validator"
)

// RegisterUseCase handles user registration
type RegisterUseCase struct {
	userRepo repository.UserRepository
}

// NewRegisterUseCase creates a new RegisterUseCase
func NewRegisterUseCase(userRepo repository.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo: userRepo,
	}
}

// RegisterInput represents the input for user registration
type RegisterInput struct {
	Email    string
	Password string
	Nickname string
}

// Execute registers a new user
func (uc *RegisterUseCase) Execute(ctx context.Context, input RegisterInput) (*entity.User, error) {
	// Validate email format
	if !validator.IsValidEmail(input.Email) {
		return nil, errors.ErrInvalidEmail
	}

	// Validate password strength
	if !validator.IsValidPassword(input.Password) {
		return nil, errors.ErrPasswordTooWeak
	}

	// Validate nickname format
	if !validator.IsValidNickname(input.Nickname) {
		return nil, errors.New(
			errors.ErrCodeValidation,
			"Nickname must be 2-20 characters and contain only alphanumeric, underscore, or hyphen",
			400,
		)
	}

	// Check if email already exists
	existingUser, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.ErrDatabaseError.WithError(err)
	}
	if existingUser != nil {
		return nil, errors.New(
			errors.ErrCodeAlreadyExists,
			"Email already exists",
			409,
		)
	}

	// Check if nickname already exists
	existingNickname, err := uc.userRepo.FindByNickname(ctx, input.Nickname)
	if err != nil {
		return nil, errors.ErrDatabaseError.WithError(err)
	}
	if existingNickname != nil {
		return nil, errors.New(
			errors.ErrCodeAlreadyExists,
			"Nickname already exists",
			409,
		)
	}

	// Hash password
	hashedPassword, err := password.Hash(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user entity
	user := &entity.User{
		Email:    input.Email,
		Password: hashedPassword,
		Nickname: input.Nickname,
		IsAdmin:  false, // Regular users are not admins by default
	}

	// Save user to database
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, errors.ErrDatabaseError.WithError(err)
	}

	return user, nil
}
