package presenter

import (
	"math"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/interface/http/dto"
	"github.com/yourusername/viblog/internal/usecase/admin"
)

// PresentDashboardStats converts dashboard stats to response DTO
func PresentDashboardStats(stats *admin.DashboardStats) dto.AdminDashboardResponse {
	return dto.AdminDashboardResponse{
		TotalUsers:     stats.TotalUsers,
		TotalPosts:     stats.TotalPosts,
		PublishedPosts: stats.PublishedPosts,
		TotalComments:  stats.TotalComments,
	}
}

// PresentUser converts a user entity to admin user response
func PresentUser(user *entity.User) dto.AdminUserResponse {
	var lastLoginAt *string
	if user.LastLoginAt != nil {
		formatted := user.LastLoginAt.Format("2006-01-02T15:04:05Z")
		lastLoginAt = &formatted
	}

	return dto.AdminUserResponse{
		ID:          user.ID,
		Email:       user.Email,
		Nickname:    user.Nickname,
		AvatarURL:   user.AvatarURL,
		Bio:         user.Bio,
		IsAdmin:     user.IsAdmin,
		CreatedAt:   user.CreatedAt.Format("2006-01-02T15:04:05Z"),
		LastLoginAt: lastLoginAt,
	}
}

// PresentUsersList converts a list of users to paginated response
func PresentUsersList(users []entity.User, total int64, page, limit int) dto.AdminUsersListResponse {
	userResponses := make([]dto.AdminUserResponse, len(users))
	for i, user := range users {
		userResponses[i] = PresentUser(&user)
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return dto.AdminUsersListResponse{
		Users:      userResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

// PresentComment converts a comment entity to admin comment response
func PresentComment(comment *entity.Comment) dto.AdminCommentResponse {
	var postTitle *string
	if comment.Post != nil {
		postTitle = &comment.Post.Title
	}

	var userName *string
	if comment.User != nil {
		userName = &comment.User.Nickname
	}

	return dto.AdminCommentResponse{
		ID:          comment.ID,
		Content:     comment.Content,
		PostID:      comment.PostID,
		PostTitle:   postTitle,
		UserID:      comment.UserID,
		UserName:    userName,
		AuthorName:  comment.AuthorName,
		AuthorEmail: comment.AuthorEmail,
		ParentID:    comment.ParentID,
		LikeCount:   comment.LikeCount,
		IsEdited:    comment.IsEdited,
		CreatedAt:   comment.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   comment.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

// PresentCommentsList converts a list of comments to paginated response
func PresentCommentsList(comments []entity.Comment, total int64, page, limit int) dto.AdminCommentsListResponse {
	commentResponses := make([]dto.AdminCommentResponse, len(comments))
	for i, comment := range comments {
		commentResponses[i] = PresentComment(&comment)
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return dto.AdminCommentsListResponse{
		Comments:   commentResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

// PresentCategory converts a category entity to response DTO
func PresentCategory(category *entity.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description,
		PostCount:   category.PostCount,
		CreatedAt:   category.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   category.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

// PresentCategoriesList converts a list of categories to response DTO
func PresentCategoriesList(categories []entity.Category) dto.CategoriesListResponse {
	categoryResponses := make([]dto.CategoryResponse, len(categories))
	for i, category := range categories {
		categoryResponses[i] = PresentCategory(&category)
	}

	return dto.CategoriesListResponse{
		Categories: categoryResponses,
		Total:      len(categories),
	}
}

// PresentTag converts a tag entity to response DTO
func PresentTag(tag *entity.Tag) dto.TagResponse {
	return dto.TagResponse{
		ID:        tag.ID,
		Name:      tag.Name,
		Slug:      tag.Slug,
		PostCount: tag.PostCount,
		CreatedAt: tag.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: tag.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

// PresentTagsList converts a list of tags to response DTO
func PresentTagsList(tags []entity.Tag) dto.TagsListResponse {
	tagResponses := make([]dto.TagResponse, len(tags))
	for i, tag := range tags {
		tagResponses[i] = PresentTag(&tag)
	}

	return dto.TagsListResponse{
		Tags:  tagResponses,
		Total: len(tags),
	}
}
