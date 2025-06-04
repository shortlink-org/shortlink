package ram

import (
	"context"
	"sync"

	rpc "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/domain/metadata/v1"
	errors "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/repository/store/error"
)

type Store struct {
	// sync.Map solver problem with cache contention
	metadata sync.Map
}

// Get - get metadata by id
func (s *Store) Get(_ context.Context, id string) (*rpc.Meta, error) {
	response, ok := s.metadata.Load(id)
	if !ok {
		return nil, &errors.MetadataNotFoundByIdError{ID: id}
	}

	v, ok := response.(*rpc.Meta)
	if !ok {
		return nil, &errors.MetadataNotFoundByIdError{ID: id}
	}

	return v, nil
}

// Add - write new metadata for a link
func (s *Store) Add(_ context.Context, source *rpc.Meta) error {
	s.metadata.Store(source.GetId(), source)

	return nil
}
