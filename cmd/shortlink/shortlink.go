package main

import (
	"context"
	"fmt"
	"github.com/batazor/shortlink/pkg/api"
	"github.com/batazor/shortlink/pkg/api/graphql"
	"github.com/batazor/shortlink/pkg/api/grpc-web"
	"github.com/batazor/shortlink/pkg/api/http-chi"
	"github.com/batazor/shortlink/pkg/logger"
	"github.com/batazor/shortlink/pkg/traicing"
)

func main() {
	// Logger
	log, err := logger.NewLogger(logger.Configuration{
		Level: logger.INFO_LEVEL,
	}, logger.Logrus)
	if err != nil {
		panic(err)
	}
	defer func() {
		if r := recover(); r != nil {
			log.Error(r.(string))
		}

		// flushes buffer, if any
		log.Close()
	}()

	// Add context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Add logger
	ctx = logger.WithLogger(ctx, log)

	// Add Tracer
	tracer, closer, err := traicing.Init()
	defer func() {
		// TODO: use logger
		if error := closer.Close(); error != nil {
			fmt.Println(error.Error())
		}
	}()
	if err != nil {
		log.Error(err.Error())
	}
	ctx = traicing.WithTraicer(ctx, tracer)

	// Test Event
	tracer.StartSpan("test").Finish()

	// start HTTP-server
	var API api.API
	serverType := "http-chi"

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
		log.Fatal(err.Error())
	}
}
