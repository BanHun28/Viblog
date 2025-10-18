//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/yourusername/viblog/internal/config"
	"github.com/yourusername/viblog/internal/infrastructure/auth"
	"github.com/yourusername/viblog/internal/infrastructure/database"
	"github.com/yourusername/viblog/internal/infrastructure/logger"
	"github.com/yourusername/viblog/internal/infrastructure/repository"
	"github.com/yourusername/viblog/internal/interface/http/handler"
	"github.com/yourusername/viblog/internal/interface/http/router"
	"github.com/yourusername/viblog/internal/usecase/post"
	"github.com/yourusername/viblog/internal/usecase/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InitializeApp initializes the application with all dependencies
func InitializeApp(cfg *config.Config) (*router.Router, func(), error) {
	wire.Build(
		// Infrastructure
		provideLogger,
		provideDatabase,
		provideJWTService,

		// Repositories
		repository.NewUserRepository,
		repository.NewPostRepository,

		// Use Cases
		user.NewRegisterUseCase,
		user.NewLoginUseCase,
		user.NewGetProfileUseCase,
		user.NewUpdateProfileUseCase,
		post.NewListUseCase,
		post.NewGetUseCase,

		// Handlers
		provideUserHandler,
		providePostHandler,
		provideCommentHandler,
		provideAdminHandler,
		provideNotificationHandler,

		// Router
		router.New,
	)
	return nil, nil, nil
}

// Provider functions

func provideLogger(cfg *config.Config) (*zap.Logger, error) {
	return logger.NewZapLogger(cfg.Logging)
}

func provideDatabase(cfg *config.Config) (*gorm.DB, func(), error) {
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		sqlDB, _ := db.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		cleanup()
		return nil, nil, err
	}

	return db, cleanup, nil
}

func provideJWTService(cfg *config.Config) *auth.JWTService {
	return auth.NewJWTService(cfg.JWT.Secret, cfg.JWT.RefreshSecret)
}

func provideUserHandler(
	registerUC *user.RegisterUseCase,
	loginUC *user.LoginUseCase,
	getProfileUC *user.GetProfileUseCase,
	updateProfileUC *user.UpdateProfileUseCase,
	jwtService *auth.JWTService,
) *handler.UserHandler {
	return handler.NewUserHandler(registerUC, loginUC, getProfileUC, updateProfileUC, jwtService)
}

func providePostHandler(
	listUC *post.ListUseCase,
	getUC *post.GetUseCase,
) *handler.PostHandler {
	return handler.NewPostHandler(listUC, getUC)
}

func provideCommentHandler() *handler.CommentHandler {
	// TODO: Implement comment use cases
	return handler.NewCommentHandler(nil)
}

func provideAdminHandler() *handler.AdminHandler {
	// TODO: Implement admin use cases
	return handler.NewAdminHandler(nil)
}

func provideNotificationHandler() *handler.NotificationHandler {
	// TODO: Implement notification use cases
	return handler.NewNotificationHandler(nil)
}
