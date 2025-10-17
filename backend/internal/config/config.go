package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds all application configuration
type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
	JWT        JWTConfig
	CORS       CORSConfig
	RateLimit  RateLimitConfig
	Logging    LoggingConfig
	Monitoring MonitoringConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port     string
	Env      string
	Timezone string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// JWTConfig holds JWT-related configuration
type JWTConfig struct {
	Secret              string
	RefreshSecret       string
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
}

// CORSConfig holds CORS-related configuration
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	APIRequests     int
	APIWindow       time.Duration
	CommentRequests int
	CommentWindow   time.Duration
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string
	Format string
	Output string
}

// MonitoringConfig holds monitoring configuration
type MonitoringConfig struct {
	Enabled bool
	Port    string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port:     getEnv("SERVER_PORT", "30000"),
			Env:      getEnv("SERVER_ENV", "development"),
			Timezone: getEnv("SERVER_TIMEZONE", "Asia/Seoul"),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "30002"),
			User:            getEnv("DB_USER", "viblog"),
			Password:        getEnv("DB_PASSWORD", "viblog123"),
			DBName:          getEnv("DB_NAME", "viblog_dev"),
			SSLMode:         getEnv("DB_SSLMODE", "disable"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		JWT: JWTConfig{
			Secret:              getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			RefreshSecret:       getEnv("JWT_REFRESH_SECRET", "your-refresh-secret-key-change-in-production"),
			AccessTokenExpires:  getEnvAsDuration("JWT_ACCESS_TOKEN_EXPIRES", 15*time.Minute),
			RefreshTokenExpires: getEnvAsDuration("JWT_REFRESH_TOKEN_EXPIRES", 168*time.Hour),
		},
		CORS: CORSConfig{
			AllowedOrigins:   getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:30001"}),
			AllowedMethods:   getEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
			AllowedHeaders:   getEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"Origin", "Content-Type", "Accept", "Authorization"}),
			AllowCredentials: getEnvAsBool("CORS_ALLOW_CREDENTIALS", true),
		},
		RateLimit: RateLimitConfig{
			APIRequests:     getEnvAsInt("RATELIMIT_API_REQUESTS", 100),
			APIWindow:       getEnvAsDuration("RATELIMIT_API_WINDOW", 1*time.Minute),
			CommentRequests: getEnvAsInt("RATELIMIT_COMMENT_REQUESTS", 5),
			CommentWindow:   getEnvAsDuration("RATELIMIT_COMMENT_WINDOW", 1*time.Minute),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "debug"),
			Format: getEnv("LOG_FORMAT", "json"),
			Output: getEnv("LOG_OUTPUT", "stdout"),
		},
		Monitoring: MonitoringConfig{
			Enabled: getEnvAsBool("METRICS_ENABLED", true),
			Port:    getEnv("METRICS_PORT", "30003"),
		},
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Server.Port == "" {
		return fmt.Errorf("SERVER_PORT is required")
	}
	if c.Database.Host == "" {
		return fmt.Errorf("DB_HOST is required")
	}
	if c.Database.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	if c.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}
	if c.JWT.RefreshSecret == "" {
		return fmt.Errorf("JWT_REFRESH_SECRET is required")
	}
	return nil
}

// GetDatabaseDSN returns the database connection string
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

// Helper functions

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	var result []string
	for i := 0; i < len(valueStr); {
		end := i
		for end < len(valueStr) && valueStr[end] != ',' {
			end++
		}
		if end > i {
			result = append(result, valueStr[i:end])
		}
		i = end + 1
	}

	if len(result) == 0 {
		return defaultValue
	}
	return result
}
