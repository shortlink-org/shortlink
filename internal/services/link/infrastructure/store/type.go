package store

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
)

type Repository interface {
	Init(ctx context.Context, db *db.Store) error
	Get(ctx context.Context, id string) (*v1.Link, error)
	List(ctx context.Context, filter *query.Filter) (*v1.Links, error)
	Add(ctx context.Context, data *v1.Link) (*v1.Link, error)
	Update(ctx context.Context, data *v1.Link) (*v1.Link, error)
	Delete(ctx context.Context, id string) error
}

// Store abstract type
type LinkStore struct { // nolint unused
	typeStore string
	Store     Repository

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}
