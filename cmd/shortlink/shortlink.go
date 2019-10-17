package main

import (
	"context"
	"github.com/batazor/shortlink/pkg/api"
	"github.com/batazor/shortlink/pkg/api/grpc-web"
	"github.com/batazor/shortlink/pkg/api/http-chi"
	log "github.com/batazor/shortlink/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Add context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Add logger
	ctx = log.WithLogger(ctx, *logger)

	// start HTTP-server
	var api api.API
	serverType := "gRPC-web"

	switch serverType {
	case "http-chi":
		api = &http_chi.API{}
	case "gRPC-web":
		api = &grpc_web.API{}
	}

	if err := api.Run(ctx); err != nil {
		logger.Panic(err.Error())
	}
}
