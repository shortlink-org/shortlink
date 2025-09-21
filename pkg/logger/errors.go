package logger

import "errors"

var (
	// ErrInvalidLogLevel is an error when log level is invalid.
	ErrInvalidLogLevel = errors.New("invalid log level")
)