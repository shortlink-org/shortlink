package link_cqrs

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/notify"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	metadata "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"
)

func (s *Service) EventHandlers() {
	// Subscribe to Event
	// Link
	notify.Subscribe(link.METHOD_ADD, s)
	notify.Subscribe(link.METHOD_UPDATE, s)
	notify.Subscribe(link.METHOD_DELETE, s)

	// Metadata
	notify.Subscribe(metadata.METHOD_ADD, s)
	notify.Subscribe(metadata.METHOD_UPDATE, s)
	notify.Subscribe(metadata.METHOD_DELETE, s)

	// Proxy
}

func (s *Service) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case link.METHOD_ADD:
		{
			_, err := s.cqsStore.LinkAdd(ctx, payload.(*link.Link))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case link.METHOD_UPDATE:
		{
			_, err := s.cqsStore.LinkUpdate(ctx, payload.(*link.Link))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case link.METHOD_DELETE:
		{
			err := s.cqsStore.LinkDelete(ctx, payload.(string))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case metadata.METHOD_ADD:
		fallthrough
	case metadata.METHOD_UPDATE:
		{
			_, err := s.cqsStore.MetadataUpdate(ctx, payload.(*metadata.Meta))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	default:
		return notify.Response{}
	}
}
