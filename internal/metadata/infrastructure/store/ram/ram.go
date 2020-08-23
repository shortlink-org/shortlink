package ram

import (
	"context"
	"fmt"
	"sync"

	rpc "github.com/batazor/shortlink/internal/metadata/domain"
	"github.com/batazor/shortlink/pkg/domain/link"
)

type Store struct {
	// sync.Map solver problem with cache contention
	metadata sync.Map
}

// Get - get metadata by id
func (s *Store) Get(_ context.Context, id string) (*rpc.Meta, error) {
	response, ok := s.metadata.Load(id)
	if !ok {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	v, ok := response.(*rpc.Meta)
	if !ok {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return v, nil
}

// Set - write new metadata for link
func (s *Store) Add(ctx context.Context, source *rpc.Meta) error {
	err := s.singleWrite(ctx, source)

	return err
}

func (s *Store) singleWrite(_ context.Context, source *rpc.Meta) error { // nolint unused
	s.metadata.Store(source.Id, source)

	return nil
}
