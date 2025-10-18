package repository

import (
	"context"
	"errors"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
	"gorm.io/gorm"
)

// tagRepository implements the TagRepository interface
type tagRepository struct {
	db *gorm.DB
}

// NewTagRepository creates a new tag repository
func NewTagRepository(db *gorm.DB) repository.TagRepository {
	return &tagRepository{db: db}
}

// Create creates a new tag
func (r *tagRepository) Create(ctx context.Context, tag *entity.Tag) error {
	return r.db.WithContext(ctx).Create(tag).Error
}

// FindByID finds a tag by ID
func (r *tagRepository) FindByID(ctx context.Context, id uint) (*entity.Tag, error) {
	var tag entity.Tag
	err := r.db.WithContext(ctx).First(&tag, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

// FindBySlug finds a tag by slug
func (r *tagRepository) FindBySlug(ctx context.Context, slug string) (*entity.Tag, error) {
	var tag entity.Tag
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

// FindAll retrieves all tags
func (r *tagRepository) FindAll(ctx context.Context) ([]entity.Tag, error) {
	var tags []entity.Tag
	err := r.db.WithContext(ctx).Order("name ASC").Find(&tags).Error
	return tags, err
}

// Update updates a tag
func (r *tagRepository) Update(ctx context.Context, tag *entity.Tag) error {
	return r.db.WithContext(ctx).Save(tag).Error
}

// Delete deletes a tag (soft delete)
func (r *tagRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Tag{}, id).Error
}

// ExistsBySlug checks if a tag exists with the given slug
func (r *tagRepository) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Tag{}).Where("slug = ?", slug).Count(&count).Error
	return count > 0, err
}

// ExistsByName checks if a tag exists with the given name
func (r *tagRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Tag{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

// IncrementPostCount increments the post count for a tag
func (r *tagRepository) IncrementPostCount(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Tag{}).
		Where("id = ?", id).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// DecrementPostCount decrements the post count for a tag
func (r *tagRepository) DecrementPostCount(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Tag{}).
		Where("id = ?", id).
		UpdateColumn("post_count", gorm.Expr("post_count - ?", 1)).Error
}
