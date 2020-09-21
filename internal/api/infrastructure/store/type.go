package store

import (
	"context"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/notify"
)

type Repository interface {
	Init(ctx context.Context, db *db.Store) error
	Get(ctx context.Context, id string) (*link.Link, error)
	List(ctx context.Context, filter *query.Filter) ([]*link.Link, error)
	Add(ctx context.Context, data *link.Link) (*link.Link, error)
	Update(ctx context.Context, data *link.Link) (*link.Link, error)
	Delete(ctx context.Context, id string) error
}

// Store abstract type
type LinkStore struct { // nolint unused
	typeStore string
	Store     Repository

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}
