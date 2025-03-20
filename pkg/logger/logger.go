package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the global logger
var Logger *zap.Logger

// Sugar is the global sugared logger
var Sugar *zap.SugaredLogger

// Init initializes the logger
func Init(level string) error {
	var zapLevel zapcore.Level

	// Parse log level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	// Configure logger
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Build logger
	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	// Create sugared logger
	Sugar = Logger.Sugar()

	return nil
}

// Sync flushes the logger
func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
}

// GetLogger returns the global logger instance
func GetLogger() *zap.Logger {
	return Logger
}

// Info wrapper
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Error wrapper
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Debug wrapper
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Warn wrapper
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}
