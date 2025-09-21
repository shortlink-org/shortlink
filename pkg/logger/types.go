package logger

import (
	"context"
	"io"
)

// Logger is our contract for the logger.
type Logger interface {
	Error(msg string, fields ...any)
	ErrorWithContext(ctx context.Context, msg string, fields ...any)

	Warn(msg string, fields ...any)
	WarnWithContext(ctx context.Context, msg string, fields ...any)

	Info(msg string, fields ...any)
	InfoWithContext(ctx context.Context, msg string, fields ...any)

	Debug(msg string, fields ...any)
	DebugWithContext(ctx context.Context, msg string, fields ...any)

	// Closer is the interface that wraps the basic Close method.
	io.Closer
}