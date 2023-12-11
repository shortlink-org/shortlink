package link_cqrs

import (
	"context"

	link "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	metadata "github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/domain/metadata/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
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

func (s *Service) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	switch event {
	case link.METHOD_ADD:
		{
			_, err := s.cqsStore.LinkAdd(ctx, payload.(*link.Link))
			if err != nil {
				s.log.ErrorWithContext(ctx, err.Error())
			}

			return notify.Response[any]{}
		}
	case link.METHOD_UPDATE:
		{
			_, err := s.cqsStore.LinkUpdate(ctx, payload.(*link.Link))
			if err != nil {
				s.log.ErrorWithContext(ctx, err.Error())
			}

			return notify.Response[any]{}
		}
	case link.METHOD_DELETE:
		{
			err := s.cqsStore.LinkDelete(ctx, payload.(string)) //nolint:forcetypeassert // simple type assertion
			if err != nil {
				s.log.ErrorWithContext(ctx, err.Error())
			}

			return notify.Response[any]{}
		}
	case metadata.METHOD_ADD:
		fallthrough
	case metadata.METHOD_UPDATE:
		{
			_, err := s.cqsStore.MetadataUpdate(ctx, payload.(*metadata.Meta))
			if err != nil {
				s.log.ErrorWithContext(ctx, err.Error())
			}

			return notify.Response[any]{}
		}
	default:
		return notify.Response[any]{}
	}
}
