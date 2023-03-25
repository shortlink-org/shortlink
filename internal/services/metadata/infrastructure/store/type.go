package meta_store

import (
	"context"

	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	rpc "github.com/shortlink-org/shortlink/internal/services/metadata/domain/metadata/v1"
)

type Repository interface {
	Get(context.Context, string) (*rpc.Meta, error)
	Add(context.Context, *rpc.Meta) error
}

// Store abstract type
type MetaStore struct {
	typeStore string
	Store     Repository

	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]
}
