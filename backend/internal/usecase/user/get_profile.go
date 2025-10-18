package user

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
	"github.com/yourusername/viblog/pkg/errors"
)

// GetProfileInput represents the input for getting user profile
type GetProfileInput struct {
	UserID uint
}

// GetProfileUseCase handles getting user profile logic
type GetProfileUseCase struct {
	userRepo repository.UserRepository
}

// NewGetProfileUseCase creates a new GetProfileUseCase
func NewGetProfileUseCase(userRepo repository.UserRepository) *GetProfileUseCase {
	return &GetProfileUseCase{
		userRepo: userRepo,
	}
}

// Execute retrieves user profile
func (uc *GetProfileUseCase) Execute(ctx context.Context, input GetProfileInput) (*entity.User, error) {
	// Get user from repository
	user, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	if user == nil {
		return nil, errors.ErrNotFound
	}

	return user, nil
}
