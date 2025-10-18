package repository

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *entity.User) error

	// FindByID finds a user by ID
	FindByID(ctx context.Context, id uint) (*entity.User, error)

	// FindByEmail finds a user by email
	FindByEmail(ctx context.Context, email string) (*entity.User, error)

	// FindByNickname finds a user by nickname
	FindByNickname(ctx context.Context, nickname string) (*entity.User, error)

	// Update updates a user
	Update(ctx context.Context, user *entity.User) error

	// Delete deletes a user (soft delete)
	Delete(ctx context.Context, id uint) error

	// ExistsByEmail checks if a user exists with the given email
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	// ExistsByNickname checks if a user exists with the given nickname
	ExistsByNickname(ctx context.Context, nickname string) (bool, error)

	// UpdateLastLoginAt updates the last login timestamp
	UpdateLastLoginAt(ctx context.Context, id uint) error

	// FindAll retrieves all users with pagination
	FindAll(ctx context.Context, page, limit int) ([]entity.User, int64, error)

	// GetTotalCount gets total count of all users
	GetTotalCount(ctx context.Context) (int64, error)
}
