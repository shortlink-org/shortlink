package http_server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

func New(ctx context.Context, h http.Handler, config Config, tracer trace.TracerProvider) *http.Server {
	handler := http.TimeoutHandler(h, config.Timeout, fmt.Sprintf(`{"error": "%s"}`, TimeoutMessage))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: handler,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},

		ReadTimeout:  5 * time.Second,                // the maximum duration for reading the entire request, including the body
		WriteTimeout: config.Timeout + 5*time.Second, // the maximum duration before timing out writes of the response
		// the maximum amount of time to wait for the next request when keep-alive is enabled
		IdleTimeout: 30 * time.Second, //nolint:gomnd
		// the amount of time allowed to read request headers
		ReadHeaderTimeout: 2 * time.Second, //nolint:gomnd
	}

	if tracer != nil {
		server.Handler = otelhttp.NewHandler(handler, "")
	}

	return server
}
