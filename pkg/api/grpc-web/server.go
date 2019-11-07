package grpcweb

import (
	"context"
	"net"
	"net/http"

	"github.com/batazor/shortlink/pkg/internal/store"
	"github.com/batazor/shortlink/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// API ...
type API struct {
	store store.DB
	ctx   context.Context
}

// Run HTTP-server
func (api *API) Run(ctx context.Context) error {
	var st store.Store

	api.ctx = ctx
	api.store = st.Use()

	logger := logger.GetLogger(ctx)
	logger.Info("Run gRPC-GateWay API")

	PORT := "7070"

	// Rug gRPC
	go func() {
		err := api.runGRPC()
		if err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	gw := runtime.NewServeMux(
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := RegisterLinkHandlerFromEndpoint(ctx, gw, "localhost:9090", opts)
	if err != nil {
		return err
	}

	srv := http.Server{Addr: ":" + PORT, Handler: gw}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = srv.ListenAndServe()
	return err
}

// runGRPC ...
func (api *API) runGRPC() error {
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterLinkServer(grpcServer, api)
	err = grpcServer.Serve(lis)
	return err
}
