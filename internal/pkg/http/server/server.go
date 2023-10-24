package http_server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

func New(ctx context.Context, h http.Handler, config Config, tracer trace.TracerProvider) *http.Server {
	viper.SetDefault("HTTP_SERVER_READ_TIMEOUT", "5s")        // the maximum duration for reading the entire request, including the body
	viper.SetDefault("HTTP_SERVER_WRITE_TIMEOUT", "5s")       // the maximum duration before timing out writes of the response
	viper.SetDefault("HTTP_SERVER_IDLE_TIMEOUT", "30s")       // the maximum amount of time to wait for the next request when keep-alive is enabled
	viper.SetDefault("HTTP_SERVER_READ_HEADER_TIMEOUT", "2s") // the amount of time allowed to read request headers

	handler := http.TimeoutHandler(h, config.Timeout, fmt.Sprintf(`{"error": %q}`, TimeoutMessage))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: handler,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},

		ReadTimeout:       viper.GetDuration("HTTP_SERVER_READ_TIMEOUT"),
		WriteTimeout:      config.Timeout + viper.GetDuration("HTTP_SERVER_WRITE_TIMEOUT"),
		IdleTimeout:       viper.GetDuration("HTTP_SERVER_IDLE_TIMEOUT"),
		ReadHeaderTimeout: viper.GetDuration("HTTP_SERVER_READ_HEADER_TIMEOUT"),
	}

	if tracer != nil {
		server.Handler = otelhttp.NewHandler(handler, "")
	}

	return server
}
