package post

import (
	"context"
	"errors"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
)

// GetUseCase handles getting a single post
type GetUseCase struct {
	postRepo repository.PostRepository
}

// NewGetUseCase creates a new GetUseCase
func NewGetUseCase(postRepo repository.PostRepository) *GetUseCase {
	return &GetUseCase{
		postRepo: postRepo,
	}
}

// Execute gets a post by ID
func (uc *GetUseCase) Execute(ctx context.Context, id uint) (*entity.Post, error) {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("post not found")
	}
	return post, nil
}

// GetLikedAndBookmarkedStatus gets like and bookmark status for a post
func (uc *GetUseCase) GetLikedAndBookmarkedStatus(ctx context.Context, postID, userID uint) (bool, bool, error) {
	liked, err := uc.postRepo.HasLiked(ctx, postID, userID)
	if err != nil {
		return false, false, err
	}

	bookmarked, err := uc.postRepo.HasBookmarked(ctx, postID, userID)
	if err != nil {
		return false, false, err
	}

	return liked, bookmarked, nil
}
