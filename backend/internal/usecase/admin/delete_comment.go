package admin

import (
	"context"
	"errors"

	"github.com/yourusername/viblog/internal/domain/repository"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
)

// DeleteCommentUseCase handles deleting a comment
type DeleteCommentUseCase struct {
	commentRepo repository.CommentRepository
	postRepo    repository.PostRepository
}

// NewDeleteCommentUseCase creates a new DeleteCommentUseCase
func NewDeleteCommentUseCase(
	commentRepo repository.CommentRepository,
	postRepo repository.PostRepository,
) *DeleteCommentUseCase {
	return &DeleteCommentUseCase{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

// Execute deletes a comment by ID
func (uc *DeleteCommentUseCase) Execute(ctx context.Context, commentID uint) error {
	// Find the comment first
	comment, err := uc.commentRepo.FindByID(ctx, commentID)
	if err != nil {
		return ErrCommentNotFound
	}

	// Delete the comment
	if err := uc.commentRepo.Delete(ctx, commentID); err != nil {
		return err
	}

	// Decrement the post's comment count
	if err := uc.postRepo.DecrementCommentCount(ctx, comment.PostID); err != nil {
		return err
	}

	return nil
}
