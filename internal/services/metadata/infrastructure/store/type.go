package meta_store

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/notify"
	rpc "github.com/batazor/shortlink/internal/services/metadata/domain"
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
	notify.Subscriber // Observer interface for subscribe on system event
}
