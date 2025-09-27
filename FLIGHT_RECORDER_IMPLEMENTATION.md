# Go 1.25 FlightRecorder Implementation

This document summarizes the implementation of Go 1.25 `trace.FlightRecorder` for perfect tracing in the shortlink project.

## Overview

The FlightRecorder provides a lightweight mechanism for continuous tracing by maintaining a moving window of the most recent execution trace data in memory. This is particularly useful for capturing trace information leading up to significant events without the overhead of writing extensive trace data to disk continuously.

## Implementation

### Core Components

1. **FlightRecorderConfig** (`pkg/observability/traicing/types.go`)
   - `Enabled`: Whether the flight recorder is active
   - `MinAge`: Minimum duration of trace data to retain (5 seconds default)
   - `MaxBytes`: Maximum buffer size (3MB default)

2. **FlightRecorder** (`pkg/observability/traicing/flight_recorder.go`)
   - Wraps Go 1.25 `trace.FlightRecorder` with additional functionality
   - Thread-safe operations with mutex protection
   - Automatic logging and error handling

3. **Utility Functions** (`pkg/observability/traicing/utils.go`)
   - Global flight recorder management
   - Error-triggered trace saving
   - Signal-based trace capture
   - Middleware for automatic trace capture

### Key Features

#### 1. Configuration
```go
// Configure the flight recorder to keep
// at least 5 seconds of trace data,
// with a maximum buffer size of 3MB.
// Both of these are hints, not strict limits.
cfg := trace.FlightRecorderConfig{
    MinAge:   5 * time.Second,
    MaxBytes: 3 << 20, // 3MB
}
```

#### 2. Environment Variables
```bash
FLIGHT_RECORDER_ENABLED=true               # Enable/disable flight recorder
FLIGHT_RECORDER_MIN_AGE=5s                 # Minimum age of trace data to retain
FLIGHT_RECORDER_MAX_BYTES=3145728          # Maximum buffer size (3MB)
```

#### 3. Automatic Integration
The FlightRecorder integrates automatically with the existing tracing system:
```go
config := traicing.Config{
    ServiceName:    "my-service",
    ServiceVersion: "1.0.0",
    URI:            "localhost:4317",
    FlightRecorder: &traicing.FlightRecorderConfig{
        Enabled:  true,
        MinAge:   5 * time.Second,
        MaxBytes: 3 << 20, // 3MB
    },
}
```

#### 4. Manual Trace Capture
```go
// Save trace on error
SaveTraceOnError(err, log)

// Save trace with context
SaveTraceWithContext(ctx, "manual_trace", map[string]interface{}{
    "user_id":    "12345",
    "request_id": "req-789",
    "endpoint":   "/api/v1/links",
}, log)

// Save trace on signal (SIGUSR1, SIGUSR2)
SaveTraceOnSignal("USR1", log)
```

#### 5. Middleware Support
```go
middleware := NewRecorderMiddleware(log)

// Automatic error tracking
wrappedFunc := middleware.WrapWithErrorTracking(someFunction)

// Automatic panic recovery with trace saving
safeFunc := middleware.WrapWithPanicRecovery(riskyFunction)
```

## Usage Examples

### Basic Usage
```go
// Create flight recorder
cfg := traicing.DefaultFlightRecorderConfig()
fr, err := traicing.NewFlightRecorder(cfg, log)
if err != nil {
    return err
}

// Start recording
if err := fr.Start(); err != nil {
    return err
}
defer fr.Stop()

// Save trace to file when needed
if err := fr.WriteToFile("trace.out"); err != nil {
    log.Error("Failed to save trace", slog.Any("err", err))
}
```

### Integration with Tracing System
```go
config := traicing.Config{
    ServiceName:    viper.GetString("SERVICE_NAME"),
    ServiceVersion: viper.GetString("SERVICE_VERSION"),
    URI:            viper.GetString("TRACER_URI"),
    FlightRecorder: &traicing.FlightRecorderConfig{
        Enabled:  true,
        MinAge:   5 * time.Second,
        MaxBytes: 3 << 20,
    },
}

tp, cleanup, err := traicing.Init(ctx, config, log)
```

### Health Check
```go
status := traicing.HealthCheck()
// Returns status information about the flight recorder
```

## Files Modified/Added

### New Files
- `pkg/observability/traicing/flight_recorder.go` - Core FlightRecorder implementation
- `pkg/observability/traicing/utils.go` - Utility functions and middleware
- `pkg/observability/traicing/example_usage.go` - Usage examples and documentation

### Modified Files
- `pkg/observability/traicing/types.go` - Added FlightRecorderConfig
- `pkg/observability/traicing/traicing.go` - Integrated FlightRecorder with tracing system
- `pkg/di/pkg/traicing/traicing.go` - Added FlightRecorder configuration to DI

## Benefits

1. **Perfect Tracing**: Captures trace data leading up to significant events
2. **Low Overhead**: In-memory ring buffer with configurable limits
3. **Automatic Integration**: Works seamlessly with existing tracing infrastructure
4. **Flexible Triggering**: Multiple ways to capture traces (errors, signals, manual)
5. **Production Ready**: Thread-safe, well-logged, and error-handled

## Performance Considerations

- **Memory Usage**: Configurable via `MaxBytes` (default 3MB)
- **CPU Overhead**: Minimal when enabled, zero when disabled
- **Storage**: Only writes to disk on demand
- **Thread Safety**: All operations are mutex-protected

## Dependencies

- Go 1.25.1 (for `trace.FlightRecorder`)
- `github.com/shortlink-org/go-sdk/logger` for logging
- Standard library packages: `runtime/trace`, `context`, `sync`, etc.

## Configuration

The FlightRecorder can be configured via:
1. Environment variables (see above)
2. Configuration structs in code
3. Default values (5s minimum age, 3MB max buffer)

All configuration values are treated as hints rather than strict limits, following Go's FlightRecorder design philosophy.