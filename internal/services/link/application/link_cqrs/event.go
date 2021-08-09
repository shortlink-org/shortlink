package link_cqrs

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	metadata_domain "github.com/batazor/shortlink/internal/services/metadata/domain"
)

func (s *Service) EventHandler() {
	// Subscribe to Event
	// Link
	notify.Subscribe(api_domain.METHOD_ADD, s)
	notify.Subscribe(api_domain.METHOD_UPDATE, s)
	notify.Subscribe(api_domain.METHOD_DELETE, s)

	// Metadata
	notify.Subscribe(metadata_domain.METHOD_ADD, s)
	notify.Subscribe(metadata_domain.METHOD_UPDATE, s)
	notify.Subscribe(metadata_domain.METHOD_DELETE, s)

	// Proxy
}

func (s *Service) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case api_domain.METHOD_ADD:
		{
			_, err := s.cqsStore.LinkAdd(ctx, payload.(*v1.Link))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case api_domain.METHOD_UPDATE:
		{
			_, err := s.cqsStore.LinkUpdate(ctx, payload.(*v1.Link))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case api_domain.METHOD_DELETE:
		{
			err := s.cqsStore.LinkDelete(ctx, payload.(string))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case metadata_domain.METHOD_ADD:
		fallthrough
	case metadata_domain.METHOD_UPDATE:
		{
			_, err := s.cqsStore.MetadataUpdate(ctx, payload.(*metadata_domain.Meta))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	default:
		return notify.Response{}
	}
}
