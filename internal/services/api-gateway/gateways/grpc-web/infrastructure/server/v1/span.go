package v1

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RegisterSpan - inject spanId in response header
func RegisterSpan(ctx context.Context) error {
	return grpc.SendHeader(ctx, metadata.New(map[string]string{
		"trace_id": trace.LinkFromContext(ctx).SpanContext.TraceID().String(),
	}))
}
