package repository

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
)

// CommentRepository defines the interface for comment data access
type CommentRepository interface {
	// Create creates a new comment
	Create(ctx context.Context, comment *entity.Comment) error

	// FindByID finds a comment by ID
	FindByID(ctx context.Context, id uint) (*entity.Comment, error)

	// FindByPostID retrieves all comments for a specific post
	FindByPostID(ctx context.Context, postID uint) ([]entity.Comment, error)

	// FindReplies retrieves all replies to a specific comment
	FindReplies(ctx context.Context, parentID uint) ([]entity.Comment, error)

	// FindAll retrieves all comments with pagination
	FindAll(ctx context.Context, page, limit int) ([]entity.Comment, int64, error)

	// Update updates a comment
	Update(ctx context.Context, comment *entity.Comment) error

	// Delete deletes a comment (soft delete)
	Delete(ctx context.Context, id uint) error

	// IncrementLikeCount increments the like count for a comment
	IncrementLikeCount(ctx context.Context, id uint) error

	// DecrementLikeCount decrements the like count for a comment
	DecrementLikeCount(ctx context.Context, id uint) error

	// GetTotalCount gets total count of all comments
	GetTotalCount(ctx context.Context) (int64, error)
}
