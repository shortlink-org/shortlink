/*
Tracing wrapping
*/
package traicing

import (
	"context"

	otelpyroscope "github.com/pyroscope-io/otel-profiling-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(ctx context.Context, cnf Config, log logger.Logger) (trace.TracerProvider, func(), error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cnf.URI)))
	if err != nil {
		return nil, nil, err
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cnf.ServiceName),
		)),
	)

	cleanup := func() {
		_ = exporter.Shutdown(ctx)
		_ = tp.Shutdown(ctx)
	}

	// Register the global Tracer provider
	otel.SetTracerProvider(otelpyroscope.NewTracerProvider(
		tp,
		otelpyroscope.WithAppName(cnf.ServiceName),
		otelpyroscope.WithPyroscopeURL(cnf.URI),
		otelpyroscope.WithRootSpanOnly(true),
		otelpyroscope.WithAddSpanName(true),
		otelpyroscope.WithProfileURL(true),
		otelpyroscope.WithProfileBaselineURL(true),
	))

	// Register the W3C trace context and baggage propagators so data is propagated across services/processes
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	log.Info(`Tracing enable`, field.Fields{
		"uri": cnf.URI,
	})

	return tp, cleanup, nil

	// cfg := &config.Configuration{
	// 	RPCMetrics:  true,
	// 	Sampler: &config.SamplerConfig{
	// 		Type:  "const",
	// 		Param: 1,
	// 	},
	// 	Reporter: &config.ReporterConfig{
	// 		LogSpans:           false,
	// 		LocalAgentHostPort: cnf.URI,
	// 	},
	// }
	// zapLogger, ok := log.Get().(*zap.Logger)
	// if !ok {
	// 	return nil, nil, errors.New("incorrect type assertion")
	// }
	//
	// tracer, closer, err := cfg.NewTracer(config.Logger(zapJaeger.NewLogger(zapLogger)))
	// if err != nil {
	// 	return nil, nil, err
	// }
}
