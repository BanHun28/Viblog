package repository

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
)

// CategoryRepository defines the interface for category data access
type CategoryRepository interface {
	// Create creates a new category
	Create(ctx context.Context, category *entity.Category) error

	// FindByID finds a category by ID
	FindByID(ctx context.Context, id uint) (*entity.Category, error)

	// FindBySlug finds a category by slug
	FindBySlug(ctx context.Context, slug string) (*entity.Category, error)

	// FindAll retrieves all categories
	FindAll(ctx context.Context) ([]entity.Category, error)

	// Update updates a category
	Update(ctx context.Context, category *entity.Category) error

	// Delete deletes a category (soft delete)
	Delete(ctx context.Context, id uint) error

	// ExistsBySlug checks if a category exists with the given slug
	ExistsBySlug(ctx context.Context, slug string) (bool, error)

	// ExistsByName checks if a category exists with the given name
	ExistsByName(ctx context.Context, name string) (bool, error)

	// IncrementPostCount increments the post count for a category
	IncrementPostCount(ctx context.Context, id uint) error

	// DecrementPostCount decrements the post count for a category
	DecrementPostCount(ctx context.Context, id uint) error
}
