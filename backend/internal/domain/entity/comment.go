package entity

import (
	"time"

	"gorm.io/gorm"
)

// Comment represents a comment on a blog post
type Comment struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Content
	Content string `gorm:"type:text;not null" json:"content"`

	// Relationships
	PostID   uint  `gorm:"not null;index" json:"post_id"`
	Post     *Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
	ParentID *uint `gorm:"index" json:"parent_id,omitempty"` // For nested comments/replies
	Parent   *Comment `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Replies  []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`

	// Author can be authenticated user or anonymous
	UserID       *uint   `gorm:"index" json:"user_id,omitempty"`
	User         *User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	AuthorName   *string `gorm:"type:varchar(100)" json:"author_name,omitempty"`   // For anonymous comments
	AuthorEmail  *string `gorm:"type:varchar(255)" json:"author_email,omitempty"`  // For anonymous comments
	AuthorPassword *string `gorm:"type:varchar(255)" json:"-"` // Hashed password for anonymous comment modification

	// Metadata
	LikeCount int  `gorm:"default:0" json:"like_count"`
	IsEdited  bool `gorm:"default:false" json:"is_edited"`
}

// IsAnonymous checks if the comment is from an anonymous user
func (c *Comment) IsAnonymous() bool {
	return c.UserID == nil
}

// IsReply checks if the comment is a reply to another comment
func (c *Comment) IsReply() bool {
	return c.ParentID != nil
}
