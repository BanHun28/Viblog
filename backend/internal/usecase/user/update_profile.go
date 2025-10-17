package user

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
	"github.com/yourusername/viblog/pkg/errors"
	"github.com/yourusername/viblog/pkg/validator"
)

// UpdateProfileUseCase handles user profile updates
type UpdateProfileUseCase struct {
	userRepo repository.UserRepository
}

// NewUpdateProfileUseCase creates a new UpdateProfileUseCase
func NewUpdateProfileUseCase(userRepo repository.UserRepository) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{
		userRepo: userRepo,
	}
}

// UpdateProfileInput represents the input for updating user profile
type UpdateProfileInput struct {
	UserID    uint
	Nickname  *string
	AvatarURL *string
	Bio       *string
}

// Execute updates a user's profile
func (uc *UpdateProfileUseCase) Execute(ctx context.Context, input UpdateProfileInput) (*entity.User, error) {
	// Get existing user
	user, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, errors.ErrDatabaseError.WithError(err)
	}
	if user == nil {
		return nil, errors.ErrNotFound
	}

	// Update nickname if provided
	if input.Nickname != nil {
		// Validate nickname format
		if !validator.IsValidNickname(*input.Nickname) {
			return nil, errors.New(
				errors.ErrCodeValidation,
				"Nickname must be 2-20 characters and contain only alphanumeric, underscore, or hyphen",
				400,
			)
		}

		// Check if nickname is different from current
		if *input.Nickname != user.Nickname {
			// Check if nickname already exists
			existingUser, err := uc.userRepo.FindByNickname(ctx, *input.Nickname)
			if err != nil {
				return nil, errors.ErrDatabaseError.WithError(err)
			}
			if existingUser != nil && existingUser.ID != user.ID {
				return nil, errors.New(
					errors.ErrCodeAlreadyExists,
					"Nickname already exists",
					409,
				)
			}

			user.Nickname = *input.Nickname
		}
	}

	// Update avatar URL if provided
	if input.AvatarURL != nil {
		// Validate URL format
		if *input.AvatarURL != "" && !validator.IsValidURL(*input.AvatarURL) {
			return nil, errors.ErrInvalidURL
		}
		user.AvatarURL = input.AvatarURL
	}

	// Update bio if provided
	if input.Bio != nil {
		user.Bio = input.Bio
	}

	// Save updated user
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, errors.ErrDatabaseError.WithError(err)
	}

	return user, nil
}
