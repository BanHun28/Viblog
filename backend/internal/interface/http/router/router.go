package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yourusername/viblog/internal/config"
	"github.com/yourusername/viblog/internal/interface/http/handler"
	"github.com/yourusername/viblog/internal/interface/http/middleware"
	"go.uber.org/zap"
)

// Router holds the HTTP router and its dependencies
type Router struct {
	engine *gin.Engine
	cfg    *config.Config
	logger *zap.Logger

	// Handlers
	userHandler         *handler.UserHandler
	postHandler         *handler.PostHandler
	commentHandler      *handler.CommentHandler
	adminHandler        *handler.AdminHandler
	notificationHandler *handler.NotificationHandler
}

// New creates a new HTTP router
func New(
	cfg *config.Config,
	logger *zap.Logger,
	userHandler *handler.UserHandler,
	postHandler *handler.PostHandler,
	commentHandler *handler.CommentHandler,
	adminHandler *handler.AdminHandler,
	notificationHandler *handler.NotificationHandler,
) *Router {
	// Set Gin mode based on environment
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	return &Router{
		engine:              engine,
		cfg:                 cfg,
		logger:              logger,
		userHandler:         userHandler,
		postHandler:         postHandler,
		commentHandler:      commentHandler,
		adminHandler:        adminHandler,
		notificationHandler: notificationHandler,
	}
}

// Setup configures all routes and middleware
func (r *Router) Setup() *gin.Engine {
	// Convert config CORS to middleware CORS
	corsConfig := middleware.CORSConfig{
		AllowedOrigins:   r.cfg.CORS.AllowedOrigins,
		AllowedMethods:   r.cfg.CORS.AllowedMethods,
		AllowedHeaders:   r.cfg.CORS.AllowedHeaders,
		AllowCredentials: r.cfg.CORS.AllowCredentials,
	}

	// Global middleware
	r.engine.Use(middleware.Logger(r.logger))
	r.engine.Use(middleware.Recovery(r.logger))
	r.engine.Use(middleware.CORS(corsConfig))
	r.engine.Use(middleware.ErrorHandler())

	// Health check endpoint
	r.engine.GET("/health", r.healthCheck)

	// Swagger documentation
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	v1 := r.engine.Group("/api/v1")
	{
		r.setupAuthRoutes(v1)
		r.setupPostRoutes(v1)
		r.setupCommentRoutes(v1)
		r.setupNotificationRoutes(v1)
		r.setupAdminRoutes(v1)
	}

	return r.engine
}

// setupAuthRoutes configures authentication and user routes
func (r *Router) setupAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		// Public routes
		auth.POST("/register", r.userHandler.Register)
		auth.POST("/login", r.userHandler.Login)
		auth.POST("/refresh", r.userHandler.RefreshToken)

		// Protected routes
		protected := auth.Group("")
		jwtConfig := middleware.JWTConfig{Secret: r.cfg.JWT.Secret}
		protected.Use(middleware.AuthMiddleware(jwtConfig))
		{
			protected.POST("/logout", r.userHandler.Logout)
			protected.GET("/me", r.userHandler.GetProfile)
			protected.PUT("/me", r.userHandler.UpdateProfile)
		}
	}
}

// setupPostRoutes configures post-related routes
func (r *Router) setupPostRoutes(rg *gin.RouterGroup) {
	jwtConfig := middleware.JWTConfig{Secret: r.cfg.JWT.Secret}

	posts := rg.Group("/posts")
	{
		// Public routes
		posts.GET("", r.postHandler.List)
		posts.GET("/:id", r.postHandler.Get)
		posts.GET("/search", r.postHandler.Search)
		posts.POST("/:id/view", r.postHandler.IncrementView) // View count tracking

		// Protected routes (admin only for write operations)
		protected := posts.Group("")
		protected.Use(middleware.AuthMiddleware(jwtConfig))
		protected.Use(middleware.RequireAdmin())
		{
			protected.POST("", r.postHandler.Create)
			protected.PUT("/:id", r.postHandler.Update)
			protected.DELETE("/:id", r.postHandler.Delete)
		}

		// Protected routes (authenticated users)
		userProtected := posts.Group("")
		userProtected.Use(middleware.AuthMiddleware(jwtConfig))
		{
			userProtected.POST("/:id/like", r.postHandler.Like)
			userProtected.DELETE("/:id/like", r.postHandler.Unlike)
			userProtected.POST("/:id/bookmark", r.postHandler.Bookmark)
			userProtected.DELETE("/:id/bookmark", r.postHandler.Unbookmark)
		}
	}

	// Categories
	categories := rg.Group("/categories")
	{
		categories.GET("", r.postHandler.ListCategories)
		categories.GET("/:slug/posts", r.postHandler.GetPostsByCategory)
	}

	// Tags
	tags := rg.Group("/tags")
	{
		tags.GET("", r.postHandler.ListTags)
		tags.GET("/:slug/posts", r.postHandler.GetPostsByTag)
	}
}

// setupCommentRoutes configures comment-related routes
func (r *Router) setupCommentRoutes(rg *gin.RouterGroup) {
	jwtConfig := middleware.JWTConfig{Secret: r.cfg.JWT.Secret}

	comments := rg.Group("/comments")
	{
		// Public routes - list comments for a post
		comments.GET("/post/:postId", r.commentHandler.List)

		// Rate-limited routes (both anonymous and authenticated)
		rateLimited := comments.Group("")
		rateLimited.Use(middleware.RateLimit(
			r.cfg.RateLimit.CommentRequests,
			r.cfg.RateLimit.CommentWindow,
		))
		{
			// Optional auth (allows both anonymous and authenticated)
			rateLimited.POST("/post/:postId", middleware.OptionalAuth(jwtConfig), r.commentHandler.Create)
			rateLimited.PUT("/:id", middleware.OptionalAuth(jwtConfig), r.commentHandler.Update)
			rateLimited.DELETE("/:id", middleware.OptionalAuth(jwtConfig), r.commentHandler.Delete)

			// Reply routes (nested comments)
			rateLimited.GET("/:id/replies", r.commentHandler.ListReplies)
			rateLimited.POST("/:id/replies", middleware.OptionalAuth(jwtConfig), r.commentHandler.CreateReply)

			// Authenticated only
			authenticated := rateLimited.Group("")
			authenticated.Use(middleware.AuthMiddleware(jwtConfig))
			{
				authenticated.POST("/:id/like", r.commentHandler.Like)
				authenticated.DELETE("/:id/like", r.commentHandler.Unlike)
			}
		}
	}
}

// setupNotificationRoutes configures notification-related routes
func (r *Router) setupNotificationRoutes(rg *gin.RouterGroup) {
	jwtConfig := middleware.JWTConfig{Secret: r.cfg.JWT.Secret}

	notifications := rg.Group("/notifications")
	notifications.Use(middleware.AuthMiddleware(jwtConfig))
	{
		notifications.GET("", r.notificationHandler.List)
		notifications.GET("/unread", r.notificationHandler.ListUnread)
		notifications.PUT("/:id/read", r.notificationHandler.MarkAsRead)
		notifications.PUT("/read-all", r.notificationHandler.MarkAllAsRead)
	}
}

// setupAdminRoutes configures admin-related routes
func (r *Router) setupAdminRoutes(rg *gin.RouterGroup) {
	jwtConfig := middleware.JWTConfig{Secret: r.cfg.JWT.Secret}

	admin := rg.Group("/admin")
	admin.Use(middleware.AuthMiddleware(jwtConfig))
	admin.Use(middleware.RequireAdmin())
	{
		// Dashboard
		admin.GET("/dashboard", r.adminHandler.GetDashboard)

		// User management
		admin.GET("/users", r.adminHandler.ListUsers)
		admin.DELETE("/users/:id", r.adminHandler.DeleteUser)

		// Comment moderation
		admin.GET("/comments", r.adminHandler.ListComments)
		admin.DELETE("/comments/:id", r.adminHandler.DeleteComment)

		// Category management
		admin.GET("/categories", r.adminHandler.ListCategories)
		admin.POST("/categories", r.adminHandler.CreateCategory)
		admin.PUT("/categories/:id", r.adminHandler.UpdateCategory)
		admin.DELETE("/categories/:id", r.adminHandler.DeleteCategory)

		// Tag management
		admin.GET("/tags", r.adminHandler.ListTags)
		admin.POST("/tags", r.adminHandler.CreateTag)
		admin.PUT("/tags/:id", r.adminHandler.UpdateTag)
		admin.DELETE("/tags/:id", r.adminHandler.DeleteTag)
	}
}

// healthCheck is a simple health check endpoint
func (r *Router) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"env":    r.cfg.Server.Env,
	})
}
