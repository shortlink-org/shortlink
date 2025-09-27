package traicing

import (
	"time"
)

// Config represents the comprehensive configuration for the tracing subsystem.
//
// This configuration struct coordinates between OpenTelemetry tracing and
// Go 1.25 FlightRecorder capabilities. It follows the Single Responsibility
// Principle by encapsulating all tracing-related configuration in one place.
//
// Configuration principles:
//   - Immutable after creation to prevent accidental modification
//   - Optional FlightRecorder to allow graceful degradation
//   - Service identification for proper trace attribution
//   - External service URIs for integration points
//
// Thread safety: This struct should be treated as immutable after creation.
// All fields should be set during initialization and not modified afterward.
type Config struct {
	// ServiceName identifies this service in distributed tracing systems.
	// This value appears in trace metadata and should follow naming conventions.
	ServiceName string
	
	// ServiceVersion provides version information for trace correlation.
	// Useful for identifying behavior changes across deployments.
	ServiceVersion string
	
	// URI specifies the OpenTelemetry trace export endpoint.
	// Typically points to Jaeger, Zipkin, or OTLP-compatible collectors.
	URI string
	
	// PyroscopeURI configures continuous profiling integration.
	// Optional field for enhanced observability and performance analysis.
	PyroscopeURI string
	
	// FlightRecorder contains Go 1.25 FlightRecorder configuration.
	// When nil, FlightRecorder functionality is disabled.
	FlightRecorder *FlightRecorderConfig
}

// FlightRecorderConfig defines the configuration parameters for Go 1.25 trace.FlightRecorder.
// This configuration enables continuous, low-overhead tracing by maintaining a rolling buffer
// of the most recent execution trace data in memory for post-mortem analysis.
type FlightRecorderConfig struct {
	// Enabled determines whether the flight recorder should be activated.
	// When false, no trace data will be collected or buffered.
	Enabled bool
	
	// MinAge specifies the minimum duration of trace data to retain in the buffer.
	// This is a hint to the Go runtime; actual retention may vary based on system conditions.
	// Typical values range from 30 seconds to several minutes depending on application needs.
	MinAge time.Duration
	
	// MaxBytes defines the maximum buffer size in bytes for storing trace data.
	// This is also a hint to the Go runtime; actual memory usage may exceed this value.
	// The buffer operates as a ring buffer, discarding oldest data when the limit is approached.
	// Recommended values are 1-10MB for most applications.
	MaxBytes int64
}
