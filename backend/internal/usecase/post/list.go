package post

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
)

// ListUseCase handles listing posts
type ListUseCase struct {
	postRepo repository.PostRepository
}

// NewListUseCase creates a new ListUseCase
func NewListUseCase(postRepo repository.PostRepository) *ListUseCase {
	return &ListUseCase{
		postRepo: postRepo,
	}
}

// Execute lists published posts with pagination
func (uc *ListUseCase) Execute(ctx context.Context, page, limit int) ([]*entity.Post, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	return uc.postRepo.ListPublished(ctx, page, limit)
}

// GetLikedAndBookmarkedStatus gets like and bookmark status for posts
func (uc *ListUseCase) GetLikedAndBookmarkedStatus(ctx context.Context, posts []*entity.Post, userID uint) (map[uint]bool, map[uint]bool, error) {
	likedPosts := make(map[uint]bool)
	bookmarkedPosts := make(map[uint]bool)

	for _, post := range posts {
		liked, err := uc.postRepo.HasLiked(ctx, post.ID, userID)
		if err != nil {
			return nil, nil, err
		}
		likedPosts[post.ID] = liked

		bookmarked, err := uc.postRepo.HasBookmarked(ctx, post.ID, userID)
		if err != nil {
			return nil, nil, err
		}
		bookmarkedPosts[post.ID] = bookmarked
	}

	return likedPosts, bookmarkedPosts, nil
}
