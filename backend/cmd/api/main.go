package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yourusername/viblog/internal/config"
	"github.com/yourusername/viblog/internal/infrastructure/logger"
	"github.com/yourusername/viblog/internal/interface/http/handler"
	"github.com/yourusername/viblog/internal/interface/http/router"
	"go.uber.org/zap"

	_ "github.com/yourusername/viblog/docs" // Import generated docs
)

// @title Viblog API
// @version 1.0
// @description Personal Blog Platform API with Clean Architecture
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.viblog.io/support
// @contact.email support@viblog.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:30000
// @BasePath /api/v1

// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	zapLogger, err := logger.NewZapLogger(cfg.Logging)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer zapLogger.Sync()

	zapLogger.Info("Starting Viblog API server",
		zap.String("env", cfg.Server.Env),
		zap.String("port", cfg.Server.Port),
	)

	// TODO: Initialize database connection
	// db, err := database.NewPostgresDB(cfg)
	// if err != nil {
	// 	zapLogger.Fatal("Failed to connect to database", zap.Error(err))
	// }

	// TODO: Run migrations
	// if err := database.RunMigrations(db); err != nil {
	// 	zapLogger.Fatal("Failed to run migrations", zap.Error(err))
	// }

	// TODO: Initialize repositories
	// userRepo := repository.NewUserRepository(db)
	// postRepo := repository.NewPostRepository(db)
	// commentRepo := repository.NewCommentRepository(db)

	// TODO: Initialize use cases
	// userUC := usecase.NewUserUseCase(userRepo)
	// postUC := usecase.NewPostUseCase(postRepo)
	// commentUC := usecase.NewCommentUseCase(commentRepo)

	// Initialize handlers (temporary stubs until use cases are implemented)
	userHandler := handler.NewUserHandler(nil)
	postHandler := handler.NewPostHandler(nil)
	commentHandler := handler.NewCommentHandler(nil)
	adminHandler := handler.NewAdminHandler(nil)
	notificationHandler := handler.NewNotificationHandler(nil)

	// Initialize router
	r := router.New(
		cfg,
		zapLogger,
		userHandler,
		postHandler,
		commentHandler,
		adminHandler,
		notificationHandler,
	)

	// Setup routes
	engine := r.Setup()

	// Create HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		zapLogger.Info("Server listening", zap.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zapLogger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zapLogger.Info("Shutting down server...")

	// Graceful shutdown with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zapLogger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	zapLogger.Info("Server exited")
}
