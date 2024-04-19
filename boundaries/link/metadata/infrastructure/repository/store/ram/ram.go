package ram

import (
	"context"
	"sync"

	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types"
	rpc "github.com/shortlink-org/shortlink/boundaries/link/metadata/domain/metadata/v1"
)

type Store struct {
	// sync.Map solver problem with cache contention
	metadata sync.Map
}

// Get - get metadata by id
func (s *Store) Get(_ context.Context, id string) (*rpc.Meta, error) {
	response, ok := s.metadata.Load(id)
	if !ok {
		return nil, &types.NotFoundByHashError{Hash: id}
	}

	v, ok := response.(*rpc.Meta)
	if !ok {
		return nil, &types.NotFoundByHashError{Hash: id}
	}

	return v, nil
}

// Add - write new metadata for a link
func (s *Store) Add(_ context.Context, source *rpc.Meta) error {
	s.metadata.Store(source.GetId(), source)

	return nil
}
