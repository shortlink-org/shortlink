package traicing_di

import (
	"context"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/observability/traicing"
	error_di "github.com/shortlink-org/shortlink/pkg/di/pkg/error"
)

// New returns a new instance of the TracerProvider.
//
//nolint:ireturn // It's make by specification
func New(ctx context.Context, log logger.Logger) (trace.TracerProvider, func(), error) {
	viper.SetDefault("TRACER_URI", "localhost:4317")                     // Tracing addr:host
	viper.SetDefault("PYROSCOPE_URI", "http://pyroscope.pyroscope:4040") // Pyroscope addr:host

	config := traicing.Config{
		ServiceName:    viper.GetString("SERVICE_NAME"),
		ServiceVersion: viper.GetString("SERVICE_VERSION"),
		URI:            viper.GetString("TRACER_URI"),
		PyroscopeURI:   viper.GetString("PYROSCOPE_URI"),
	}

	tracer, tracerClose, err := traicing.Init(ctx, config, log)
	if err != nil {
		return nil, nil, &error_di.BaseError{Err: err}
	}

	if tracer == nil {
		return nil, func() {}, nil
	}

	return tracer, tracerClose, nil
}
