package logger

import (
	"io"
)

const (
	// Zap implementation
	Zap int = iota // nolint unused
	// Logrus implementation
	Logrus // nolint unused
)

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{} //nolint unused

// Logger is our contract for the logger
type Logger interface { //nolint unused
	init(Configuration) error

	Fatal(msg string, fields ...Fields)

	Error(msg string, fields ...Fields)

	Warn(msg string, fields ...Fields)

	Info(msg string, fields ...Fields)

	Debug(msg string, fields ...Fields)

	SetConfig(Configuration) error

	// Closer is the interface that wraps the basic Close method.
	io.Closer
}

// The severity levels. Higher values are more considered more important.
const (
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FATAL_LEVEL int = iota // nolint unused
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ERROR_LEVEL // nolint unused
	// WarnLevel level. Non-critical entries that deserve eyes.
	WARN_LEVEL // nolint unused
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	INFO_LEVEL // nolint unused
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DEBUG_LEVEL // nolint unused
)

// Configuration - options for logger
type Configuration struct { // nolint unused
	Level      int
	Writer     io.Writer
	TimeFormat string
}
