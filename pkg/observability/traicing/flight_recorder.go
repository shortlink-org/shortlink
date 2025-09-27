// Package traicing provides observability and tracing capabilities using Go 1.25 FlightRecorder.
//
// The FlightRecorder implementation offers continuous, low-overhead tracing by maintaining
// a rolling buffer of execution trace data in memory. This approach enables "perfect tracing"
// where trace data is always available for post-mortem analysis without the performance cost
// of continuous disk writes.
//
// Key features:
//   - Thread-safe operations with comprehensive error handling
//   - Configurable buffer size and retention policies
//   - Automatic trace capture on errors, panics, and signals
//   - Integration with existing observability infrastructure
//   - Clean architecture principles with proper separation of concerns
//
// The implementation follows modern Go practices including:
//   - Dependency injection pattern
//   - Interface-based design for testability
//   - Proper resource lifecycle management
//   - Structured logging with contextual information
package traicing

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime/trace"
	"sync"
	"time"

	"github.com/shortlink-org/go-sdk/logger"
)

// FlightRecorder wraps the Go 1.25 trace.FlightRecorder with additional functionality
type FlightRecorder struct {
	fr      *trace.FlightRecorder
	config  FlightRecorderConfig
	log     logger.Logger
	mu      sync.RWMutex
	running bool
}

// NewFlightRecorder creates a new FlightRecorder instance
func NewFlightRecorder(config FlightRecorderConfig, log logger.Logger) (*FlightRecorder, error) {
	if !config.Enabled {
		return nil, nil
	}

	// Configure the flight recorder to keep
	// at least 5 seconds of trace data,
	// with a maximum buffer size of 3MB.
	// Both of these are hints, not strict limits.
	cfg := trace.FlightRecorderConfig{
		MinAge:   config.MinAge,
		MaxBytes: uint64(config.MaxBytes),
	}

	// Create the flight recorder
	fr := trace.NewFlightRecorder(cfg)

	return &FlightRecorder{
		fr:     fr,
		config: config,
		log:    log,
	}, nil
}

// Start begins recording trace data
func (fr *FlightRecorder) Start() error {
	if fr == nil {
		return nil
	}

	fr.mu.Lock()
	defer fr.mu.Unlock()

	if fr.running {
		return fmt.Errorf("flight recorder is already running")
	}

	if err := fr.fr.Start(); err != nil {
		return fmt.Errorf("failed to start flight recorder: %w", err)
	}

	fr.running = true

	fr.log.Info("Flight recorder started", 
		slog.String("min_age", fr.config.MinAge.String()),
		slog.Int64("max_bytes", fr.config.MaxBytes))

	return nil
}

// Stop stops recording trace data
func (fr *FlightRecorder) Stop() error {
	if fr == nil {
		return nil
	}

	fr.mu.Lock()
	defer fr.mu.Unlock()

	if !fr.running {
		return nil
	}

	fr.fr.Stop()

	fr.running = false

	fr.log.Info("Flight recorder stopped")

	return nil
}

// WriteTo writes the current trace data to the provided writer
func (fr *FlightRecorder) WriteTo(w io.Writer) (int64, error) {
	if fr == nil {
		return 0, fmt.Errorf("flight recorder is not initialized")
	}

	fr.mu.RLock()
	defer fr.mu.RUnlock()

	if !fr.running {
		return 0, fmt.Errorf("flight recorder is not running")
	}

	n, err := fr.fr.WriteTo(w)
	if err != nil {
		fr.log.Error("Failed to write trace data", slog.Any("err", err))
		return n, fmt.Errorf("failed to write trace data: %w", err)
	}

	fr.log.Info("Trace data written", slog.Int64("bytes_written", n))

	return n, nil
}

// WriteToFile writes the current trace data to a file
func (fr *FlightRecorder) WriteToFile(filename string) error {
	if fr == nil {
		return fmt.Errorf("flight recorder is not initialized")
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create trace file %s: %w", filename, err)
	}
	defer file.Close()

	_, err = fr.WriteTo(file)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}

	fr.log.Info("Trace data saved to file", slog.String("filename", filename))

	return nil
}

// IsRunning returns whether the flight recorder is currently running
func (fr *FlightRecorder) IsRunning() bool {
	if fr == nil {
		return false
	}

	fr.mu.RLock()
	defer fr.mu.RUnlock()

	return fr.running
}

// GetConfig returns the flight recorder configuration
func (fr *FlightRecorder) GetConfig() FlightRecorderConfig {
	if fr == nil {
		return FlightRecorderConfig{}
	}

	return fr.config
}

// SaveTraceOnPanic saves trace data when a panic occurs
func (fr *FlightRecorder) SaveTraceOnPanic() {
	if fr == nil || !fr.IsRunning() {
		return
	}

	if r := recover(); r != nil {
		filename := fmt.Sprintf("panic_trace_%d.out", time.Now().Unix())
		if err := fr.WriteToFile(filename); err != nil {
			fr.log.Error("Failed to save panic trace", 
				slog.Any("err", err),
				slog.String("filename", filename),
				slog.Any("panic", r))
		} else {
			fr.log.Info("Panic trace saved", 
				slog.String("filename", filename),
				slog.Any("panic", r))
		}

		// Re-panic to preserve the original panic behavior
		panic(r)
	}
}

// DefaultFlightRecorderConfig returns a default configuration for the flight recorder
func DefaultFlightRecorderConfig() FlightRecorderConfig {
	return FlightRecorderConfig{
		Enabled:  true,
		MinAge:   1 * time.Minute,
		MaxBytes: 3 << 20, // 3MB
	}
}