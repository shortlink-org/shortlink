package tracer

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const callersSkip = 3

// ErrLogField is a sentinel used to wrap error messages extracted from log fields.
var ErrLogField = errors.New("log field error")

// NewTraceFromContext
// - If an active span exists: add an Event ("log.<LEVEL>") with attributes.
// - If there is no active span: create a short span only for WARN/ERROR.
// - Always return fields augmented with traceID/spanID when a span exists.
func NewTraceFromContext(
	ctx context.Context,
	level string, // "INFO"|"WARN"|"ERROR"|...
	msg string,
	tags []attribute.KeyValue, // extra attributes
	fields ...any, // pairs key,value for logs
) ([]any, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	levelUpper := strings.ToUpper(level)

	// 1) Normalize fields → attributes with a consistent schema
	attrs, capturedErr := normalizeToAttrs(msg, levelUpper, tags, fields...)

	// 2) If an active span exists — write an Event
	if span := trace.SpanFromContext(ctx); span != nil && span.SpanContext().IsValid() {
		span.AddEvent("log."+levelUpper, trace.WithAttributes(attrs...))
		annotateByLevel(span, levelUpper, msg, capturedErr)

		out := append(append([]any{}, fields...),
			"traceID", span.SpanContext().TraceID().String(),
			"spanID", span.SpanContext().SpanID().String(),
		)

		return out, nil
	}

	// 3) No active span — create only for important levels
	if isSmallLevel(levelUpper) {
		// Small levels (INFO/DEBUG/TRACE): do not create a span; just return fields
		return fields, nil
	}

	// Short correlation span for WARN/ERROR
	_, span := otel.Tracer("logger").Start(ctx, getNameFunc())
	defer span.End()

	span.SetAttributes(attrs...)
	annotateByLevel(span, levelUpper, msg, capturedErr)

	out := append(append([]any{}, fields...),
		"traceID", span.SpanContext().TraceID().String(),
		"spanID", span.SpanContext().SpanID().String(),
	)

	return out, nil
}

func isSmallLevel(levelUpper string) bool {
	switch levelUpper {
	case "DEBUG", "INFO", "TRACE":
		return true
	default:
		return false
	}
}

// getNameFunc returns the caller function name to use as a span name.
func getNameFunc() string {
	pc := make([]uintptr, 1)
	if n := runtime.Callers(callersSkip, pc); n > 0 {
		if f := runtime.FuncForPC(pc[0]); f != nil {
			return f.Name()
		}
	}

	return "log"
}

// annotateByLevel sets ERROR status and exception.* if error/level indicates failure.
func annotateByLevel(span trace.Span, level, msg string, err error) {
	if err != nil {
		span.SetAttributes(
			attribute.String("exception.type", fmt.Sprintf("%T", err)),
			attribute.String("exception.message", err.Error()),
		)
	}

	switch level {
	case "ERROR", "FATAL":
		span.SetStatus(codes.Error, msg)
	}
}

// normalizeToAttrs:
//   - adds log.severity, log.message
//   - normalizes err/error → exception.message + exception.type
//   - maps is_error → log.is_error
//   - converts the rest of key/value pairs to attributes
func normalizeToAttrs(msg, level string, tags []attribute.KeyValue, fields ...any) ([]attribute.KeyValue, error) {
	attrs := make([]attribute.KeyValue, 0, len(fields)/2+4+len(tags))
	attrs = append(attrs,
		attribute.String("log.severity", level),
		attribute.String("log.message", msg),
	)
	attrs = append(attrs, tags...)

	var capturedErr error

	for i := 0; i+1 < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok || key == "" {
			continue
		}

		val := fields[i+1]

		switch strings.ToLower(key) {
		case "err", "error":
			switch typedVal := val.(type) {
			case error:
				capturedErr = typedVal
				attrs = append(attrs, attribute.String("exception.message", typedVal.Error()))
				attrs = append(attrs, attribute.String("exception.type", fmt.Sprintf("%T", typedVal)))
			case string:
				if typedVal != "" {
					capturedErr = fmt.Errorf("%w: %s", ErrLogField, typedVal) // wrap sentinel (err113)
					attrs = append(attrs, attribute.String("exception.message", typedVal))
				}
			default:
				asStr := fmt.Sprintf("%v", typedVal)
				if asStr != "" {
					capturedErr = fmt.Errorf("%w: %s", ErrLogField, asStr) // wrap sentinel (err113)
					attrs = append(attrs, attribute.String("exception.message", asStr))
				}
			}
		case "is_error", "iserror", "error_flag":
			attrs = append(attrs, attribute.Bool("log.is_error", toBool(val)))
		default:
			attrs = append(attrs, kv(key, val))
		}
	}

	return attrs, capturedErr
}

// kv converts most common Go types to OTel attributes.
func kv(key string, v any) attribute.KeyValue {
	switch typed := v.(type) {
	case string:
		return attribute.String(key, typed)
	case bool:
		return attribute.Bool(key, typed)
	case int:
		return attribute.Int(key, typed)
	case int32:
		return attribute.Int(key, int(typed))
	case int64:
		return attribute.Int64(key, typed)
	case float32:
		return attribute.Float64(key, float64(typed))
	case float64:
		return attribute.Float64(key, typed)
	case error:
		return attribute.String(key, typed.Error())
	default:
		return attribute.String(key, fmt.Sprintf("%v", typed))
	}
}

func toBool(v any) bool {
	switch raw := v.(type) {
	case bool:
		return raw
	case string:
		return strings.EqualFold(raw, "true") || raw == "1" || strings.EqualFold(raw, "yes")
	case int:
		return raw == 1
	case int32:
		return raw == 1
	case int64:
		return raw == 1
	default:
		return false
	}
}