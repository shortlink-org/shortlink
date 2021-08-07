package link

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain"
)

func (s *Service) EventHandler() {
	// Subscribe to Event
	// Link
	notify.Subscribe(api_domain.METHOD_ADD, s)
	notify.Subscribe(api_domain.METHOD_UPDATE, s)
	notify.Subscribe(api_domain.METHOD_DELETE, s)

	// Metadata

	// Proxy
}

func (s *Service) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case api_domain.METHOD_ADD:
		{
			return notify.Response{}
		}
	case api_domain.METHOD_UPDATE:
		{
			return notify.Response{}
		}
	case api_domain.METHOD_DELETE:
		{
			return notify.Response{}
		}
	default:
		return notify.Response{}
	}
}
