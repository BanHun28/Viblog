package entity

import (
	"time"
)

// Bookmark represents a user's bookmarked post
type Bookmark struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	// User who bookmarked
	UserID uint  `gorm:"not null;uniqueIndex:idx_user_post" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// Bookmarked post
	PostID uint  `gorm:"not null;uniqueIndex:idx_user_post" json:"post_id"`
	Post   *Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
}
