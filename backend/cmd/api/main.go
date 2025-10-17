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
	"github.com/yourusername/viblog/wire"

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

	// Initialize application with dependency injection
	r, cleanup, err := wire.InitializeApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer cleanup()

	log.Printf("Starting Viblog API server on port %s (env: %s)", cfg.Server.Port, cfg.Server.Env)

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
		log.Printf("Server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
