package repository

import (
	"context"
	"errors"

	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
	"gorm.io/gorm"
)

// categoryRepository implements the CategoryRepository interface
type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{db: db}
}

// Create creates a new category
func (r *categoryRepository) Create(ctx context.Context, category *entity.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

// FindByID finds a category by ID
func (r *categoryRepository) FindByID(ctx context.Context, id uint) (*entity.Category, error) {
	var category entity.Category
	err := r.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

// FindBySlug finds a category by slug
func (r *categoryRepository) FindBySlug(ctx context.Context, slug string) (*entity.Category, error) {
	var category entity.Category
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

// FindAll retrieves all categories
func (r *categoryRepository) FindAll(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.WithContext(ctx).Order("name ASC").Find(&categories).Error
	return categories, err
}

// Update updates a category
func (r *categoryRepository) Update(ctx context.Context, category *entity.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

// Delete deletes a category (soft delete)
func (r *categoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Category{}, id).Error
}

// ExistsBySlug checks if a category exists with the given slug
func (r *categoryRepository) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("slug = ?", slug).Count(&count).Error
	return count > 0, err
}

// ExistsByName checks if a category exists with the given name
func (r *categoryRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

// IncrementPostCount increments the post count for a category
func (r *categoryRepository) IncrementPostCount(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Category{}).
		Where("id = ?", id).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// DecrementPostCount decrements the post count for a category
func (r *categoryRepository) DecrementPostCount(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.Category{}).
		Where("id = ?", id).
		UpdateColumn("post_count", gorm.Expr("post_count - ?", 1)).Error
}
