package logger

import (
	"context"
	"io"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

const (
	// Zap implementation
	Zap int = iota
	// Logrus implementation
	Logrus
)

// Logger is our contract for the logger
type Logger interface { //nolint:decorder
	Fatal(msg string, fields ...field.Fields)
	FatalWithContext(ctx context.Context, msg string, fields ...field.Fields)

	Error(msg string, fields ...field.Fields)
	ErrorWithContext(ctx context.Context, msg string, fields ...field.Fields)

	Warn(msg string, fields ...field.Fields)
	WarnWithContext(ctx context.Context, msg string, fields ...field.Fields)

	Info(msg string, fields ...field.Fields)
	InfoWithContext(ctx context.Context, msg string, fields ...field.Fields)

	Debug(msg string, fields ...field.Fields)
	DebugWithContext(ctx context.Context, msg string, fields ...field.Fields)

	Get() any

	// Closer is the interface that wraps the basic Close method.
	io.Closer
}
