package server

import (
	"time"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

type Option func(*Server)

// WithLogger sets the logger for the server
func WithLogger(logger logger.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

// WithElectionTime sets the election time for the server
func WithElectionTime(duration time.Duration) Option {
	return func(s *Server) {
		s.electionResetTimer = duration
	}
}
