// Package domain defines the core business entities and interfaces for the FlightRecorder domain.
// This layer contains the essential business logic and is independent of external concerns.
package domain

import (
	"context"
	"io"
	"time"
)

// RecorderState represents the operational state of a FlightRecorder.
type RecorderState uint8

const (
	// StateStopped indicates the recorder is not actively collecting traces.
	StateStopped RecorderState = iota
	// StateRunning indicates the recorder is actively collecting trace data.
	StateRunning
	// StateError indicates the recorder encountered an error and requires intervention.
	StateError
)

// String provides a human-readable representation of the recorder state.
func (s RecorderState) String() string {
	switch s {
	case StateStopped:
		return "stopped"
	case StateRunning:
		return "running"
	case StateError:
		return "error"
	default:
		return "unknown"
	}
}

// Configuration encapsulates the FlightRecorder configuration parameters.
// This value object ensures immutability and validation of configuration data.
type Configuration struct {
	enabled  bool
	minAge   time.Duration
	maxBytes uint64
}

// NewConfiguration creates a validated Configuration instance.
// It ensures all parameters meet the minimum requirements for safe operation.
func NewConfiguration(enabled bool, minAge time.Duration, maxBytes uint64) (*Configuration, error) {
	if enabled {
		if minAge < time.Second {
			return nil, ErrInvalidMinAge
		}
		if maxBytes < 1024*1024 { // Minimum 1MB
			return nil, ErrInvalidMaxBytes
		}
	}

	return &Configuration{
		enabled:  enabled,
		minAge:   minAge,
		maxBytes: maxBytes,
	}, nil
}

// Enabled returns whether the recorder should be active.
func (c *Configuration) Enabled() bool {
	return c.enabled
}

// MinAge returns the minimum trace data retention duration.
func (c *Configuration) MinAge() time.Duration {
	return c.minAge
}

// MaxBytes returns the maximum buffer size in bytes.
func (c *Configuration) MaxBytes() uint64 {
	return c.maxBytes
}

// Recorder defines the core FlightRecorder behavior.
// This interface abstracts the recording capabilities from implementation details.
type Recorder interface {
	// Start initiates trace data collection.
	// Returns ErrAlreadyRunning if the recorder is already active.
	Start(ctx context.Context) error

	// Stop terminates trace data collection gracefully.
	// This operation is idempotent and safe to call multiple times.
	Stop(ctx context.Context) error

	// State returns the current operational state of the recorder.
	State() RecorderState

	// WriteTo exports the current trace buffer to the provided writer.
	// Returns the number of bytes written and any error encountered.
	WriteTo(w io.Writer) (int64, error)

	// Configuration returns the current recorder configuration.
	Configuration() *Configuration
}

// Repository defines persistence operations for trace data.
// This interface allows for different storage implementations (file, S3, etc.).
type Repository interface {
	// Save persists trace data with the given identifier.
	Save(ctx context.Context, id string, data io.Reader) error

	// Load retrieves trace data by identifier.
	Load(ctx context.Context, id string) (io.Reader, error)

	// Delete removes trace data by identifier.
	Delete(ctx context.Context, id string) error

	// List returns available trace identifiers.
	List(ctx context.Context) ([]string, error)
}

// EventHandler defines the contract for handling recorder events.
// This enables the implementation of cross-cutting concerns like logging and monitoring.
type EventHandler interface {
	// OnStarted is called when the recorder successfully starts.
	OnStarted(ctx context.Context, config *Configuration)

	// OnStopped is called when the recorder stops.
	OnStopped(ctx context.Context)

	// OnError is called when the recorder encounters an error.
	OnError(ctx context.Context, err error)

	// OnTraceSaved is called when trace data is successfully saved.
	OnTraceSaved(ctx context.Context, id string, bytes int64)
}