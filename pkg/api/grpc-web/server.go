package grpcweb

import (
	"context"
	"fmt"
	"github.com/batazor/shortlink/internal/traicing"
	"github.com/opentracing/opentracing-go/ext"
	"net"
	"net/http"

	"github.com/batazor/shortlink/internal/freeport"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/pkg/api"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// API ...
type API struct { // nolint unused
	store store.DB
	ctx   context.Context
}

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

// Run HTTP-server
func (api *API) Run(ctx context.Context, db store.DB, config api.Config) error {
	api.ctx = ctx
	api.store = db

	// Get free port
	port, err := freeport.GetFreePort()
	if err != nil {
		return err
	}

	// Get logger
	log := logger.GetLogger(ctx)
	log.Info(fmt.Sprintf("Run gRPC-GateWay on localhost:%d", port))

	// Rug gRPC
	go func() {
		if errRunGRPC := api.runGRPC(port); errRunGRPC != nil {
			log.Fatal(errRunGRPC.Error())
		}
	}()

	// Get tracer
	tracer := traicing.GetTraicer(ctx)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	gw := runtime.NewServeMux(
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(tracer),
			),
		),
	}
	err = RegisterLinkHandlerFromEndpoint(ctx, gw, fmt.Sprintf("localhost:%d", port), opts)
	if err != nil {
		return err
	}

	srv := http.Server{Addr: fmt.Sprintf(":%d", config.Port), Handler: api.tracingWrapper(gw)}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = srv.ListenAndServe()
	return err
}

// runGRPC ...
func (api *API) runGRPC(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterLinkServer(grpcServer, api)
	err = grpcServer.Serve(lis)
	return err
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
