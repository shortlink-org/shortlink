package main

import (
	"context"
	"fmt"
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
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func() {
		// flushes buffer, if any
		if error := logger.Sync(); error != nil {
			// TODO: use logger
			fmt.Println(error.Error())
		}
	}()

	// Add context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Add logger
	ctx = log.WithLogger(ctx, *logger)

	// Add Tracer
	tracer, closer, err := traicing.Init()
	defer func() {
		// TODO: use logger
		if error := closer.Close(); error != nil {
			fmt.Println(error.Error())
		}
	}()
	if err != nil {
		logger.Error(err.Error())
	}
	ctx = traicing.WithTraicer(ctx, tracer)

	// Test Event
	tracer.StartSpan("test").Finish()

	// start HTTP-server
	var API api.API
	serverType := "graphql"

	switch serverType {
	case "http-chi":
		API = &httpchi.API{}
	case "gRPC-web":
		API = &grpcweb.API{}
	case "graphql":
		API = &graphql.API{}
	default:
		API = &httpchi.API{}
	}

	if err := API.Run(ctx); err != nil {
		logger.Panic(err.Error())
	}
}
