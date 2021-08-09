package ram

import (
	"context"
	"fmt"
	"sync"

	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	rpc "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"
)

type Store struct {
	// sync.Map solver problem with cache contention
	metadata sync.Map
}

// Get - get metadata by id
func (s *Store) Get(_ context.Context, id string) (*rpc.Meta, error) {
	response, ok := s.metadata.Load(id)
	if !ok {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	v, ok := response.(*rpc.Meta)
	if !ok {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return v, nil
}

// Set - write new metadata for link
func (s *Store) Add(_ context.Context, source *rpc.Meta) error {
	s.metadata.Store(source.Id, source)

	return nil
}
