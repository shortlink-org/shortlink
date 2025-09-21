package logger

import (
	"context"
	"log/slog"
)

// Warn ================================================================================================================

func (log *SlogLogger) Warn(msg string, fields ...any) {
	log.logger.Warn(msg, fields...)
}

func (log *SlogLogger) WarnWithContext(ctx context.Context, msg string, fields ...any) {
	log.logWithContext(ctx, slog.LevelWarn, msg, fields...)
}

// Error ===============================================================================================================

func (log *SlogLogger) Error(msg string, fields ...any) {
	log.logger.Error(msg, fields...)
}

func (log *SlogLogger) ErrorWithContext(ctx context.Context, msg string, fields ...any) {
	// Add error: true field to indicate this is an error log
	errorFields := append([]any{"error", true}, fields...)
	log.logWithContext(ctx, slog.LevelError, msg, errorFields...)
}

// Info ================================================================================================================

func (log *SlogLogger) Info(msg string, fields ...any) {
	log.logger.Info(msg, fields...)
}

func (log *SlogLogger) InfoWithContext(ctx context.Context, msg string, fields ...any) {
	log.logWithContext(ctx, slog.LevelInfo, msg, fields...)
}

// Debug ===============================================================================================================

func (log *SlogLogger) Debug(msg string, fields ...any) {
	log.logger.Debug(msg, fields...)
}

func (log *SlogLogger) DebugWithContext(ctx context.Context, msg string, fields ...any) {
	log.logWithContext(ctx, slog.LevelDebug, msg, fields...)
}