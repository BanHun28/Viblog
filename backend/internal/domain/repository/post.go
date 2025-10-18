package repository

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
)

// PostRepository defines the interface for post data access
type PostRepository interface {
	// Create creates a new post
	Create(ctx context.Context, post *entity.Post) error

	// FindByID finds a post by ID
	FindByID(ctx context.Context, id uint) (*entity.Post, error)

	// FindBySlug finds a post by slug
	FindBySlug(ctx context.Context, slug string) (*entity.Post, error)

	// FindAll retrieves all posts with pagination
	FindAll(ctx context.Context, page, limit int) ([]entity.Post, int64, error)

	// FindPublished retrieves published posts with pagination
	FindPublished(ctx context.Context, page, limit int) ([]entity.Post, int64, error)

	// FindByCategory retrieves posts by category with pagination
	FindByCategory(ctx context.Context, categoryID uint, page, limit int) ([]entity.Post, int64, error)

	// FindByTag retrieves posts by tag with pagination
	FindByTag(ctx context.Context, tagID uint, page, limit int) ([]entity.Post, int64, error)

	// Search searches posts by query
	Search(ctx context.Context, query string) ([]entity.Post, error)

	// Update updates a post
	Update(ctx context.Context, post *entity.Post) error

	// Delete deletes a post (soft delete)
	Delete(ctx context.Context, id uint) error

	// IncrementViewCount increments the view count for a post
	IncrementViewCount(ctx context.Context, id uint) error

	// IncrementLikeCount increments the like count for a post
	IncrementLikeCount(ctx context.Context, id uint) error

	// DecrementLikeCount decrements the like count for a post
	DecrementLikeCount(ctx context.Context, id uint) error

	// IncrementCommentCount increments the comment count for a post
	IncrementCommentCount(ctx context.Context, id uint) error

	// DecrementCommentCount decrements the comment count for a post
	DecrementCommentCount(ctx context.Context, id uint) error

	// GetTotalCount gets total count of all posts
	GetTotalCount(ctx context.Context) (int64, error)

	// GetPublishedCount gets count of published posts
	GetPublishedCount(ctx context.Context) (int64, error)
}
