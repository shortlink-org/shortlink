// Package infrastructure provides concrete implementations of domain interfaces.
// This layer handles external dependencies and technical concerns.
package infrastructure

import (
	"context"
	"fmt"
	"io"
	"runtime/trace"
	"sync"

	"github.com/shortlink-org/shortlink/pkg/observability/traicing/domain"
)

// GoFlightRecorder implements domain.Recorder using Go 1.25's trace.FlightRecorder.
// This adapter bridges the domain interface with the Go runtime implementation.
type GoFlightRecorder struct {
	recorder *trace.FlightRecorder
	config   *domain.Configuration
	state    domain.RecorderState
	mu       sync.RWMutex
}

// NewGoFlightRecorder creates a new GoFlightRecorder instance with the provided configuration.
// It initializes the underlying Go trace.FlightRecorder with validated parameters.
func NewGoFlightRecorder(config *domain.Configuration) (*GoFlightRecorder, error) {
	if config == nil {
		return nil, domain.ErrInvalidConfiguration
	}

	if !config.Enabled() {
		return &GoFlightRecorder{
			config: config,
			state:  domain.StateStopped,
		}, nil
	}

	// Configure the Go 1.25 FlightRecorder with domain configuration
	runtimeConfig := trace.FlightRecorderConfig{
		MinAge:   config.MinAge(),
		MaxBytes: config.MaxBytes(),
	}

	recorder := trace.NewFlightRecorder(runtimeConfig)

	return &GoFlightRecorder{
		recorder: recorder,
		config:   config,
		state:    domain.StateStopped,
	}, nil
}

// Start initiates trace data collection.
// This method is thread-safe and idempotent.
func (r *GoFlightRecorder) Start(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.config.Enabled() {
		return domain.ErrRecorderDisabled
	}

	if r.state == domain.StateRunning {
		return domain.ErrAlreadyRunning
	}

	if r.recorder == nil {
		r.state = domain.StateError
		return fmt.Errorf("recorder not initialized")
	}

	if err := r.recorder.Start(); err != nil {
		r.state = domain.StateError
		return fmt.Errorf("failed to start Go flight recorder: %w", err)
	}

	r.state = domain.StateRunning
	return nil
}

// Stop terminates trace data collection.
// This method is thread-safe and idempotent.
func (r *GoFlightRecorder) Stop(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.state != domain.StateRunning {
		return nil // Idempotent operation
	}

	if r.recorder == nil {
		r.state = domain.StateStopped
		return nil
	}

	// Go 1.25 FlightRecorder.Stop() doesn't return an error
	r.recorder.Stop()
	r.state = domain.StateStopped
	return nil
}

// State returns the current operational state.
// This method is thread-safe.
func (r *GoFlightRecorder) State() domain.RecorderState {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.state
}

// WriteTo exports the current trace buffer to the provided writer.
// This method is thread-safe and validates the recorder state.
func (r *GoFlightRecorder) WriteTo(w io.Writer) (int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if !r.config.Enabled() {
		return 0, domain.ErrRecorderDisabled
	}

	if r.state != domain.StateRunning {
		return 0, domain.ErrNotRunning
	}

	if r.recorder == nil {
		return 0, fmt.Errorf("recorder not initialized")
	}

	bytesWritten, err := r.recorder.WriteTo(w)
	if err != nil {
		r.state = domain.StateError
		return bytesWritten, fmt.Errorf("failed to write trace data: %w", err)
	}

	return bytesWritten, nil
}

// Configuration returns the recorder configuration.
// This method is thread-safe.
func (r *GoFlightRecorder) Configuration() *domain.Configuration {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.config
}

// Ensure GoFlightRecorder implements domain.Recorder
var _ domain.Recorder = (*GoFlightRecorder)(nil)