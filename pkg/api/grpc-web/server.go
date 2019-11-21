package grpcweb

import (
	"context"
	"fmt"
	"github.com/batazor/shortlink/internal/freeport"
	"github.com/batazor/shortlink/pkg/api"
	"net"
	"net/http"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/store"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// API ...
type API struct { // nolint unused
	store store.DB
	ctx   context.Context
}

// Run HTTP-server
func (api *API) Run(ctx context.Context, db store.DB, config api.Config) error {
	api.ctx = ctx
	api.store = db

	// Get free port
	port, err := freeport.GetFreePort()
	if err != nil {
		return err
	}

	log := logger.GetLogger(ctx)
	log.Info(fmt.Sprintf("Run gRPC-GateWay on localhost:%d", port))

	// Rug gRPC
	go func() {
		if errRunGRPC := api.runGRPC(port); errRunGRPC != nil {
			log.Fatal(errRunGRPC.Error())
		}
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	gw := runtime.NewServeMux(
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = RegisterLinkHandlerFromEndpoint(ctx, gw, fmt.Sprintf("localhost:%d", port), opts)
	if err != nil {
		return err
	}

	srv := http.Server{Addr: fmt.Sprintf(":%d", config.Port), Handler: gw}

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
