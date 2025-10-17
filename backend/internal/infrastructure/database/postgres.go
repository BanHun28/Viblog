package database

import (
	"fmt"
	"time"

	"github.com/yourusername/viblog/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDatabaseDSN()

	// Configure GORM logger
	var logLevel logger.LogLevel
	switch cfg.Logging.Level {
	case "debug":
		logLevel = logger.Info
	case "info":
		logLevel = logger.Warn
	default:
		logLevel = logger.Error
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying SQL DB for connection pool configuration
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
