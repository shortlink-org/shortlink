package metadata

import (
	"context"

	rpc "github.com/batazor/shortlink/internal/metadata/domain"
)

type Store interface {
	Get(context.Context, string) (rpc.Meta, error)
	Set(context.Context, rpc.Meta) (rpc.Meta, error)
}

type MetaStore struct{}
