package admin

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
)

// ListCommentsUseCase handles listing all comments with pagination
type ListCommentsUseCase struct {
	commentRepo repository.CommentRepository
}

// NewListCommentsUseCase creates a new ListCommentsUseCase
func NewListCommentsUseCase(commentRepo repository.CommentRepository) *ListCommentsUseCase {
	return &ListCommentsUseCase{
		commentRepo: commentRepo,
	}
}

// Execute retrieves all comments with pagination
func (uc *ListCommentsUseCase) Execute(ctx context.Context, page, limit int) ([]entity.Comment, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	return uc.commentRepo.FindAll(ctx, page, limit)
}
