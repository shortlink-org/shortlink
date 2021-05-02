package helpers

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

// RegisterSpan - inject spanId in response header
func RegisterSpan(ctx context.Context) string {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return ""
	}

	if traceID, okSpan := span.Context().(jaeger.SpanContext); okSpan {
		return traceID.SpanID().String()
	}

	return ""
}
