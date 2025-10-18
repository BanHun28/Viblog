package dto

import "time"

// CreatePostRequest represents the request to create a new post
type CreatePostRequest struct {
	Title           string   `json:"title" binding:"required,min=1,max=255"`
	Slug            string   `json:"slug" binding:"required,min=1,max=255"`
	Content         string   `json:"content" binding:"required"`
	Excerpt         *string  `json:"excerpt,omitempty"`
	FeaturedImage   *string  `json:"featured_image,omitempty"`
	Status          string   `json:"status" binding:"required,oneof=draft published scheduled"`
	PublishedAt     *time.Time `json:"published_at,omitempty"`
	MetaTitle       *string  `json:"meta_title,omitempty"`
	MetaDescription *string  `json:"meta_description,omitempty"`
	MetaKeywords    *string  `json:"meta_keywords,omitempty"`
	CategoryID      *uint    `json:"category_id,omitempty"`
	TagIDs          []uint   `json:"tag_ids,omitempty" binding:"max=10"` // Max 10 tags per post
}

// UpdatePostRequest represents the request to update a post
type UpdatePostRequest struct {
	Title           *string    `json:"title,omitempty" binding:"omitempty,min=1,max=255"`
	Slug            *string    `json:"slug,omitempty" binding:"omitempty,min=1,max=255"`
	Content         *string    `json:"content,omitempty"`
	Excerpt         *string    `json:"excerpt,omitempty"`
	FeaturedImage   *string    `json:"featured_image,omitempty"`
	Status          *string    `json:"status,omitempty" binding:"omitempty,oneof=draft published scheduled"`
	PublishedAt     *time.Time `json:"published_at,omitempty"`
	MetaTitle       *string    `json:"meta_title,omitempty"`
	MetaDescription *string    `json:"meta_description,omitempty"`
	MetaKeywords    *string    `json:"meta_keywords,omitempty"`
	CategoryID      *uint      `json:"category_id,omitempty"`
	TagIDs          []uint     `json:"tag_ids,omitempty" binding:"max=10"`
}

// PostListRequest represents the query parameters for listing posts
type PostListRequest struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

// PostSearchRequest represents the query parameters for searching posts
type PostSearchRequest struct {
	Query string `form:"q" binding:"required,min=1"`
	Page  int    `form:"page" binding:"omitempty,min=1"`
	Limit int    `form:"limit" binding:"omitempty,min=1,max=100"`
}

// PostResponse represents a post in the response
type PostResponse struct {
	ID              uint                  `json:"id"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
	Title           string                `json:"title"`
	Slug            string                `json:"slug"`
	Content         string                `json:"content"`
	Excerpt         string                `json:"excerpt,omitempty"`
	FeaturedImage   *string               `json:"featured_image,omitempty"`
	Status          string                `json:"status"`
	PublishedAt     *time.Time            `json:"published_at,omitempty"`
	ViewCount       int                   `json:"view_count"`
	LikeCount       int                   `json:"like_count"`
	CommentCount    int                   `json:"comment_count"`
	BookmarkCount   int                   `json:"bookmark_count"`
	MetaTitle       *string               `json:"meta_title,omitempty"`
	MetaDescription *string               `json:"meta_description,omitempty"`
	MetaKeywords    *string               `json:"meta_keywords,omitempty"`
	Author          *AuthorResponse       `json:"author,omitempty"`
	Category        *CategoryResponse     `json:"category,omitempty"`
	Tags            []TagResponse         `json:"tags,omitempty"`
	IsLiked         bool                  `json:"is_liked,omitempty"`         // Whether current user liked the post
	IsBookmarked    bool                  `json:"is_bookmarked,omitempty"`    // Whether current user bookmarked the post
}

// PostListResponse represents a paginated list of posts
type PostListResponse struct {
	Posts      []PostResponse     `json:"posts"`
	Pagination PaginationResponse `json:"pagination"`
}

// AuthorResponse represents the author information in post response
type AuthorResponse struct {
	ID        uint   `json:"id"`
	Nickname  string `json:"nickname"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

// PaginationResponse represents pagination metadata
type PaginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}
