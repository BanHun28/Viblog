package dto

// AdminDashboardResponse represents the admin dashboard response
type AdminDashboardResponse struct {
	TotalUsers     int64 `json:"total_users"`
	TotalPosts     int64 `json:"total_posts"`
	PublishedPosts int64 `json:"published_posts"`
	TotalComments  int64 `json:"total_comments"`
}

// AdminUserResponse represents a user in admin context
type AdminUserResponse struct {
	ID          uint    `json:"id"`
	Email       string  `json:"email"`
	Nickname    string  `json:"nickname"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	Bio         *string `json:"bio,omitempty"`
	IsAdmin     bool    `json:"is_admin"`
	CreatedAt   string  `json:"created_at"`
	LastLoginAt *string `json:"last_login_at,omitempty"`
}

// AdminUsersListResponse represents paginated users list response
type AdminUsersListResponse struct {
	Users      []AdminUserResponse `json:"users"`
	Total      int64               `json:"total"`
	Page       int                 `json:"page"`
	Limit      int                 `json:"limit"`
	TotalPages int                 `json:"total_pages"`
}

// AdminCommentResponse represents a comment in admin context
type AdminCommentResponse struct {
	ID          uint    `json:"id"`
	Content     string  `json:"content"`
	PostID      uint    `json:"post_id"`
	PostTitle   *string `json:"post_title,omitempty"`
	UserID      *uint   `json:"user_id,omitempty"`
	UserName    *string `json:"user_name,omitempty"`
	AuthorName  *string `json:"author_name,omitempty"`
	AuthorEmail *string `json:"author_email,omitempty"`
	ParentID    *uint   `json:"parent_id,omitempty"`
	LikeCount   int     `json:"like_count"`
	IsEdited    bool    `json:"is_edited"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// AdminCommentsListResponse represents paginated comments list response
type AdminCommentsListResponse struct {
	Comments   []AdminCommentResponse `json:"comments"`
	Total      int64                  `json:"total"`
	Page       int                    `json:"page"`
	Limit      int                    `json:"limit"`
	TotalPages int                    `json:"total_pages"`
}

// CreateCategoryRequest represents a category creation request
type CreateCategoryRequest struct {
	Name        string  `json:"name" binding:"required,min=1,max=100"`
	Description *string `json:"description,omitempty"`
}

// UpdateCategoryRequest represents a category update request
type UpdateCategoryRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=1,max=100"`
	Description *string `json:"description,omitempty"`
}

// CategoryResponse represents a category response
type CategoryResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description,omitempty"`
	PostCount   int     `json:"post_count"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// CategoriesListResponse represents a list of categories
type CategoriesListResponse struct {
	Categories []CategoryResponse `json:"categories"`
	Total      int                `json:"total"`
}

// CreateTagRequest represents a tag creation request
type CreateTagRequest struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
}

// UpdateTagRequest represents a tag update request
type UpdateTagRequest struct {
	Name *string `json:"name,omitempty" binding:"omitempty,min=1,max=50"`
}

// TagResponse represents a tag response
type TagResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	PostCount int    `json:"post_count"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// TagsListResponse represents a list of tags
type TagsListResponse struct {
	Tags  []TagResponse `json:"tags"`
	Total int           `json:"total"`
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message"`
}
