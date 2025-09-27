// Package domain defines domain-specific errors for the FlightRecorder.
// These errors represent business rule violations and expected failure scenarios.
package domain

import "errors"

var (
	// ErrRecorderDisabled indicates an operation was attempted on a disabled recorder.
	ErrRecorderDisabled = errors.New("flight recorder is disabled")

	// ErrAlreadyRunning indicates an attempt to start an already running recorder.
	ErrAlreadyRunning = errors.New("flight recorder is already running")

	// ErrNotRunning indicates an operation that requires a running recorder.
	ErrNotRunning = errors.New("flight recorder is not running")

	// ErrInvalidConfiguration indicates the provided configuration is invalid.
	ErrInvalidConfiguration = errors.New("invalid flight recorder configuration")

	// ErrInvalidMinAge indicates the minimum age parameter is invalid.
	ErrInvalidMinAge = errors.New("minimum age must be at least 1 second")

	// ErrInvalidMaxBytes indicates the maximum bytes parameter is invalid.
	ErrInvalidMaxBytes = errors.New("maximum bytes must be at least 1MB")

	// ErrTraceNotFound indicates the requested trace data was not found.
	ErrTraceNotFound = errors.New("trace data not found")

	// ErrRepositoryUnavailable indicates the storage repository is unavailable.
	ErrRepositoryUnavailable = errors.New("trace repository is unavailable")
)

// ValidationError wraps validation errors with additional context.
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
	Cause   error
}

// Error returns the error message.
func (e *ValidationError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

// Unwrap returns the underlying error.
func (e *ValidationError) Unwrap() error {
	return e.Cause
}

// NewValidationError creates a new validation error with context.
func NewValidationError(field string, value interface{}, message string, cause error) *ValidationError {
	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
		Cause:   cause,
	}
}