package presenter

import (
	"math"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/interface/http/dto"
)

// ToPostResponse converts a post entity to a post response DTO
func ToPostResponse(post *entity.Post, isLiked, isBookmarked bool) dto.PostResponse {
	response := dto.PostResponse{
		ID:              post.ID,
		CreatedAt:       post.CreatedAt,
		UpdatedAt:       post.UpdatedAt,
		Title:           post.Title,
		Slug:            post.Slug,
		Content:         post.Content,
		FeaturedImage:   post.FeaturedImage,
		Status:          post.Status,
		PublishedAt:     post.PublishedAt,
		ViewCount:       post.ViewCount,
		LikeCount:       post.LikeCount,
		CommentCount:    post.CommentCount,
		BookmarkCount:   post.BookmarkCount,
		MetaTitle:       post.MetaTitle,
		MetaDescription: post.MetaDescription,
		MetaKeywords:    post.MetaKeywords,
		IsLiked:         isLiked,
		IsBookmarked:    isBookmarked,
	}

	// Set excerpt
	if post.Excerpt != "" {
		response.Excerpt = post.Excerpt
	}

	// Add author information
	if post.Author != nil {
		response.Author = &dto.AuthorResponse{
			ID:        post.Author.ID,
			Nickname:  post.Author.Nickname,
			AvatarURL: post.Author.AvatarURL,
		}
	}

	// Add category information
	if post.Category != nil {
		response.Category = &dto.CategoryResponse{
			ID:   post.Category.ID,
			Name: post.Category.Name,
			Slug: post.Category.Slug,
		}
	}

	// Add tags information
	if len(post.Tags) > 0 {
		response.Tags = make([]dto.TagResponse, len(post.Tags))
		for i, tag := range post.Tags {
			response.Tags[i] = dto.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
				Slug: tag.Slug,
			}
		}
	}

	return response
}

// ToPostListResponse converts a list of post entities to a post list response
func ToPostListResponse(posts []*entity.Post, total int64, page, limit int, userID *uint, likedPosts, bookmarkedPosts map[uint]bool) dto.PostListResponse {
	postResponses := make([]dto.PostResponse, len(posts))

	for i, post := range posts {
		isLiked := false
		isBookmarked := false

		if userID != nil {
			isLiked = likedPosts[post.ID]
			isBookmarked = bookmarkedPosts[post.ID]
		}

		postResponses[i] = ToPostResponse(post, isLiked, isBookmarked)
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return dto.PostListResponse{
		Posts: postResponses,
		Pagination: dto.PaginationResponse{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}
}
