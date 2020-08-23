package metadata

import (
	"context"

	rpc "github.com/batazor/shortlink/internal/metadata/domain"
	"github.com/batazor/shortlink/internal/store"
)

type Store interface {
	Get(context.Context, string) (rpc.Meta, error)
	Set(context.Context, rpc.Meta) (rpc.Meta, error)
}

type MetaStore struct {
	store store.DB
}
