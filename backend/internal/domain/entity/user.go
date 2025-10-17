package entity

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user account
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Authentication
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"` // Hashed password

	// Profile
	Nickname  string  `gorm:"type:varchar(100);uniqueIndex;not null" json:"nickname"`
	AvatarURL *string `gorm:"type:varchar(500)" json:"avatar_url,omitempty"`
	Bio       *string `gorm:"type:text" json:"bio,omitempty"`

	// Role
	IsAdmin bool `gorm:"default:false" json:"is_admin"`

	// Metadata
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`

	// Relationships
	Posts         []Post         `gorm:"foreignKey:AuthorID" json:"posts,omitempty"`
	Comments      []Comment      `gorm:"foreignKey:UserID" json:"comments,omitempty"`
	Notifications []Notification `gorm:"foreignKey:UserID" json:"notifications,omitempty"`
}
