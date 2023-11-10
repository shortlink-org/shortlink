/*
Tracing wrapping
*/
package traicing

import (
	"context"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/common"
)

// Init returns an instance of Tracer Provider that samples 100% of traces and logs all spans to stdout.
func Init(ctx context.Context, cnf Config, log logger.Logger) (*trace.TracerProvider, func(), error) {
	// Setup resource.
	res, err := common.NewResource(ctx, cnf.ServiceName, cnf.ServiceVersion)
	if err != nil {
		return nil, nil, err
	}

	// Setup trace provider.
	tp, err := newTraceProvider(ctx, res, cnf.URI)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		err := tp.Shutdown(ctx)
		if err != nil {
			log.Error(`Tracing disable`, field.Fields{
				"uri": cnf.URI,
				"err": err,
			})
		}
	}

	// Register the global Tracer provider
	// TODO: fix this
	// otel.SetTracerProvider(otelpyroscope.NewTracerProvider(
	// 	tp,
	// 	otelpyroscope.WithAppName(cnf.ServiceName),
	// 	otelpyroscope.WithPyroscopeURL(cnf.PyroscopeURI),
	// 	otelpyroscope.WithRootSpanOnly(true),
	// 	otelpyroscope.WithAddSpanName(true),
	// 	otelpyroscope.WithProfileURL(true),
	// 	otelpyroscope.WithProfileBaselineURL(true),
	// ))

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
}

func newTraceProvider(ctx context.Context, res *resource.Resource, uri string) (*trace.TracerProvider, error) {
	viper.SetDefault("TRACING_INITIAL_INTERVAL", "5s")
	viper.SetDefault("TRACING_MAX_INTERVAL", "30s")
	viper.SetDefault("TRACING_MAX_ELAPSED_TIME", "1m")

	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(uri),
		otlptracegrpc.WithRetry(otlptracegrpc.RetryConfig{
			Enabled:         true,
			InitialInterval: viper.GetDuration("TRACING_INITIAL_INTERVAL"),
			MaxInterval:     viper.GetDuration("TRACING_MAX_INTERVAL"),
			MaxElapsedTime:  viper.GetDuration("TRACING_MAX_ELAPSED_TIME"),
		}),
	)
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter, trace.WithBatchTimeout(viper.GetDuration("TRACING_INITIAL_INTERVAL"))),
		trace.WithResource(res),
	)

	return traceProvider, nil
}
