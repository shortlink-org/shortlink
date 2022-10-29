package traicing_di

import (
	"context"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/traicing"
)

func New(ctx context.Context, log logger.Logger) (*trace.TracerProvider, func(), error) {
	viper.SetDefault("TRACER_URI", "localhost:14268") // Tracing addr:host

	config := traicing.Config{
		ServiceName: viper.GetString("SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	tracer, tracerClose, err := traicing.Init(ctx, config, log)
	if err != nil {
		return nil, nil, err
	}
	if tracer == nil {
		return nil, func() {}, nil
	}

	return &tracer, tracerClose, nil
}
