package main

import (
	"context"
	"github.com/batazor/shortlink/pkg/api"
	"github.com/batazor/shortlink/pkg/api/graphql"
	grpcweb "github.com/batazor/shortlink/pkg/api/grpc-web"
	httpchi "github.com/batazor/shortlink/pkg/api/http-chi"
	"github.com/batazor/shortlink/pkg/logger"
	"github.com/batazor/shortlink/pkg/traicing"
	"github.com/opentracing/opentracing-go"
	"io"
)

type Service struct {
	log         logger.Logger
	tracer      opentracing.Tracer
	tracerClose io.Closer
}

func (s *Service) initLogger() {
	var err error
	conf := logger.Configuration{
		Level: logger.INFO_LEVEL,
	}

	if s.log, err = logger.NewLogger(logger.Zap, conf); err != nil {
		panic(err)
	}
}

func (s *Service) initTracer() {
	var err error
	if s.tracer, s.tracerClose, err = traicing.Init(); err != nil {
		s.log.Error(err.Error())
	}
}

// runAPIServer - start HTTP-server
func (s *Service) runAPIServer(ctx context.Context) {
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
		s.log.Fatal(err.Error())
	}
}

// Start - run this a service
func (s *Service) Start() {
	// Create a new context
	ctx := context.Background()

	// Logger
	s.initLogger()
	ctx = logger.WithLogger(ctx, s.log) // Add logger to context

	// Add Tracer
	s.initTracer()
	ctx = traicing.WithTraicer(ctx, s.tracer) // Add tracer to context

	// Run API server
	s.runAPIServer(ctx)
}

// Stop - stop this a service
func (s *Service) Stop() {
	if err := s.tracerClose.Close(); err != nil {
		s.log.Error(err.Error())
	}

	// flushes buffer, if any
	s.log.Close()
}
