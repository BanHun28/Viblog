package entity

import (
	"time"

	"gorm.io/gorm"
)

// Post represents a blog post
type Post struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Content
	Title        string `gorm:"type:varchar(255);not null" json:"title"`
	Slug         string `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Content      string `gorm:"type:text;not null" json:"content"`
	Excerpt      string `gorm:"type:text" json:"excerpt"`
	FeaturedImage *string `gorm:"type:varchar(500)" json:"featured_image,omitempty"`

	// Metadata
	Status        string `gorm:"type:varchar(20);default:'draft'" json:"status"` // draft, published, scheduled
	PublishedAt   *time.Time `json:"published_at,omitempty"`
	ViewCount     int    `gorm:"default:0" json:"view_count"`
	LikeCount     int    `gorm:"default:0" json:"like_count"`
	CommentCount  int    `gorm:"default:0" json:"comment_count"`
	BookmarkCount int    `gorm:"default:0" json:"bookmark_count"`

	// SEO
	MetaTitle       *string `gorm:"type:varchar(255)" json:"meta_title,omitempty"`
	MetaDescription *string `gorm:"type:varchar(500)" json:"meta_description,omitempty"`
	MetaKeywords    *string `gorm:"type:varchar(255)" json:"meta_keywords,omitempty"`

	// Relationships
	AuthorID   uint       `gorm:"not null;index" json:"author_id"`
	Author     *User      `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	CategoryID *uint      `gorm:"index" json:"category_id,omitempty"`
	Category   *Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags       []Tag      `gorm:"many2many:post_tags" json:"tags,omitempty"`
	Comments   []Comment  `gorm:"foreignKey:PostID" json:"comments,omitempty"`
	ViewLogs   []ViewLog  `gorm:"foreignKey:PostID" json:"-"` // For tracking unique views
}

// ViewLog represents a view record for IP-based duplicate prevention
type ViewLog struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	PostID    uint   `gorm:"not null;index:idx_post_ip" json:"post_id"`
	IPAddress string `gorm:"type:varchar(45);not null;index:idx_post_ip" json:"ip_address"` // Supports both IPv4 and IPv6
	UserAgent string `gorm:"type:varchar(500)" json:"user_agent"`
}
