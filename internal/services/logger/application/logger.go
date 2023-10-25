package logger_application

import (
	"context"
	"fmt"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
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
