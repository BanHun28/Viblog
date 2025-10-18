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
	"github.com/yourusername/viblog/internal/usecase/admin"
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
		repository.NewCommentRepository,
		repository.NewCategoryRepository,
		repository.NewTagRepository,

		// User Use Cases
		user.NewRegisterUseCase,
		user.NewLoginUseCase,
		user.NewGetProfileUseCase,
		user.NewUpdateProfileUseCase,

		// Admin Use Cases
		admin.NewGetDashboardUseCase,
		admin.NewListUsersUseCase,
		admin.NewDeleteUserUseCase,
		admin.NewListCommentsUseCase,
		admin.NewDeleteCommentUseCase,
		admin.NewListCategoriesUseCase,
		admin.NewCreateCategoryUseCase,
		admin.NewUpdateCategoryUseCase,
		admin.NewDeleteCategoryUseCase,
		admin.NewListTagsUseCase,
		admin.NewCreateTagUseCase,
		admin.NewUpdateTagUseCase,
		admin.NewDeleteTagUseCase,

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

func providePostHandler() *handler.PostHandler {
	// TODO: Implement post use cases
	return handler.NewPostHandler(nil)
}

func provideCommentHandler() *handler.CommentHandler {
	// TODO: Implement comment use cases
	return handler.NewCommentHandler(nil)
}

func provideAdminHandler(
	dashboardUC *admin.GetDashboardUseCase,
	listUsersUC *admin.ListUsersUseCase,
	deleteUserUC *admin.DeleteUserUseCase,
	listCommentsUC *admin.ListCommentsUseCase,
	deleteCommentUC *admin.DeleteCommentUseCase,
	listCategoriesUC *admin.ListCategoriesUseCase,
	createCategoryUC *admin.CreateCategoryUseCase,
	updateCategoryUC *admin.UpdateCategoryUseCase,
	deleteCategoryUC *admin.DeleteCategoryUseCase,
	listTagsUC *admin.ListTagsUseCase,
	createTagUC *admin.CreateTagUseCase,
	updateTagUC *admin.UpdateTagUseCase,
	deleteTagUC *admin.DeleteTagUseCase,
) *handler.AdminHandler {
	return handler.NewAdminHandler(
		dashboardUC,
		listUsersUC,
		deleteUserUC,
		listCommentsUC,
		deleteCommentUC,
		listCategoriesUC,
		createCategoryUC,
		updateCategoryUC,
		deleteCategoryUC,
		listTagsUC,
		createTagUC,
		updateTagUC,
		deleteTagUC,
	)
}

func provideNotificationHandler() *handler.NotificationHandler {
	// TODO: Implement notification use cases
	return handler.NewNotificationHandler(nil)
}
