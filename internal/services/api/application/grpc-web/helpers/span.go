package helpers

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RegisterSpan - inject spanId in response header
func RegisterSpan(ctx context.Context) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return
	}
	if traceID, okSpan := span.Context().(jaeger.SpanContext); okSpan {
		err := grpc.SendHeader(ctx, metadata.New(map[string]string{
			"span-id": traceID.SpanID().String(),
		}))
		if err != nil {
			return
		}
	}
}
