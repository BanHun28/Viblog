package admin

import (
	"context"

	"github.com/yourusername/viblog/internal/domain/repository"
)

// DashboardStats represents admin dashboard statistics
type DashboardStats struct {
	TotalUsers     int64 `json:"total_users"`
	TotalPosts     int64 `json:"total_posts"`
	PublishedPosts int64 `json:"published_posts"`
	TotalComments  int64 `json:"total_comments"`
}

// GetDashboardUseCase handles retrieving dashboard statistics
type GetDashboardUseCase struct {
	userRepo    repository.UserRepository
	postRepo    repository.PostRepository
	commentRepo repository.CommentRepository
}

// NewGetDashboardUseCase creates a new GetDashboardUseCase
func NewGetDashboardUseCase(
	userRepo repository.UserRepository,
	postRepo repository.PostRepository,
	commentRepo repository.CommentRepository,
) *GetDashboardUseCase {
	return &GetDashboardUseCase{
		userRepo:    userRepo,
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}

// Execute retrieves dashboard statistics
func (uc *GetDashboardUseCase) Execute(ctx context.Context) (*DashboardStats, error) {
	// Get total user count
	totalUsers, err := uc.userRepo.GetTotalCount(ctx)
	if err != nil {
		return nil, err
	}

	// Get total post count
	totalPosts, err := uc.postRepo.GetTotalCount(ctx)
	if err != nil {
		return nil, err
	}

	// Get published post count
	publishedPosts, err := uc.postRepo.GetPublishedCount(ctx)
	if err != nil {
		return nil, err
	}

	// Get total comment count
	totalComments, err := uc.commentRepo.GetTotalCount(ctx)
	if err != nil {
		return nil, err
	}

	return &DashboardStats{
		TotalUsers:     totalUsers,
		TotalPosts:     totalPosts,
		PublishedPosts: publishedPosts,
		TotalComments:  totalComments,
	}, nil
}
