package logger

import (
	"os"

	"github.com/yourusername/viblog/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger creates a new Zap logger based on configuration
func NewZapLogger(cfg config.LoggingConfig) (*zap.Logger, error) {
	// Parse log level
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	// Encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Choose encoder based on format
	var encoder zapcore.Encoder
	if cfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Output configuration
	var writeSyncer zapcore.WriteSyncer
	if cfg.Output == "stdout" {
		writeSyncer = zapcore.AddSync(os.Stdout)
	} else {
		// For file output, you can add file handling here
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	// Create core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger, nil
}

// NewDevelopmentLogger creates a development logger with pretty printing
func NewDevelopmentLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

// NewProductionLogger creates a production logger with JSON output
func NewProductionLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
