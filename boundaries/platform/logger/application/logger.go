package logger_application

import (
	"context"
	"fmt"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1"
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

func (s *Service) Log(ctx context.Context, link *domain.Link) {
	s.log.InfoWithContext(ctx, fmt.Sprintf("GET URL: %s", link.GetUrl()))
}
