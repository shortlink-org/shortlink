package main

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/internal/traicing"
	"github.com/batazor/shortlink/pkg/api"
)

// Service - heplers
type Service struct {
	log         logger.Logger
	tracer      opentracing.Tracer
	tracerClose func() error
	db          store.DB
	mq          *mq.MQ
	api         api.Server
}

func (s *Service) initLogger() {
	log, err := di.InitLogger()
	if err != nil {
		panic(err)
	}

	s.log = *log
}

func (s *Service) initTracer() {
	tracer, tracerClose, err := di.InitTracer()
	if err != nil {
		panic(err)
	}

	s.tracer = *tracer
	s.tracerClose = tracerClose
}

func (s *Service) initMonitoring() *http.ServeMux {
	commonMux := di.InitMonitoring()

	return commonMux
}

func (s *Service) initMQ(ctx context.Context) {
	service, err := di.InitMQ(ctx)
	if err != nil {
		panic(err)
	}

	if service != nil {
		s.mq = service
		return
	}

	s.log.Info("MQ Disabled")
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

	// Add Store
	var st store.Store
	s.db = st.Use(ctx)

	// Add MQ
	s.initMQ(ctx)

	// Monitoring endpoints
	monitoringServer := s.initMonitoring()
	go http.ListenAndServe("0.0.0.0:9090", monitoringServer) // nolint errcheck

	// Run API server
	s.api.RunAPIServer(ctx)
}

// Stop - stop this a service
func (s *Service) Stop() {
	// close DB
	if err := s.db.Close(); err != nil {
		s.log.Error(err.Error())
	}

	// close tracer
	if err := s.tracerClose(); err != nil {
		s.log.Error(err.Error())
	}

	// flushes buffer, if any
	s.log.Close()
}
