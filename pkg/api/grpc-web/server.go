package grpcweb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/logger"
	api_type "github.com/batazor/shortlink/pkg/api/type"

	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// API ...
type API struct { // nolint unused
	ctx  context.Context
	http http.Server
	RPC  *di.RPCServer
}

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

// Run HTTP-server
func (api *API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer opentracing.Tracer) error {
	api.ctx = ctx

	service := &LinkService{
		GetLinks:   api.GetLinks,
		GetLink:    api.GetLink,
		CreateLink: api.CreateLink,
		DeleteLink: api.DeleteLink,
	}

	// Rug gRPC
	RegisterLinkService(api.RPC.Server, service)
	api.RPC.Run()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	gw := runtime.NewServeMux(
		// Register custom error handler
		runtime.WithErrorHandler(api.CustomHTTPError),
	)

	// DefaultContextTimeout is used for gRPC call context.WithTimeout whenever a Grpc-Timeout inbound
	// header isn't present. If the value is 0 the sent `context` will not have a timeout.
	runtime.DefaultContextTimeout = config.Timeout * time.Second

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(tracer),
			),
		),
	}
	err := RegisterLinkHandlerFromEndpoint(ctx, gw, api.RPC.Endpoint, opts)
	if err != nil {
		return err
	}

	api.http = http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: api.tracingWrapper(gw),

		ReadTimeout:       1 * time.Second,                     // the maximum duration for reading the entire request, including the body
		WriteTimeout:      (config.Timeout + 30) * time.Second, // the maximum duration before timing out writes of the response
		IdleTimeout:       30 * time.Second,                    // the maximum amount of time to wait for the next request when keep-alive is enabled
		ReadHeaderTimeout: 2 * time.Second,                     // the amount of time allowed to read request headers
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Info(fmt.Sprintf("API run on port %d", config.Port))
	err = api.http.ListenAndServe()
	return err
}

// Close ...
func (api *API) Close() error {
	if err := api.http.Close(); err != nil {
		return err
	}

	return nil
}

func (api *API) tracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err == nil || err == opentracing.ErrSpanContextNotFound {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				"ServeHTTP",
				// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}

func (api *API) CustomHTTPError(_ context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s := status.Convert(err)
	pb := s.Proto()

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaler.ContentType(pb)
	w.Header().Set("Content-Type", contentType)

	w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))

	jErr := json.NewEncoder(w).Encode(customError{
		Error: status.Convert(err).Message(),
	})

	if jErr != nil {
		_, _ = w.Write([]byte(fallback)) // nolint gosec
	}
}
