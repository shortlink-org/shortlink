package config

import (
	"io"
	"os"
	"time"
)

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

func (c *Configuration) Validate() error {
	if c.Writer == nil {
		c.Writer = os.Stdout
	}

	if c.TimeFormat == "" {
		c.TimeFormat = time.RFC3339Nano
	}

	return nil
}
