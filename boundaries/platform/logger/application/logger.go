package logger_application

import (
	"context"
	"fmt"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

type Service struct {
	log logger.Logger
}

func New(log logger.Logger) (*Service, error) {
	service := &Service{
		log: log,
	}

	return service, nil
}

func (s *Service) Log(ctx context.Context, payload any) {
	s.log.InfoWithContext(ctx, fmt.Sprintf("GET URL: %s", payload))
}
