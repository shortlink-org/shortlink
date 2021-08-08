package query

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
)

type Repository interface {
	Init(ctx context.Context, db *db.Store) error

	Get(ctx context.Context, id string) (*v1.Link, error)
	List(ctx context.Context, filter *query.Filter) (*v1.Links, error)
}

// Store abstract type
type Store struct { // nolint unused
	typeStore string
	Repository
}
