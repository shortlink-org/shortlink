package logger_application

import (
	"context"
	"fmt"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

type Service struct {
	logger logger.Logger
}

func New(logger logger.Logger) (*Service, error) {
	service := &Service{
		logger: logger,
	}

	return service, nil
}

func (s *Service) Log(ctx context.Context, link *v1.Link) {
	s.logger.InfoWithContext(ctx, fmt.Sprintf("GET URL: %s", link.GetUrl()))
}
