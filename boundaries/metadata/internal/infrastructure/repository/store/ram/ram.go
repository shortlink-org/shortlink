package ram

import (
	"context"
	"sync"

	rpc "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
	errors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/repository/store/error"
)

type Store struct {
	// sync.Map solver problem with cache contention
	metadata sync.Map
}

// Get - get metadata by id
func (s *Store) Get(_ context.Context, linkID string) (*rpc.Meta, error) {
	response, ok := s.metadata.Load(linkID)
	if !ok {
		return nil, &errors.MetadataNotFoundByIdError{ID: linkID}
	}

	v, ok := response.(*rpc.Meta)
	if !ok {
		return nil, &errors.MetadataNotFoundByIdError{ID: linkID}
	}

	return v, nil
}

// Add - write new metadata for a link
func (s *Store) Add(_ context.Context, source *rpc.Meta) error {
	s.metadata.Store(source.GetId(), source)

	return nil
}
