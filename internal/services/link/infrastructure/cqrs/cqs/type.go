package cqs

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
)

type Repository interface {
	Init(ctx context.Context, db *db.Store) error

	Add(ctx context.Context, data *v1.Link) (*v1.Link, error)
	Update(ctx context.Context, data *v1.Link) (*v1.Link, error)
	Delete(ctx context.Context, id string) error
}

// Store abstract type
type Store struct { // nolint unused
	typeStore string
	Repository
}
