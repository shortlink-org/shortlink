package main

import (
	"context"
	"time"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/mq/kafka"
	"github.com/spf13/viper"
)

// Service - heplers
type Service struct {
	log logger.Logger
	mq  mq.MQ
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

func (s *Service) initMQ(ctx context.Context) {
	s.mq = &kafka.Kafka{}
	if err := s.mq.Init(ctx); err != nil {
		panic(err)
	}
}

// Start - run this a service
func (s *Service) Start() {
	// Create a new context
	ctx := context.Background()

	// Logger
	s.initLogger()
	ctx = logger.WithLogger(ctx, s.log) // Add logger to context

	// Add MQ
	s.initMQ(ctx)
}

// Stop - stop this a service
func (s *Service) Stop() {
	// flushes buffer, if any
	s.log.Close()
}
