package traicing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(cnf Config) (opentracing.Tracer, io.Closer, error) { // nolint unused
	cfg := &config.Configuration{
		ServiceName: cnf.ServiceName,
		RPCMetrics:  true,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cnf.URI,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, nil, err
	}

	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)

	return tracer, closer, nil
}

// WithTraicer set logger
func WithTraicer(ctx context.Context, traicer opentracing.Tracer) context.Context { // nolint unused
	return context.WithValue(ctx, keyTraicer, traicer)
}

// GetTraicer return logger
func GetTraicer(ctx context.Context) opentracing.Tracer { // nolint unused
	return ctx.Value(keyTraicer).(opentracing.Tracer)
}
