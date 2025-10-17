package entity

import (
	"time"

	"gorm.io/gorm"
)

// Tag represents a blog post tag
type Tag struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Content
	Name string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Slug string `gorm:"type:varchar(50);uniqueIndex;not null" json:"slug"`

	// Metadata
	PostCount int `gorm:"default:0" json:"post_count"`

	// Relationships
	Posts []Post `gorm:"many2many:post_tags" json:"posts,omitempty"`
}
