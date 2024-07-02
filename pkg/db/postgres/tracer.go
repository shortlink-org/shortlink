package postgres

import (
	"context"

	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	trace.TracerProvider
}

func (t *Tracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		semconv.DBSystemPostgreSQL,
		semconv.DBStatementKey.String(data.SQL),
	)

	return ctx
}

func (t *Tracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	span := trace.SpanFromContext(ctx)

	if data.Err != nil {
		span.RecordError(data.Err)
	}

	span.End()
}
