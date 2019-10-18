package logger

import (
	"context"
	"go.uber.org/zap"
)

// WithLogger set logger
func WithLogger(ctx context.Context, logger zap.Logger) context.Context {
	return context.WithValue(ctx, KeyLogger, logger)
}

// GetLogger return logger
func GetLogger(ctx context.Context) zap.Logger {
	return ctx.Value(KeyLogger).(zap.Logger)
}
