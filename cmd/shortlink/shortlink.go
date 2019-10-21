package main

import (
	"context"
	"github.com/batazor/shortlink/pkg/api"
	"github.com/batazor/shortlink/pkg/api/graphql"
	"github.com/batazor/shortlink/pkg/api/grpc-web"
	"github.com/batazor/shortlink/pkg/api/http-chi"
	log "github.com/batazor/shortlink/pkg/logger"
	"github.com/batazor/shortlink/pkg/traicing"
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

	// Add Tracer
	tracer, closer, err := traicing.Init("hello-world")
	defer closer.Close()
	if err != nil {
		logger.Error(err.Error())
	}
	ctx = traicing.WithTraicer(ctx, tracer)

	// Test Event
	tracer.StartSpan("test").Finish()

	// start HTTP-server
	var api api.API
	serverType := "graphql"

	switch serverType {
	case "http-chi":
		api = &http_chi.API{}
	case "gRPC-web":
		api = &grpc_web.API{}
	case "graphql":
		api = &graphql.API{}
	default:
		api = &http_chi.API{}
	}

	if err := api.Run(ctx); err != nil {
		logger.Panic(err.Error())
	}
}
