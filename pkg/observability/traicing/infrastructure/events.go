// Package infrastructure provides event handling implementations.
// This layer manages cross-cutting concerns like logging, monitoring, and notifications.
package infrastructure

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/traicing/domain"
)

// LoggingEventHandler implements domain.EventHandler using structured logging.
// This implementation provides comprehensive audit trails and monitoring integration.
type LoggingEventHandler struct {
	logger logger.Logger
}

// NewLoggingEventHandler creates a new logging-based event handler.
// It requires a structured logger that supports contextual logging.
func NewLoggingEventHandler(log logger.Logger) *LoggingEventHandler {
	return &LoggingEventHandler{
		logger: log,
	}
}

// OnStarted logs when the recorder successfully starts.
// This event includes configuration details for audit purposes.
func (h *LoggingEventHandler) OnStarted(ctx context.Context, config *domain.Configuration) {
	h.logger.InfoWithContext(ctx, "FlightRecorder started successfully",
		slog.Bool("enabled", config.Enabled()),
		slog.String("min_age", config.MinAge().String()),
		slog.Uint64("max_bytes", config.MaxBytes()),
	)
}

// OnStopped logs when the recorder stops.
// This provides visibility into recorder lifecycle events.
func (h *LoggingEventHandler) OnStopped(ctx context.Context) {
	h.logger.InfoWithContext(ctx, "FlightRecorder stopped")
}

// OnError logs recorder errors with appropriate severity.
// This enables proper error monitoring and alerting.
func (h *LoggingEventHandler) OnError(ctx context.Context, err error) {
	h.logger.ErrorWithContext(ctx, "FlightRecorder error occurred",
		slog.String("error", err.Error()),
		slog.String("error_type", fmt.Sprintf("%T", err)),
	)
}

// OnTraceSaved logs successful trace save operations.
// This provides audit trails and performance metrics.
func (h *LoggingEventHandler) OnTraceSaved(ctx context.Context, id string, bytes int64) {
	h.logger.InfoWithContext(ctx, "Trace data saved successfully",
		slog.String("trace_id", id),
		slog.Int64("bytes_written", bytes),
	)
}

// CompositeEventHandler combines multiple event handlers.
// This enables sending events to multiple destinations (logs, metrics, notifications).
type CompositeEventHandler struct {
	handlers []domain.EventHandler
}

// NewCompositeEventHandler creates a new composite event handler.
// It allows combining multiple event handling strategies.
func NewCompositeEventHandler(handlers ...domain.EventHandler) *CompositeEventHandler {
	return &CompositeEventHandler{
		handlers: handlers,
	}
}

// OnStarted notifies all registered handlers.
func (h *CompositeEventHandler) OnStarted(ctx context.Context, config *domain.Configuration) {
	for _, handler := range h.handlers {
		handler.OnStarted(ctx, config)
	}
}

// OnStopped notifies all registered handlers.
func (h *CompositeEventHandler) OnStopped(ctx context.Context) {
	for _, handler := range h.handlers {
		handler.OnStopped(ctx)
	}
}

// OnError notifies all registered handlers.
func (h *CompositeEventHandler) OnError(ctx context.Context, err error) {
	for _, handler := range h.handlers {
		handler.OnError(ctx, err)
	}
}

// OnTraceSaved notifies all registered handlers.
func (h *CompositeEventHandler) OnTraceSaved(ctx context.Context, id string, bytes int64) {
	for _, handler := range h.handlers {
		handler.OnTraceSaved(ctx, id, bytes)
	}
}

// MetricsEventHandler implements domain.EventHandler for metrics collection.
// This handler can be used to integrate with monitoring systems like Prometheus.
type MetricsEventHandler struct {
	// Metrics collectors would be injected here
	// For example: prometheus.CounterVec, prometheus.HistogramVec, etc.
}

// NewMetricsEventHandler creates a new metrics-based event handler.
func NewMetricsEventHandler() *MetricsEventHandler {
	return &MetricsEventHandler{}
}

// OnStarted records start metrics.
func (h *MetricsEventHandler) OnStarted(ctx context.Context, config *domain.Configuration) {
	// Increment start counter
	// Record configuration metrics
}

// OnStopped records stop metrics.
func (h *MetricsEventHandler) OnStopped(ctx context.Context) {
	// Increment stop counter
}

// OnError records error metrics.
func (h *MetricsEventHandler) OnError(ctx context.Context, err error) {
	// Increment error counter by error type
}

// OnTraceSaved records trace save metrics.
func (h *MetricsEventHandler) OnTraceSaved(ctx context.Context, id string, bytes int64) {
	// Record bytes written histogram
	// Increment save counter
}

// Ensure implementations satisfy the interface
var (
	_ domain.EventHandler = (*LoggingEventHandler)(nil)
	_ domain.EventHandler = (*CompositeEventHandler)(nil)
	_ domain.EventHandler = (*MetricsEventHandler)(nil)
)