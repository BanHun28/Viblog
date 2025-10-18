package repository

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
)

// PostRepository defines methods for post persistence
type PostRepository interface {
	// Basic CRUD
	Create(ctx context.Context, post *entity.Post) error
	GetByID(ctx context.Context, id uint) (*entity.Post, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Post, error)
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id uint) error

	// List operations
	List(ctx context.Context, page, limit int) ([]*entity.Post, int64, error)
	ListPublished(ctx context.Context, page, limit int) ([]*entity.Post, int64, error)
	ListByCategory(ctx context.Context, categorySlug string, page, limit int) ([]*entity.Post, int64, error)
	ListByTag(ctx context.Context, tagSlug string, page, limit int) ([]*entity.Post, int64, error)
	ListByAuthor(ctx context.Context, authorID uint, page, limit int) ([]*entity.Post, int64, error)

	// Search
	Search(ctx context.Context, query string, page, limit int) ([]*entity.Post, int64, error)

	// View tracking
	IncrementViewCount(ctx context.Context, postID uint) error
	HasViewedRecently(ctx context.Context, postID uint, ipAddress string) (bool, error)
	RecordView(ctx context.Context, postID uint, ipAddress, userAgent string) error

	// Like operations
	AddLike(ctx context.Context, postID, userID uint) error
	RemoveLike(ctx context.Context, postID, userID uint) error
	HasLiked(ctx context.Context, postID, userID uint) (bool, error)
	IncrementLikeCount(ctx context.Context, postID uint) error
	DecrementLikeCount(ctx context.Context, postID uint) error

	// Bookmark operations
	AddBookmark(ctx context.Context, postID, userID uint) error
	RemoveBookmark(ctx context.Context, postID, userID uint) error
	HasBookmarked(ctx context.Context, postID, userID uint) (bool, error)
	IncrementBookmarkCount(ctx context.Context, postID uint) error
	DecrementBookmarkCount(ctx context.Context, postID uint) error
	ListBookmarkedByUser(ctx context.Context, userID uint, page, limit int) ([]*entity.Post, int64, error)

	// Utilities
	SlugExists(ctx context.Context, slug string, excludeID *uint) (bool, error)
	IncrementCommentCount(ctx context.Context, postID uint) error
	DecrementCommentCount(ctx context.Context, postID uint) error
}
