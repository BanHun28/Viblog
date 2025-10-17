package entity

import (
	"time"

	"gorm.io/gorm"
)

// Category represents a blog post category
type Category struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Content
	Name        string  `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Slug        string  `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Metadata
	PostCount int `gorm:"default:0" json:"post_count"`

	// Relationships
	Posts []Post `gorm:"foreignKey:CategoryID" json:"posts,omitempty"`
}
