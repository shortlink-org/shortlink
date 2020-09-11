package tracer

import (
	"context"
	"runtime"

	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"

	"github.com/batazor/shortlink/internal/logger/field"
)

func NewTraceFromContext(ctx context.Context, msg string, fields ...field.Fields) error {
	if ctx == nil {
		ctx = context.Background()
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, getNameFunc())
	defer span.Finish()

	span.LogFields(ZapFieldsToOpentracing(fields...)...)
	span.LogFields(opentracinglog.String("log", msg))

	//if traceID, ok := span.Context().(jaeger.SpanContext); ok {
	//	fields = append(fields, zap.String("traceID", traceID.TraceID().String()))
	//}

	return nil
}

func getNameFunc() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(4, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// ZapFieldsToOpentracing returns a table of standard opentracing field based on
// the inputed table of Zap field.
func ZapFieldsToOpentracing(fields ...field.Fields) []opentracinglog.Field {
	opentracingFields := make([]opentracinglog.Field, len(fields))

	for key, _ := range fields {
		for k, _ := range fields[key] {
			switch v := fields[key][k].(type) {
			case string:
				opentracingFields = append(opentracingFields, opentracinglog.String(k, v))
			case bool:
				opentracingFields = append(opentracingFields, opentracinglog.Bool(k, v))
			case int:
				opentracingFields = append(opentracingFields, opentracinglog.Int(k, v))
			case int32:
				opentracingFields = append(opentracingFields, opentracinglog.Int32(k, v))
			case int64:
				opentracingFields = append(opentracingFields, opentracinglog.Int64(k, v))
			case error:
				opentracingFields = append(opentracingFields, opentracinglog.String(k, v.Error()))
			}
		}
	}

	return opentracingFields
}
