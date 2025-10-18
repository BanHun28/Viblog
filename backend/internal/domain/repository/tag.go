package repository

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
)

// TagRepository defines the interface for tag data access
type TagRepository interface {
	// Create creates a new tag
	Create(ctx context.Context, tag *entity.Tag) error

	// FindByID finds a tag by ID
	FindByID(ctx context.Context, id uint) (*entity.Tag, error)

	// FindBySlug finds a tag by slug
	FindBySlug(ctx context.Context, slug string) (*entity.Tag, error)

	// FindAll retrieves all tags
	FindAll(ctx context.Context) ([]entity.Tag, error)

	// Update updates a tag
	Update(ctx context.Context, tag *entity.Tag) error

	// Delete deletes a tag (soft delete)
	Delete(ctx context.Context, id uint) error

	// ExistsBySlug checks if a tag exists with the given slug
	ExistsBySlug(ctx context.Context, slug string) (bool, error)

	// ExistsByName checks if a tag exists with the given name
	ExistsByName(ctx context.Context, name string) (bool, error)

	// IncrementPostCount increments the post count for a tag
	IncrementPostCount(ctx context.Context, id uint) error

	// DecrementPostCount decrements the post count for a tag
	DecrementPostCount(ctx context.Context, id uint) error
}
