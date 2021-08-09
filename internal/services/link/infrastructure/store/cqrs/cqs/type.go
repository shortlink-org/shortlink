package cqs

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/metadata/domain"
)

type Repository interface {
	Init(ctx context.Context, db *db.Store) error

	LinkAdd(ctx context.Context, data *v1.Link) (*v1.Link, error)
	LinkUpdate(ctx context.Context, data *v1.Link) (*v1.Link, error)
	LinkDelete(ctx context.Context, id string) error

	MetadataUpdate(ctx context.Context, data *domain.Meta) (*domain.Meta, error)
}

// Store abstract type
type Store struct { // nolint unused
	typeStore string
	Repository
}
