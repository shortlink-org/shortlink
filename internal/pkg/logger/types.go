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
type Logger interface { // nolint:decorder
	init(Configuration) error

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

	Get() interface{}

	// Closer is the interface that wraps the basic Close method.
	io.Closer
}

// The severity levels. Higher values are more considered more important.
const (
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FATAL_LEVEL int = iota
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ERROR_LEVEL
	// WarnLevel level. Non-critical entries that deserve eyes.
	WARN_LEVEL
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	INFO_LEVEL
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DEBUG_LEVEL
)

// Configuration - options for logger
type Configuration struct {
	Writer     io.Writer
	TimeFormat string
	Level      int
}
