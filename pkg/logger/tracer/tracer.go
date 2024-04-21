package tracer

import (
	"context"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

func NewTraceFromContext(
	ctx context.Context, //nolint:contextcheck,maintidx // contextcheck: ctx is not nil
	msg string,
	tags []attribute.KeyValue,
	fields ...field.Fields,
) ([]field.Fields, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	_, span := otel.Tracer("logger").Start(ctx, getNameFunc())
	defer span.End()

	span.SetAttributes(ZapFieldsToOpenTelemetry(fields...)...)
	span.SetAttributes(attribute.String("log", msg))
	span.SetAttributes(tags...)

	// Get span ID
	if len(fields) == 0 {
		fields = make([]field.Fields, 0)
	}

	fields = append(fields, field.Fields{ // nozero
		"traceID": span.SpanContext().TraceID().String(),
	})

	return fields, nil
}

func getNameFunc() string {
	// at least one entry needed
	pc := make([]uintptr, 10) //nolint:mnd,revive // 10
	runtime.Callers(4, pc)    //nolint:mnd,revive // 4

	f := runtime.FuncForPC(pc[0])

	return f.Name()
}

// ZapFieldsToOpenTelemetry returns a table of standard openTelemetry field based on
// the inputed table of Zap field.
func ZapFieldsToOpenTelemetry(fields ...field.Fields) []attribute.KeyValue {
	openTelemetryFields := make([]attribute.KeyValue, 0, len(fields))

	for key := range fields {
		for k := range fields[key] {
			switch v := fields[key][k].(type) {
			case string:
				openTelemetryFields = append(openTelemetryFields, attribute.String(k, v))
			case bool:
				openTelemetryFields = append(openTelemetryFields, attribute.Bool(k, v))
			case int:
				openTelemetryFields = append(openTelemetryFields, attribute.Int(k, v))
			case int32:
				openTelemetryFields = append(openTelemetryFields, attribute.Int(k, int(v)))
			case int64:
				openTelemetryFields = append(openTelemetryFields, attribute.Int64(k, v))
			case error:
				openTelemetryFields = append(openTelemetryFields, attribute.String(k, v.Error()))
			}
		}
	}

	return openTelemetryFields
}
