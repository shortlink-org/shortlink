// Package traicing provides enterprise-grade observability and tracing capabilities.
//
// This package implements a comprehensive tracing solution using Go 1.25's FlightRecorder
// with clean architecture principles. The design separates concerns into distinct layers:
//
// Domain Layer (pkg/observability/traicing/domain/):
//   - Core business entities and interfaces
//   - Domain-specific error definitions  
//   - Value objects with validation
//
// Application Layer (pkg/observability/traicing/application/):
//   - Use case orchestration and business workflows
//   - Cross-cutting concern coordination
//   - Service composition and transaction management
//
// Infrastructure Layer (pkg/observability/traicing/infrastructure/):
//   - External system integrations (Go runtime, filesystem)
//   - Technical implementation details
//   - Framework and library adapters
//
// Factory Pattern (pkg/observability/traicing/factory.go):
//   - Dependency injection and component wiring
//   - Configuration management and validation
//   - Clean separation of object creation concerns
//
// This legacy file is maintained for backward compatibility but the new
// clean architecture implementation should be used for new development.
// 
// Deprecated: Use the new factory-based approach with proper layered architecture.
// Example: factory.CreateRecorderService() instead of NewFlightRecorder()
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

// FlightRecorder provides a legacy wrapper around Go 1.25's trace.FlightRecorder.
//
// This implementation maintains backward compatibility while encouraging migration
// to the new clean architecture approach. It encapsulates the Go runtime's
// FlightRecorder with additional safety measures and operational visibility.
//
// Key characteristics:
//   - Thread-safe operations using read-write mutex
//   - Comprehensive error handling and validation
//   - Structured logging for operational visibility
//   - Configuration validation and immutability
//   - Graceful degradation when disabled
//
// Deprecated: This struct is maintained for backward compatibility.
// New code should use the factory pattern with proper dependency injection:
//   factory := traicing.NewFactory(config)
//   service := factory.CreateRecorderService()
type FlightRecorder struct {
	// fr holds the underlying Go 1.25 FlightRecorder instance
	fr *trace.FlightRecorder
	
	// config stores the immutable configuration for this recorder
	config FlightRecorderConfig
	
	// log provides structured logging for operational visibility
	log logger.Logger
	
	// mu protects concurrent access to mutable state
	mu sync.RWMutex
	
	// running tracks the current operational state
	running bool
}

// NewFlightRecorder creates a new FlightRecorder instance with comprehensive validation.
//
// This constructor performs extensive validation and initialization to ensure
// the recorder is properly configured and ready for operation. It follows the
// fail-fast principle by validating configuration early.
//
// Parameters:
//   - config: Immutable configuration defining recorder behavior
//   - log: Structured logger for operational visibility and debugging
//
// Returns:
//   - *FlightRecorder: Configured recorder instance ready for use
//   - error: Validation or initialization errors
//
// Error conditions:
//   - Invalid configuration parameters
//   - Logger instance is nil
//   - Go runtime FlightRecorder creation failure
//
// Deprecated: Use factory.CreateRecorderService() for new implementations.
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

// Start initiates trace data collection with comprehensive error handling.
//
// This method performs pre-flight validation, state checking, and safe
// initialization of the underlying Go runtime FlightRecorder. It implements
// the Command pattern for state-changing operations.
//
// Operational characteristics:
//   - Thread-safe using write mutex
//   - Idempotent operation (safe to call multiple times)
//   - Comprehensive logging for operational visibility
//   - Graceful handling of disabled configurations
//
// Returns:
//   - nil: Successful start or already running
//   - error: Configuration, state, or runtime errors
//
// Error conditions:
//   - Recorder is disabled in configuration
//   - Already running (idempotent, returns error for clarity)
//   - Go runtime FlightRecorder initialization failure
//
// Deprecated: Use service.StartRecording(ctx) for new implementations.
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