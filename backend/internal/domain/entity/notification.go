package entity

import (
	"time"

	"gorm.io/gorm"
)

// NotificationType defines the type of notification
type NotificationType string

const (
	NotificationTypeCommentReply   NotificationType = "comment_reply"   // 댓글에 대한 답글
	NotificationTypePostComment    NotificationType = "post_comment"    // 내 글에 댓글
	NotificationTypeCommentLike    NotificationType = "comment_like"    // 댓글 좋아요
	NotificationTypePostLike       NotificationType = "post_like"       // 글 좋아요
	NotificationTypePostBookmark   NotificationType = "post_bookmark"   // 글 북마크
)

// Notification represents a user notification
type Notification struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Target user
	UserID uint  `gorm:"not null;index" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// Notification details
	Type    NotificationType `gorm:"type:varchar(50);not null" json:"type"`
	Title   string           `gorm:"type:varchar(255);not null" json:"title"`
	Message string           `gorm:"type:text;not null" json:"message"`

	// Relationships - nullable for flexibility
	PostID    *uint    `gorm:"index" json:"post_id,omitempty"`
	Post      *Post    `gorm:"foreignKey:PostID" json:"post,omitempty"`
	CommentID *uint    `gorm:"index" json:"comment_id,omitempty"`
	Comment   *Comment `gorm:"foreignKey:CommentID" json:"comment,omitempty"`
	ActorID   *uint    `gorm:"index" json:"actor_id,omitempty"` // User who triggered the notification
	Actor     *User    `gorm:"foreignKey:ActorID" json:"actor,omitempty"`

	// Status
	IsRead bool       `gorm:"default:false;index" json:"is_read"`
	ReadAt *time.Time `json:"read_at,omitempty"`

	// Link to the resource
	Link string `gorm:"type:varchar(500)" json:"link"`
}
