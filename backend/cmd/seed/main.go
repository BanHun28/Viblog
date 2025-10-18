package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yourusername/viblog/internal/config"
	"github.com/yourusername/viblog/internal/domain/entity"
	"github.com/yourusername/viblog/internal/infrastructure/database"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	fmt.Println("Creating seed data...")

	// Create admin user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := &entity.User{
		Email:    "admin@viblog.com",
		Password: string(hashedPassword),
		Nickname: "Admin",
		IsAdmin:  true,
	}
	db.Create(admin)
	fmt.Println("✓ Created admin user (email: admin@viblog.com, password: admin123)")

	// Create categories
	categories := []*entity.Category{
		{Name: "Technology", Slug: "technology", Description: strPtr("Technology and programming articles")},
		{Name: "Lifestyle", Slug: "lifestyle", Description: strPtr("Lifestyle and personal development")},
		{Name: "Travel", Slug: "travel", Description: strPtr("Travel guides and experiences")},
	}
	for _, cat := range categories {
		db.Create(cat)
	}
	fmt.Println("✓ Created 3 categories")

	// Create tags
	tags := []*entity.Tag{
		{Name: "Go", Slug: "go"},
		{Name: "Web Development", Slug: "web-development"},
		{Name: "Tutorial", Slug: "tutorial"},
		{Name: "Tips", Slug: "tips"},
		{Name: "Review", Slug: "review"},
	}
	for _, tag := range tags {
		db.Create(tag)
	}
	fmt.Println("✓ Created 5 tags")

	// Create sample posts
	now := time.Now()
	posts := []*entity.Post{
		{
			Title:       "Getting Started with Go",
			Slug:        "getting-started-with-go",
			Content:     "# Getting Started with Go\n\nGo is a powerful programming language...",
			Excerpt:     "Learn the basics of Go programming language",
			Status:      "published",
			PublishedAt: &now,
			AuthorID:    admin.ID,
			CategoryID:  &categories[0].ID,
			ViewCount:   150,
			LikeCount:   25,
		},
		{
			Title:       "Building RESTful APIs with Gin",
			Slug:        "building-restful-apis-with-gin",
			Content:     "# Building RESTful APIs with Gin\n\nGin is a high-performance HTTP web framework...",
			Excerpt:     "Complete guide to building APIs using Gin framework",
			Status:      "published",
			PublishedAt: &now,
			AuthorID:    admin.ID,
			CategoryID:  &categories[0].ID,
			ViewCount:   230,
			LikeCount:   42,
		},
		{
			Title:       "Traveling in Southeast Asia",
			Slug:        "traveling-in-southeast-asia",
			Content:     "# Traveling in Southeast Asia\n\nSoutheast Asia offers incredible experiences...",
			Excerpt:     "A guide to exploring Southeast Asia on a budget",
			Status:      "published",
			PublishedAt: &now,
			AuthorID:    admin.ID,
			CategoryID:  &categories[2].ID,
			ViewCount:   89,
			LikeCount:   15,
		},
		{
			Title:       "Productivity Tips for Developers",
			Slug:        "productivity-tips-for-developers",
			Content:     "# Productivity Tips for Developers\n\nStay productive with these essential tips...",
			Excerpt:     "Boost your productivity as a software developer",
			Status:      "published",
			PublishedAt: &now,
			AuthorID:    admin.ID,
			CategoryID:  &categories[1].ID,
			ViewCount:   178,
			LikeCount:   31,
		},
		{
			Title:       "Understanding Clean Architecture",
			Slug:        "understanding-clean-architecture",
			Content:     "# Understanding Clean Architecture\n\nClean Architecture is a software design philosophy...",
			Excerpt:     "Learn the principles of Clean Architecture",
			Status:      "published",
			PublishedAt: &now,
			AuthorID:    admin.ID,
			CategoryID:  &categories[0].ID,
			ViewCount:   420,
			LikeCount:   78,
		},
	}

	for _, post := range posts {
		db.Create(post)
		// Add tags to posts
		db.Model(post).Association("Tags").Append([]*entity.Tag{tags[0], tags[1], tags[2]})
	}
	fmt.Println("✓ Created 5 sample posts with tags")

	fmt.Println("\n✅ Seed data created successfully!")
	fmt.Println("\n📝 You can now login with:")
	fmt.Println("   Email: admin@viblog.com")
	fmt.Println("   Password: admin123")
}

func strPtr(s string) *string {
	return &s
}
