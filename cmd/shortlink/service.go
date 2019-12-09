package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/mq/kafka"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/internal/traicing"
	"github.com/batazor/shortlink/pkg/api"
)

// Service - heplers
type Service struct {
	log         logger.Logger
	tracer      opentracing.Tracer
	tracerClose io.Closer
	db          store.DB
	mq          mq.MQ
	api         api.Server
}

func (s *Service) initLogger() {
	var err error

	viper.SetDefault("LOG_LEVEL", logger.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := logger.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	if s.log, err = logger.NewLogger(logger.Zap, conf); err != nil {
		panic(err)
	}
}

func (s *Service) initTracer() {
	var err error

	viper.SetDefault("TRACER_SERVICE_NAME", "ShortLink")
	viper.SetDefault("TRACER_URI", "localhost:6831")

	config := traicing.Config{
		ServiceName: viper.GetString("TRACER_SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	if s.tracer, s.tracerClose, err = traicing.Init(config); err != nil {
		s.log.Error(err.Error())
	}
}

func (s *Service) initMonitoring() *http.ServeMux {
	// Create a new Prometheus registry
	registry := prometheus.NewRegistry()

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(registry, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	// Create an "common" listener
	commonMux := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	commonMux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	// Expose a liveness check on /live
	commonMux.HandleFunc("/live", health.LiveEndpoint)

	// Expose a readiness check on /ready
	commonMux.HandleFunc("/ready", health.ReadyEndpoint)

	return commonMux
}

func (s *Service) initMQ(ctx context.Context) {
	s.mq = &kafka.Kafka{}
	if err := s.mq.Init(ctx); err != nil {
		panic(err)
	}

	s.log.Info("Run MQ")
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
	if err := s.tracerClose.Close(); err != nil {
		s.log.Error(err.Error())
	}

	// flushes buffer, if any
	s.log.Close()
}
