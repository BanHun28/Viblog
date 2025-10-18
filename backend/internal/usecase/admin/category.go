package admin

import (
	"context"
	"errors"
	"strings"

	"github.com/gosimple/slug"
	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/domain/repository"
)

var (
	ErrCategoryNotFound     = errors.New("category not found")
	ErrCategoryNameRequired = errors.New("category name is required")
	ErrCategorySlugExists   = errors.New("category slug already exists")
	ErrCategoryNameExists   = errors.New("category name already exists")
)

// ListCategoriesUseCase handles listing all categories
type ListCategoriesUseCase struct {
	categoryRepo repository.CategoryRepository
}

// NewListCategoriesUseCase creates a new ListCategoriesUseCase
func NewListCategoriesUseCase(categoryRepo repository.CategoryRepository) *ListCategoriesUseCase {
	return &ListCategoriesUseCase{
		categoryRepo: categoryRepo,
	}
}

// Execute retrieves all categories
func (uc *ListCategoriesUseCase) Execute(ctx context.Context) ([]entity.Category, error) {
	return uc.categoryRepo.FindAll(ctx)
}

// CreateCategoryInput represents input for creating a category
type CreateCategoryInput struct {
	Name        string
	Description *string
}

// CreateCategoryUseCase handles creating a new category
type CreateCategoryUseCase struct {
	categoryRepo repository.CategoryRepository
}

// NewCreateCategoryUseCase creates a new CreateCategoryUseCase
func NewCreateCategoryUseCase(categoryRepo repository.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

// Execute creates a new category
func (uc *CreateCategoryUseCase) Execute(ctx context.Context, input CreateCategoryInput) (*entity.Category, error) {
	// Validate input
	if strings.TrimSpace(input.Name) == "" {
		return nil, ErrCategoryNameRequired
	}

	// Check if category name already exists
	exists, err := uc.categoryRepo.ExistsByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrCategoryNameExists
	}

	// Generate slug from name
	categorySlug := slug.Make(input.Name)

	// Check if slug already exists
	slugExists, err := uc.categoryRepo.ExistsBySlug(ctx, categorySlug)
	if err != nil {
		return nil, err
	}
	if slugExists {
		return nil, ErrCategorySlugExists
	}

	// Create category
	category := &entity.Category{
		Name:        input.Name,
		Slug:        categorySlug,
		Description: input.Description,
	}

	if err := uc.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

// UpdateCategoryInput represents input for updating a category
type UpdateCategoryInput struct {
	Name        *string
	Description *string
}

// UpdateCategoryUseCase handles updating a category
type UpdateCategoryUseCase struct {
	categoryRepo repository.CategoryRepository
}

// NewUpdateCategoryUseCase creates a new UpdateCategoryUseCase
func NewUpdateCategoryUseCase(categoryRepo repository.CategoryRepository) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

// Execute updates a category
func (uc *UpdateCategoryUseCase) Execute(ctx context.Context, categoryID uint, input UpdateCategoryInput) (*entity.Category, error) {
	// Find the category
	category, err := uc.categoryRepo.FindByID(ctx, categoryID)
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	// Update fields if provided
	if input.Name != nil && strings.TrimSpace(*input.Name) != "" {
		// Check if new name already exists (excluding current category)
		exists, err := uc.categoryRepo.ExistsByName(ctx, *input.Name)
		if err != nil {
			return nil, err
		}
		if exists && *input.Name != category.Name {
			return nil, ErrCategoryNameExists
		}

		category.Name = *input.Name
		category.Slug = slug.Make(*input.Name)
	}

	if input.Description != nil {
		category.Description = input.Description
	}

	// Update category
	if err := uc.categoryRepo.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategoryUseCase handles deleting a category
type DeleteCategoryUseCase struct {
	categoryRepo repository.CategoryRepository
}

// NewDeleteCategoryUseCase creates a new DeleteCategoryUseCase
func NewDeleteCategoryUseCase(categoryRepo repository.CategoryRepository) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

// Execute deletes a category
func (uc *DeleteCategoryUseCase) Execute(ctx context.Context, categoryID uint) error {
	// Check if category exists
	_, err := uc.categoryRepo.FindByID(ctx, categoryID)
	if err != nil {
		return ErrCategoryNotFound
	}

	// Delete the category
	return uc.categoryRepo.Delete(ctx, categoryID)
}
