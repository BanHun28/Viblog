package repository

import (
	"context"
	"errors"
	"time"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
	"gorm.io/gorm"
)

// postRepository implements the PostRepository interface
type postRepository struct {
	db *gorm.DB
}

// NewPostRepository creates a new post repository
func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db: db}
}

// Create creates a new post
func (r *postRepository) Create(ctx context.Context, post *entity.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

// GetByID retrieves a post by ID with all associations
func (r *postRepository) GetByID(ctx context.Context, id uint) (*entity.Post, error) {
	var post entity.Post
	err := r.db.WithContext(ctx).
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		First(&post, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

// GetBySlug retrieves a post by slug with all associations
func (r *postRepository) GetBySlug(ctx context.Context, slug string) (*entity.Post, error) {
	var post entity.Post
	err := r.db.WithContext(ctx).
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Where("slug = ?", slug).
		First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

// Update updates a post
func (r *postRepository) Update(ctx context.Context, post *entity.Post) error {
	return r.db.WithContext(ctx).Save(post).Error
}

// Delete deletes a post (soft delete)
func (r *postRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Post{}, id).Error
}

// List retrieves all posts with pagination
func (r *postRepository) List(ctx context.Context, page, limit int) ([]*entity.Post, int64, error) {
	var posts []*entity.Post
	var total int64

	offset := (page - 1) * limit

	// Count total
	if err := r.db.WithContext(ctx).Model(&entity.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get posts with pagination
	err := r.db.WithContext(ctx).
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return posts, total, err
}

// ListPublished retrieves published posts with pagination
func (r *postRepository) ListPublished(ctx context.Context, page, limit int) ([]*entity.Post, int64, error) {
	var posts []*entity.Post
	var total int64

	offset := (page - 1) * limit

	query := r.db.WithContext(ctx).
		Where("status = ?", "published").
		Where("published_at <= ?", time.Now())

	// Count total
	if err := query.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get posts with pagination
	err := query.
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Order("published_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return posts, total, err
}

// ListByCategory retrieves posts by category slug
func (r *postRepository) ListByCategory(ctx context.Context, categorySlug string, page, limit int) ([]*entity.Post, int64, error) {
	var posts []*entity.Post
	var total int64

	offset := (page - 1) * limit

	query := r.db.WithContext(ctx).
		Joins("JOIN categories ON categories.id = posts.category_id").
		Where("categories.slug = ?", categorySlug).
		Where("posts.status = ?", "published").
		Where("posts.published_at <= ?", time.Now())

	// Count total
	if err := query.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get posts
	err := query.
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Order("posts.published_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return posts, total, err
}

// ListByTag retrieves posts by tag slug
func (r *postRepository) ListByTag(ctx context.Context, tagSlug string, page, limit int) ([]*entity.Post, int64, error) {
	var posts []*entity.Post
	var total int64

	offset := (page - 1) * limit

	query := r.db.WithContext(ctx).
		Joins("JOIN post_tags ON post_tags.post_id = posts.id").
		Joins("JOIN tags ON tags.id = post_tags.tag_id").
		Where("tags.slug = ?", tagSlug).
		Where("posts.status = ?", "published").
		Where("posts.published_at <= ?", time.Now())

	// Count total
	if err := query.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get posts
	err := query.
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Order("posts.published_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return posts, total, err
}

// ListByAuthor retrieves posts by author ID
func (r *postRepository) ListByAuthor(ctx context.Context, authorID uint, page, limit int) ([]*entity.Post, int64, error) {
	var posts []*entity.Post
	var total int64

	offset := (page - 1) * limit

	query := r.db.WithContext(ctx).
		Where("author_id = ?", authorID).
		Where("status = ?", "published").
		Where("published_at <= ?", time.Now())

	// Count total
	if err := query.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get posts
	err := query.
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Order("published_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return posts, total, err
}

// Search performs full-text search on posts
func (r *postRepository) Search(ctx context.Context, query string, page, limit int) ([]*entity.Post, int64, error) {
	var posts []*entity.Post
	var total int64

	offset := (page - 1) * limit

	// PostgreSQL full-text search
	searchQuery := r.db.WithContext(ctx).
		Where("status = ?", "published").
		Where("published_at <= ?", time.Now()).
		Where("title ILIKE ? OR content ILIKE ? OR excerpt ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%")

	// Count total
	if err := searchQuery.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get posts
	err := searchQuery.
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Order("published_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return posts, total, err
}

// IncrementViewCount increments the view count of a post
func (r *postRepository) IncrementViewCount(ctx context.Context, postID uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Where("id = ?", postID).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}

// HasViewedRecently checks if an IP has viewed a post in the last 24 hours
func (r *postRepository) HasViewedRecently(ctx context.Context, postID uint, ipAddress string) (bool, error) {
	var count int64
	twentyFourHoursAgo := time.Now().Add(-24 * time.Hour)

	err := r.db.WithContext(ctx).
		Model(&entity.ViewLog{}).
		Where("post_id = ? AND ip_address = ? AND created_at > ?", postID, ipAddress, twentyFourHoursAgo).
		Count(&count).Error

	return count > 0, err
}

// RecordView records a view log entry
func (r *postRepository) RecordView(ctx context.Context, postID uint, ipAddress, userAgent string) error {
	viewLog := &entity.ViewLog{
		PostID:    postID,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
	return r.db.WithContext(ctx).Create(viewLog).Error
}

// AddLike adds a like to a post
func (r *postRepository) AddLike(ctx context.Context, postID, userID uint) error {
	like := &entity.Like{
		UserID: userID,
		PostID: &postID,
	}
	return r.db.WithContext(ctx).Create(like).Error
}

// RemoveLike removes a like from a post
func (r *postRepository) RemoveLike(ctx context.Context, postID, userID uint) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Delete(&entity.Like{}).Error
}

// HasLiked checks if a user has liked a post
func (r *postRepository) HasLiked(ctx context.Context, postID, userID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.Like{}).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Count(&count).Error
	return count > 0, err
}

// IncrementLikeCount increments the like count of a post
func (r *postRepository) IncrementLikeCount(ctx context.Context, postID uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Where("id = ?", postID).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
}

// DecrementLikeCount decrements the like count of a post
func (r *postRepository) DecrementLikeCount(ctx context.Context, postID uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Where("id = ?", postID).
		UpdateColumn("like_count", gorm.Expr("like_count - ?", 1)).Error
}

// AddBookmark adds a bookmark to a post
func (r *postRepository) AddBookmark(ctx context.Context, postID, userID uint) error {
	bookmark := &entity.Bookmark{
		UserID: userID,
		PostID: postID,
	}
	return r.db.WithContext(ctx).Create(bookmark).Error
}

// RemoveBookmark removes a bookmark from a post
func (r *postRepository) RemoveBookmark(ctx context.Context, postID, userID uint) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Delete(&entity.Bookmark{}).Error
}

// HasBookmarked checks if a user has bookmarked a post
func (r *postRepository) HasBookmarked(ctx context.Context, postID, userID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.Bookmark{}).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Count(&count).Error
	return count > 0, err
}

// IncrementBookmarkCount increments the bookmark count of a post
func (r *postRepository) IncrementBookmarkCount(ctx context.Context, postID uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Where("id = ?", postID).
		UpdateColumn("bookmark_count", gorm.Expr("bookmark_count + ?", 1)).Error
}

// DecrementBookmarkCount decrements the bookmark count of a post
func (r *postRepository) DecrementBookmarkCount(ctx context.Context, postID uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Where("id = ?", postID).
		UpdateColumn("bookmark_count", gorm.Expr("bookmark_count - ?", 1)).Error
}

// ListBookmarkedByUser retrieves bookmarked posts by user
func (r *postRepository) ListBookmarkedByUser(ctx context.Context, userID uint, page, limit int) ([]*entity.Post, int64, error) {
	var posts []*entity.Post
	var total int64

	offset := (page - 1) * limit

	query := r.db.WithContext(ctx).
		Joins("JOIN bookmarks ON bookmarks.post_id = posts.id").
		Where("bookmarks.user_id = ?", userID).
		Where("posts.status = ?", "published")

	// Count total
	if err := query.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get posts
	err := query.
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Order("bookmarks.created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return posts, total, err
}

// SlugExists checks if a slug exists (optionally excluding a specific post ID)
func (r *postRepository) SlugExists(ctx context.Context, slug string, excludeID *uint) (bool, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&entity.Post{}).Where("slug = ?", slug)

	if excludeID != nil {
		query = query.Where("id != ?", *excludeID)
	}

	err := query.Count(&count).Error
	return count > 0, err
}

// IncrementCommentCount increments the comment count of a post
func (r *postRepository) IncrementCommentCount(ctx context.Context, postID uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Where("id = ?", postID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
}

// DecrementCommentCount decrements the comment count of a post
func (r *postRepository) DecrementCommentCount(ctx context.Context, postID uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Where("id = ?", postID).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
}

// GetTotalCount gets the total count of all posts
func (r *postRepository) GetTotalCount(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Post{}).Count(&count).Error
	return count, err
}

// GetPublishedCount gets the count of published posts
func (r *postRepository) GetPublishedCount(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Where("status = ?", "published").
		Where("published_at <= ?", time.Now()).
		Count(&count).Error
	return count, err
}
