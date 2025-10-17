package database

import (
	"github.com/yourusername/viblog/internal/domain/entity"
	"gorm.io/gorm"
)

// RunMigrations runs all database migrations
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Comment{},
		&entity.Category{},
		&entity.Tag{},
		&entity.Like{},
		&entity.Bookmark{},
		&entity.Notification{},
	)
}
