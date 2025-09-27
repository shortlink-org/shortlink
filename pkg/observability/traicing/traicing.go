/*
Tracing wrapping with Go 1.25 FlightRecorder support
*/
package traicing

import (
	"context"
	"log/slog"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/common"
)

// Init returns an instance of Tracer Provider that samples 100% of traces and logs all spans to stdout.
// It also initializes Go 1.25 FlightRecorder if configured.
func Init(ctx context.Context, cnf Config, log *logger.SlogLogger) (*trace.TracerProvider, func(), error) {
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

	// Initialize Flight Recorder if configured
	var flightRecorder *FlightRecorder
	if cnf.FlightRecorder != nil && cnf.FlightRecorder.Enabled {
		flightRecorder, err = NewFlightRecorder(*cnf.FlightRecorder, log)
		if err != nil {
			log.Error("Failed to create flight recorder", slog.Any("err", err))
		} else if flightRecorder != nil {
			if err := flightRecorder.Start(); err != nil {
				log.Error("Failed to start flight recorder", slog.Any("err", err))
			} else {
				// Set as global flight recorder for easy access
				SetGlobalFlightRecorder(flightRecorder)
			}
		}
	}

	cleanup := func() {
		// Stop flight recorder first
		if flightRecorder != nil {
			if err := flightRecorder.Stop(); err != nil {
				log.Error("Error stopping flight recorder", slog.Any("err", err))
			}
		}

		// Then shutdown trace provider
		errShutdown := tp.Shutdown(ctx)
		if errShutdown != nil {
			log.Error(`Tracing disable`, 
				slog.String("uri", cnf.URI),
				slog.Any("err", errShutdown))
		}
	}

	log.Info(`Tracing enable`, 
		slog.String("uri", cnf.URI),
		slog.Bool("flight_recorder", cnf.FlightRecorder != nil && cnf.FlightRecorder.Enabled))

	// Gracefully shutdown the trace provider and flight recorder on exit
	go func() {
		<-ctx.Done()

		// Stop flight recorder
		if flightRecorder != nil {
			if err := flightRecorder.Stop(); err != nil {
				log.Error("error stopping flight recorder", slog.String("err", err.Error()))
			}
		}

		// Shutdown will flush any remaining spans and shut down the exporter.
		if errShutdown := tp.Shutdown(ctx); errShutdown != nil {
			log.Error("error shutting down trace provider", slog.String("err", errShutdown.Error()))
		}
	}()

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
		trace.WithSampler(trace.AlwaysSample()),
	)

	otel.SetTracerProvider(traceProvider)

	// Register the W3C trace context and baggage propagators so data is propagated across services/processes
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return traceProvider, nil
}
