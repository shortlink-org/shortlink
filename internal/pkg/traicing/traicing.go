/*
Tracing wrapping
*/
package traicing

import (
	"errors"
	"io"
	"net"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	zapJaeger "github.com/uber/jaeger-client-go/log/zap"
	"go.uber.org/zap"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(cnf Config, log logger.Logger) (*opentracing.Tracer, io.Closer, error) {
	// Check lookup site
	addr := strings.Split(cnf.URI, ":")
	if len(addr) == 0 {
		return nil, nil, errors.New("Not found jaeger URI host:port")
	}
	_, err := net.LookupIP(addr[0])
	if err != nil {
		// Ignore error with lookup Jaeger
		log.Warn("don't lookup Jaeger", field.Fields{"addr": cnf.URI})

		t := &jaeger.Tracer{}
		return nil, t, nil // nolint nilerr
	}

	cfg := &config.Configuration{
		ServiceName: cnf.ServiceName,
		RPCMetrics:  true,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: cnf.URI,
		},
	}
	zapLogger := log.Get().(*zap.Logger)
	tracer, closer, err := cfg.NewTracer(config.Logger(zapJaeger.NewLogger(zapLogger)))
	if err != nil {
		return nil, nil, err
	}

	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)

	return &tracer, closer, nil
}
