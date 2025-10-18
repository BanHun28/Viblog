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
	ErrTagNotFound     = errors.New("tag not found")
	ErrTagNameRequired = errors.New("tag name is required")
	ErrTagSlugExists   = errors.New("tag slug already exists")
	ErrTagNameExists   = errors.New("tag name already exists")
)

// ListTagsUseCase handles listing all tags
type ListTagsUseCase struct {
	tagRepo repository.TagRepository
}

// NewListTagsUseCase creates a new ListTagsUseCase
func NewListTagsUseCase(tagRepo repository.TagRepository) *ListTagsUseCase {
	return &ListTagsUseCase{
		tagRepo: tagRepo,
	}
}

// Execute retrieves all tags
func (uc *ListTagsUseCase) Execute(ctx context.Context) ([]entity.Tag, error) {
	return uc.tagRepo.FindAll(ctx)
}

// CreateTagInput represents input for creating a tag
type CreateTagInput struct {
	Name string
}

// CreateTagUseCase handles creating a new tag
type CreateTagUseCase struct {
	tagRepo repository.TagRepository
}

// NewCreateTagUseCase creates a new CreateTagUseCase
func NewCreateTagUseCase(tagRepo repository.TagRepository) *CreateTagUseCase {
	return &CreateTagUseCase{
		tagRepo: tagRepo,
	}
}

// Execute creates a new tag
func (uc *CreateTagUseCase) Execute(ctx context.Context, input CreateTagInput) (*entity.Tag, error) {
	// Validate input
	if strings.TrimSpace(input.Name) == "" {
		return nil, ErrTagNameRequired
	}

	// Check if tag name already exists
	exists, err := uc.tagRepo.ExistsByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrTagNameExists
	}

	// Generate slug from name
	tagSlug := slug.Make(input.Name)

	// Check if slug already exists
	slugExists, err := uc.tagRepo.ExistsBySlug(ctx, tagSlug)
	if err != nil {
		return nil, err
	}
	if slugExists {
		return nil, ErrTagSlugExists
	}

	// Create tag
	tag := &entity.Tag{
		Name: input.Name,
		Slug: tagSlug,
	}

	if err := uc.tagRepo.Create(ctx, tag); err != nil {
		return nil, err
	}

	return tag, nil
}

// UpdateTagInput represents input for updating a tag
type UpdateTagInput struct {
	Name *string
}

// UpdateTagUseCase handles updating a tag
type UpdateTagUseCase struct {
	tagRepo repository.TagRepository
}

// NewUpdateTagUseCase creates a new UpdateTagUseCase
func NewUpdateTagUseCase(tagRepo repository.TagRepository) *UpdateTagUseCase {
	return &UpdateTagUseCase{
		tagRepo: tagRepo,
	}
}

// Execute updates a tag
func (uc *UpdateTagUseCase) Execute(ctx context.Context, tagID uint, input UpdateTagInput) (*entity.Tag, error) {
	// Find the tag
	tag, err := uc.tagRepo.FindByID(ctx, tagID)
	if err != nil {
		return nil, ErrTagNotFound
	}

	// Update fields if provided
	if input.Name != nil && strings.TrimSpace(*input.Name) != "" {
		// Check if new name already exists (excluding current tag)
		exists, err := uc.tagRepo.ExistsByName(ctx, *input.Name)
		if err != nil {
			return nil, err
		}
		if exists && *input.Name != tag.Name {
			return nil, ErrTagNameExists
		}

		tag.Name = *input.Name
		tag.Slug = slug.Make(*input.Name)
	}

	// Update tag
	if err := uc.tagRepo.Update(ctx, tag); err != nil {
		return nil, err
	}

	return tag, nil
}

// DeleteTagUseCase handles deleting a tag
type DeleteTagUseCase struct {
	tagRepo repository.TagRepository
}

// NewDeleteTagUseCase creates a new DeleteTagUseCase
func NewDeleteTagUseCase(tagRepo repository.TagRepository) *DeleteTagUseCase {
	return &DeleteTagUseCase{
		tagRepo: tagRepo,
	}
}

// Execute deletes a tag
func (uc *DeleteTagUseCase) Execute(ctx context.Context, tagID uint) error {
	// Check if tag exists
	_, err := uc.tagRepo.FindByID(ctx, tagID)
	if err != nil {
		return ErrTagNotFound
	}

	// Delete the tag
	return uc.tagRepo.Delete(ctx, tagID)
}
