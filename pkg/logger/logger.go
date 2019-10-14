package logger

import (
	"context"
	"go.uber.org/zap"
)

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, "logger", &logger)
}

func GetLogger(ctx context.Context) *zap.Logger {
	return ctx.Value("logger").(*zap.Logger)
}
