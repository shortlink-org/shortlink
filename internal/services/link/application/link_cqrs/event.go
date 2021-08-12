package link_cqrs

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/notify"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	v12 "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"
)

func (s *Service) EventHandlers() {
	// Subscribe to Event
	// Link
	notify.Subscribe(v1.METHOD_ADD, s)
	notify.Subscribe(v1.METHOD_UPDATE, s)
	notify.Subscribe(v1.METHOD_DELETE, s)

	// Metadata
	notify.Subscribe(v12.METHOD_ADD, s)
	notify.Subscribe(v12.METHOD_UPDATE, s)
	notify.Subscribe(v12.METHOD_DELETE, s)

	// Proxy
}

func (s *Service) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case v1.METHOD_ADD:
		{
			_, err := s.cqsStore.LinkAdd(ctx, payload.(*v1.Link))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case v1.METHOD_UPDATE:
		{
			_, err := s.cqsStore.LinkUpdate(ctx, payload.(*v1.Link))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case v1.METHOD_DELETE:
		{
			err := s.cqsStore.LinkDelete(ctx, payload.(string))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	case v12.METHOD_ADD:
		fallthrough
	case v12.METHOD_UPDATE:
		{
			_, err := s.cqsStore.MetadataUpdate(ctx, payload.(*v12.Meta))
			if err != nil {
				s.logger.ErrorWithContext(ctx, err.Error())
			}
			return notify.Response{}
		}
	default:
		return notify.Response{}
	}
}
