package entity

import (
	"time"
)

// Like represents a like on a post or comment
type Like struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	// User who liked
	UserID uint  `gorm:"not null;index:idx_user_post,idx_user_comment" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// Target (either post or comment, not both)
	PostID    *uint    `gorm:"index:idx_user_post" json:"post_id,omitempty"`
	Post      *Post    `gorm:"foreignKey:PostID" json:"post,omitempty"`
	CommentID *uint    `gorm:"index:idx_user_comment" json:"comment_id,omitempty"`
	Comment   *Comment `gorm:"foreignKey:CommentID" json:"comment,omitempty"`
}
