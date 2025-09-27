// Package application implements the FlightRecorder application services.
// This layer orchestrates domain operations and coordinates between different bounded contexts.
package application

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/shortlink-org/shortlink/pkg/observability/traicing/domain"
)

// RecorderService provides high-level operations for FlightRecorder management.
// It orchestrates domain entities and coordinates cross-cutting concerns.
type RecorderService struct {
	recorder      domain.Recorder
	repository    domain.Repository
	eventHandler  domain.EventHandler
	config        *domain.Configuration
}

// NewRecorderService creates a new RecorderService with the provided dependencies.
// This constructor ensures all required dependencies are properly injected.
func NewRecorderService(
	recorder domain.Recorder,
	repository domain.Repository,
	eventHandler domain.EventHandler,
	config *domain.Configuration,
) *RecorderService {
	return &RecorderService{
		recorder:     recorder,
		repository:   repository,
		eventHandler: eventHandler,
		config:       config,
	}
}

// StartRecording initiates trace data collection with proper error handling and event notification.
// This operation is idempotent and includes comprehensive validation.
func (s *RecorderService) StartRecording(ctx context.Context) error {
	if !s.config.Enabled() {
		return domain.ErrRecorderDisabled
	}

	if s.recorder.State() == domain.StateRunning {
		return domain.ErrAlreadyRunning
	}

	if err := s.recorder.Start(ctx); err != nil {
		s.eventHandler.OnError(ctx, err)
		return fmt.Errorf("failed to start flight recorder: %w", err)
	}

	s.eventHandler.OnStarted(ctx, s.config)
	return nil
}

// StopRecording terminates trace data collection gracefully.
// This operation ensures proper cleanup and event notification.
func (s *RecorderService) StopRecording(ctx context.Context) error {
	if s.recorder.State() != domain.StateRunning {
		return nil // Idempotent operation
	}

	if err := s.recorder.Stop(ctx); err != nil {
		s.eventHandler.OnError(ctx, err)
		return fmt.Errorf("failed to stop flight recorder: %w", err)
	}

	s.eventHandler.OnStopped(ctx)
	return nil
}

// CaptureTrace saves the current trace buffer to persistent storage.
// It generates a unique identifier and handles the complete save operation.
func (s *RecorderService) CaptureTrace(ctx context.Context, reason string) (string, error) {
	if s.recorder.State() != domain.StateRunning {
		return "", domain.ErrNotRunning
	}

	// Generate unique trace identifier
	traceID := s.generateTraceID(reason)

	// Create a pipe for streaming trace data
	pr, pw := io.Pipe()
	defer pr.Close()

	// Start the save operation in a goroutine
	errCh := make(chan error, 1)
	go func() {
		defer pw.Close()
		errCh <- s.repository.Save(ctx, traceID, pr)
	}()

	// Write trace data to the pipe
	bytesWritten, err := s.recorder.WriteTo(pw)
	if err != nil {
		return "", fmt.Errorf("failed to write trace data: %w", err)
	}

	// Wait for save operation to complete
	if err := <-errCh; err != nil {
		return "", fmt.Errorf("failed to save trace data: %w", err)
	}

	s.eventHandler.OnTraceSaved(ctx, traceID, bytesWritten)
	return traceID, nil
}

// GetRecorderStatus returns comprehensive status information about the recorder.
// This includes operational state, configuration, and performance metrics.
func (s *RecorderService) GetRecorderStatus() *RecorderStatus {
	return &RecorderStatus{
		State:         s.recorder.State(),
		Configuration: s.config,
		Enabled:       s.config.Enabled(),
	}
}

// ListTraces returns available trace identifiers from the repository.
func (s *RecorderService) ListTraces(ctx context.Context) ([]string, error) {
	traces, err := s.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list traces: %w", err)
	}
	return traces, nil
}

// LoadTrace retrieves trace data by identifier.
func (s *RecorderService) LoadTrace(ctx context.Context, traceID string) (io.Reader, error) {
	data, err := s.repository.Load(ctx, traceID)
	if err != nil {
		return nil, fmt.Errorf("failed to load trace %s: %w", traceID, err)
	}
	return data, nil
}

// DeleteTrace removes trace data by identifier.
func (s *RecorderService) DeleteTrace(ctx context.Context, traceID string) error {
	if err := s.repository.Delete(ctx, traceID); err != nil {
		return fmt.Errorf("failed to delete trace %s: %w", traceID, err)
	}
	return nil
}

// generateTraceID creates a unique identifier for trace data.
// The format includes timestamp and reason for better traceability.
func (s *RecorderService) generateTraceID(reason string) string {
	timestamp := time.Now().UTC().Format("20060102_150405")
	return fmt.Sprintf("%s_%s_%d", reason, timestamp, time.Now().UnixNano()%1000000)
}

// RecorderStatus provides comprehensive information about the recorder state.
type RecorderStatus struct {
	State         domain.RecorderState
	Configuration *domain.Configuration
	Enabled       bool
}

// IsHealthy returns whether the recorder is in a healthy operational state.
func (s *RecorderStatus) IsHealthy() bool {
	if !s.Enabled {
		return true // Disabled recorder is considered healthy
	}
	return s.State == domain.StateRunning || s.State == domain.StateStopped
}