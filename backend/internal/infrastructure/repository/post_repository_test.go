package repository

import (
	"context"
	"testing"
	"time"

	"github.com/yourusername/viblog/internal/domain/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates an in-memory SQLite database for testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Auto-migrate all entities - migrate them separately to avoid index conflicts
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		t.Fatalf("Failed to migrate User: %v", err)
	}
	err = db.AutoMigrate(&entity.Category{})
	if err != nil {
		t.Fatalf("Failed to migrate Category: %v", err)
	}
	err = db.AutoMigrate(&entity.Tag{})
	if err != nil {
		t.Fatalf("Failed to migrate Tag: %v", err)
	}
	err = db.AutoMigrate(&entity.Post{})
	if err != nil {
		t.Fatalf("Failed to migrate Post: %v", err)
	}
	err = db.AutoMigrate(&entity.ViewLog{})
	if err != nil {
		t.Fatalf("Failed to migrate ViewLog: %v", err)
	}
	err = db.AutoMigrate(&entity.Like{})
	if err != nil {
		t.Fatalf("Failed to migrate Like: %v", err)
	}
	err = db.AutoMigrate(&entity.Bookmark{})
	if err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	return db
}

func TestPostRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)
	ctx := context.Background()

	// Create a test user
	user := &entity.User{
		Email:        "test@example.com",
		Password: "hashedpassword",
		Nickname:     "testuser",
	}
	if err := db.Create(user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Create a test post
	now := time.Now()
	post := &entity.Post{
		Title:       "Test Post",
		Slug:        "test-post",
		Content:     "This is a test post content",
		Excerpt:     "Test excerpt",
		Status:      "published",
		PublishedAt: &now,
		AuthorID:    user.ID,
	}

	err := repo.Create(ctx, post)
	if err != nil {
		t.Errorf("Create() error = %v", err)
	}

	if post.ID == 0 {
		t.Error("Expected post ID to be set after creation")
	}
}

func TestPostRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)
	ctx := context.Background()

	// Create test user
	user := &entity.User{
		Email:        "test@example.com",
		Password: "hashedpassword",
		Nickname:     "testuser",
	}
	db.Create(user)

	// Create test post
	now := time.Now()
	post := &entity.Post{
		Title:       "Test Post",
		Slug:        "test-post",
		Content:     "Content",
		Status:      "published",
		PublishedAt: &now,
		AuthorID:    user.ID,
	}
	db.Create(post)

	// Test GetByID
	retrieved, err := repo.GetByID(ctx, post.ID)
	if err != nil {
		t.Errorf("GetByID() error = %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected post to be retrieved")
	}

	if retrieved.Title != post.Title {
		t.Errorf("Expected title %s, got %s", post.Title, retrieved.Title)
	}
}

func TestPostRepository_GetBySlug(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)
	ctx := context.Background()

	// Create test user
	user := &entity.User{
		Email:        "test@example.com",
		Password: "hashedpassword",
		Nickname:     "testuser",
	}
	db.Create(user)

	// Create test post
	now := time.Now()
	post := &entity.Post{
		Title:       "Test Post",
		Slug:        "test-post-slug",
		Content:     "Content",
		Status:      "published",
		PublishedAt: &now,
		AuthorID:    user.ID,
	}
	db.Create(post)

	// Test GetBySlug
	retrieved, err := repo.GetBySlug(ctx, "test-post-slug")
	if err != nil {
		t.Errorf("GetBySlug() error = %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected post to be retrieved")
	}

	if retrieved.Slug != "test-post-slug" {
		t.Errorf("Expected slug test-post-slug, got %s", retrieved.Slug)
	}
}

func TestPostRepository_ListPublished(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)
	ctx := context.Background()

	// Create test user
	user := &entity.User{
		Email:        "test@example.com",
		Password: "hashedpassword",
		Nickname:     "testuser",
	}
	db.Create(user)

	// Create multiple test posts
	now := time.Now()
	posts := []*entity.Post{
		{
			Title:       "Published Post 1",
			Slug:        "published-1",
			Content:     "Content 1",
			Status:      "published",
			PublishedAt: &now,
			AuthorID:    user.ID,
		},
		{
			Title:       "Published Post 2",
			Slug:        "published-2",
			Content:     "Content 2",
			Status:      "published",
			PublishedAt: &now,
			AuthorID:    user.ID,
		},
		{
			Title:    "Draft Post",
			Slug:     "draft-1",
			Content:  "Content 3",
			Status:   "draft",
			AuthorID: user.ID,
		},
	}

	for _, p := range posts {
		db.Create(p)
	}

	// Test ListPublished
	retrieved, total, err := repo.ListPublished(ctx, 1, 10)
	if err != nil {
		t.Errorf("ListPublished() error = %v", err)
	}

	if total != 2 {
		t.Errorf("Expected 2 published posts, got %d", total)
	}

	if len(retrieved) != 2 {
		t.Errorf("Expected 2 posts in result, got %d", len(retrieved))
	}
}

func TestPostRepository_SlugExists(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)
	ctx := context.Background()

	// Create test user
	user := &entity.User{
		Email:        "test@example.com",
		Password: "hashedpassword",
		Nickname:     "testuser",
	}
	db.Create(user)

	// Create test post
	now := time.Now()
	post := &entity.Post{
		Title:       "Test Post",
		Slug:        "existing-slug",
		Content:     "Content",
		Status:      "published",
		PublishedAt: &now,
		AuthorID:    user.ID,
	}
	db.Create(post)

	// Test existing slug
	exists, err := repo.SlugExists(ctx, "existing-slug", nil)
	if err != nil {
		t.Errorf("SlugExists() error = %v", err)
	}
	if !exists {
		t.Error("Expected slug to exist")
	}

	// Test non-existing slug
	exists, err = repo.SlugExists(ctx, "non-existing-slug", nil)
	if err != nil {
		t.Errorf("SlugExists() error = %v", err)
	}
	if exists {
		t.Error("Expected slug to not exist")
	}

	// Test with exclude ID
	exists, err = repo.SlugExists(ctx, "existing-slug", &post.ID)
	if err != nil {
		t.Errorf("SlugExists() error = %v", err)
	}
	if exists {
		t.Error("Expected slug to not exist when excluded")
	}
}

func TestPostRepository_IncrementViewCount(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)
	ctx := context.Background()

	// Create test user
	user := &entity.User{
		Email:        "test@example.com",
		Password: "hashedpassword",
		Nickname:     "testuser",
	}
	db.Create(user)

	// Create test post
	now := time.Now()
	post := &entity.Post{
		Title:       "Test Post",
		Slug:        "test-post",
		Content:     "Content",
		Status:      "published",
		PublishedAt: &now,
		AuthorID:    user.ID,
		ViewCount:   0,
	}
	db.Create(post)

	// Increment view count
	err := repo.IncrementViewCount(ctx, post.ID)
	if err != nil {
		t.Errorf("IncrementViewCount() error = %v", err)
	}

	// Verify view count increased
	var updated entity.Post
	db.First(&updated, post.ID)
	if updated.ViewCount != 1 {
		t.Errorf("Expected view count 1, got %d", updated.ViewCount)
	}
}

func TestPostRepository_LikeOperations(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)
	ctx := context.Background()

	// Create test user
	user := &entity.User{
		Email:        "test@example.com",
		Password: "hashedpassword",
		Nickname:     "testuser",
	}
	db.Create(user)

	// Create test post
	now := time.Now()
	post := &entity.Post{
		Title:       "Test Post",
		Slug:        "test-post",
		Content:     "Content",
		Status:      "published",
		PublishedAt: &now,
		AuthorID:    user.ID,
		LikeCount:   0,
	}
	db.Create(post)

	// Test AddLike
	err := repo.AddLike(ctx, post.ID, user.ID)
	if err != nil {
		t.Errorf("AddLike() error = %v", err)
	}

	// Test HasLiked
	liked, err := repo.HasLiked(ctx, post.ID, user.ID)
	if err != nil {
		t.Errorf("HasLiked() error = %v", err)
	}
	if !liked {
		t.Error("Expected user to have liked the post")
	}

	// Test IncrementLikeCount
	err = repo.IncrementLikeCount(ctx, post.ID)
	if err != nil {
		t.Errorf("IncrementLikeCount() error = %v", err)
	}

	var updated entity.Post
	db.First(&updated, post.ID)
	if updated.LikeCount != 1 {
		t.Errorf("Expected like count 1, got %d", updated.LikeCount)
	}

	// Test RemoveLike
	err = repo.RemoveLike(ctx, post.ID, user.ID)
	if err != nil {
		t.Errorf("RemoveLike() error = %v", err)
	}

	liked, err = repo.HasLiked(ctx, post.ID, user.ID)
	if err != nil {
		t.Errorf("HasLiked() error = %v", err)
	}
	if liked {
		t.Error("Expected user to have unliked the post")
	}
}

func TestPostRepository_BookmarkOperations(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)
	ctx := context.Background()

	// Create test user
	user := &entity.User{
		Email:        "test@example.com",
		Password: "hashedpassword",
		Nickname:     "testuser",
	}
	db.Create(user)

	// Create test post
	now := time.Now()
	post := &entity.Post{
		Title:         "Test Post",
		Slug:          "test-post",
		Content:       "Content",
		Status:        "published",
		PublishedAt:   &now,
		AuthorID:      user.ID,
		BookmarkCount: 0,
	}
	db.Create(post)

	// Test AddBookmark
	err := repo.AddBookmark(ctx, post.ID, user.ID)
	if err != nil {
		t.Errorf("AddBookmark() error = %v", err)
	}

	// Test HasBookmarked
	bookmarked, err := repo.HasBookmarked(ctx, post.ID, user.ID)
	if err != nil {
		t.Errorf("HasBookmarked() error = %v", err)
	}
	if !bookmarked {
		t.Error("Expected user to have bookmarked the post")
	}

	// Test IncrementBookmarkCount
	err = repo.IncrementBookmarkCount(ctx, post.ID)
	if err != nil {
		t.Errorf("IncrementBookmarkCount() error = %v", err)
	}

	var updated entity.Post
	db.First(&updated, post.ID)
	if updated.BookmarkCount != 1 {
		t.Errorf("Expected bookmark count 1, got %d", updated.BookmarkCount)
	}

	// Test RemoveBookmark
	err = repo.RemoveBookmark(ctx, post.ID, user.ID)
	if err != nil {
		t.Errorf("RemoveBookmark() error = %v", err)
	}

	bookmarked, err = repo.HasBookmarked(ctx, post.ID, user.ID)
	if err != nil {
		t.Errorf("HasBookmarked() error = %v", err)
	}
	if bookmarked {
		t.Error("Expected user to have removed bookmark")
	}
}
