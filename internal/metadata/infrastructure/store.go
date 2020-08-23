package store

import (
	"context"

	"github.com/batazor/shortlink/internal/metadata/domain"
	"github.com/batazor/shortlink/internal/store"
)

type Store interface {
	Get(context.Context, string) (*rpc.Meta, error)
	Add(context.Context, rpc.Meta) error
}

type MetaStore struct {
	store.DB
}
