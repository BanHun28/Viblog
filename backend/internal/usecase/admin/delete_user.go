package admin

import (
	"context"
	"errors"

	"github.com/yourusername/viblog/internal/domain/repository"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrCannotDeleteAdmin  = errors.New("cannot delete admin user")
)

// DeleteUserUseCase handles deleting a user
type DeleteUserUseCase struct {
	userRepo repository.UserRepository
}

// NewDeleteUserUseCase creates a new DeleteUserUseCase
func NewDeleteUserUseCase(userRepo repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepo: userRepo,
	}
}

// Execute deletes a user by ID
func (uc *DeleteUserUseCase) Execute(ctx context.Context, userID uint) error {
	// Find the user first
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return ErrUserNotFound
	}

	// Prevent deletion of admin users
	if user.IsAdmin {
		return ErrCannotDeleteAdmin
	}

	// Delete the user
	return uc.userRepo.Delete(ctx, userID)
}
