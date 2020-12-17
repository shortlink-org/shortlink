package traicing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
)

func NewTracer(ctx context.Context, log logger.Logger) (*opentracing.Tracer, func(), error) {
	viper.SetDefault("TRACER_URI", "localhost:6831") // Tracing addr:host

	config := Config{
		ServiceName: viper.GetString("SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	tracer, tracerClose, err := Init(config, log)
	if err != nil {
		return nil, nil, err
	}
	if tracer == nil {
		return nil, func() {}, nil
	}

	cleanup := func() {
		if err := tracerClose.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return tracer, cleanup, nil
}
