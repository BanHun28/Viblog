package repository

import (
	"context"
	"errors"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
	"gorm.io/gorm"
)

// commentRepository implements the CommentRepository interface
type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository creates a new comment repository
func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{db: db}
}

// Create creates a new comment
func (r *commentRepository) Create(ctx context.Context, comment *entity.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

// FindByID finds a comment by ID
func (r *commentRepository) FindByID(ctx context.Context, id uint) (*entity.Comment, error) {
	var comment entity.Comment
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Post").
		First(&comment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

// FindByPostID retrieves all comments for a specific post
func (r *commentRepository) FindByPostID(ctx context.Context, postID uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.WithContext(ctx).
		Where("post_id = ?", postID).
		Preload("User").
		Order("created_at ASC").
		Find(&comments).Error
	return comments, err
}

// FindReplies retrieves all replies to a specific comment
func (r *commentRepository) FindReplies(ctx context.Context, parentID uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Preload("User").
		Order("created_at ASC").
		Find(&comments).Error
	return comments, err
}

// FindAll retrieves all comments with pagination
func (r *commentRepository) FindAll(ctx context.Context, page, limit int) ([]entity.Comment, int64, error) {
	var comments []entity.Comment
	var total int64

	offset := (page - 1) * limit

	// Count total
	if err := r.db.WithContext(ctx).Model(&entity.Comment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get comments with pagination
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Post").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&comments).Error

	return comments, total, err
}

// Update updates a comment
func (r *commentRepository) Update(ctx context.Context, comment *entity.Comment) error {
	return r.db.WithContext(ctx).Save(comment).Error
}

// Delete deletes a comment (soft delete)
func (r *commentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Comment{}, id).Error
}

// IncrementLikeCount increments the like count for a comment
func (r *commentRepository) IncrementLikeCount(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Comment{}).
		Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
}

// DecrementLikeCount decrements the like count for a comment
func (r *commentRepository) DecrementLikeCount(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Comment{}).
		Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count - ?", 1)).Error
}

// GetTotalCount gets total count of all comments
func (r *commentRepository) GetTotalCount(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Comment{}).Count(&count).Error
	return count, err
}
