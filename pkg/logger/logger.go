package logger

import (
	"context"
	"go.uber.org/zap"
)

// WithLogger set logger
func WithLogger(ctx context.Context, logger zap.Logger) context.Context { //nolint unused
	return context.WithValue(ctx, keyLogger, logger)
}

// GetLogger return logger
func GetLogger(ctx context.Context) zap.Logger { //nolint unused
	return ctx.Value(keyLogger).(zap.Logger)
}
