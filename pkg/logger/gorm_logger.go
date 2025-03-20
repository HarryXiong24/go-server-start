package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

// GormLogger implements GORM's logger.Interface with zap
type GormLogger struct {
	ZapLogger *zap.Logger
	LogLevel  logger.LogLevel
}

// NewGormLogger creates a new GORM logger using zap
func NewGormLogger() logger.Interface {
	return &GormLogger{
		ZapLogger: Logger,
		LogLevel:  logger.Info,
	}
}

// LogMode sets the log level
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs info messages
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.ZapLogger.Info(fmt.Sprintf(msg, data...))
	}
}

// Warn logs warn messages
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.ZapLogger.Warn(fmt.Sprintf(msg, data...))
	}
}

// Error logs error messages
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.ZapLogger.Error(fmt.Sprintf(msg, data...))
	}
}

// Trace logs SQL and execution time
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	fields := []zap.Field{
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.Duration("elapsed", elapsed),
	}

	switch {
	case err != nil && l.LogLevel >= logger.Error:
		l.ZapLogger.Error("GORM Error", append(fields, zap.Error(err))...)
	case elapsed > 200*time.Millisecond && l.LogLevel >= logger.Warn:
		l.ZapLogger.Warn("GORM Slow SQL", fields...)
	case l.LogLevel >= logger.Info:
		l.ZapLogger.Info("GORM SQL", fields...)
	}
}
