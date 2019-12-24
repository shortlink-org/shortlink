package main

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
)

// Service - heplers
type Service struct {
	log logger.Logger
	mq  *mq.MQ
}

func (s *Service) initLogger() {
	log, err := di.InitLogger()
	if err != nil {
		panic(err)
	}

	s.log = *log
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

	// Add MQ
	s.initMQ(ctx)

	test := make(chan []byte)

	go func() {
		if s.mq != nil {
			service := *s.mq
			if err := service.Subscribe(test); err != nil {
				s.log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			s.log.Info(fmt.Sprintf("GET: %s", string(<-test)))
		}
	}()
}

// Stop - stop this a service
func (s *Service) Stop() {
	// flushes buffer, if any
	s.log.Close()
}
