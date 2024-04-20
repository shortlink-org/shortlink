package redis

import (
	"context"
	"encoding/json"

	"github.com/redis/rueidis"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
)

// Store implementation of db interface
type Store struct {
	client rueidis.Client
}

// New store
func New(ctx context.Context, store db.DB) (*Store, error) {
	conn, ok := store.GetConn().(rueidis.Client)
	if !ok {
		return nil, db.ErrGetConnection
	}

	s := &Store{
		client: conn,
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		conn.Close()
	}()

	return s, nil
}

// Get - get
func (s *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	val, err := s.client.Do(ctx, s.client.B().Get().Key(id).Build()).ToString()
	if err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	var response v1.Link
	if err = json.Unmarshal([]byte(val), &response); err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	return &response, nil
}

// List - list
func (s *Store) List(ctx context.Context, _ *types.FilterLink) (*v1.Links, error) {
	list, err := s.client.Do(ctx, s.client.B().Scan().Cursor(0).Match("*").Count(100).Build()).AsScanEntry()
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}

	values, err := s.client.Do(ctx, s.client.B().Mget().Key(list.Elements...).Build()).ToArray()
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}

	links := v1.NewLinks()

	for _, item := range values {
		var response v1.Link

		value, errAsBytes := item.AsBytes()
		if errAsBytes != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}}
		}

		if err = json.Unmarshal(value, &response); err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}}
		}

		links.Push(&response)
	}

	return links, nil
}

// Add - add a new link
func (s *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	val, err := json.Marshal(source)
	if err != nil {
		return nil, &v1.NotFoundError{Link: source}
	}

	err = s.client.Do(ctx, s.client.B().Set().Key(source.GetHash()).Value(rueidis.BinaryString(val)).Build()).Error()
	if err != nil {
		return nil, &v1.NotFoundError{Link: source}
	}

	return source, nil
}

// Update - update
func (s *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete - delete
func (s *Store) Delete(ctx context.Context, id string) error {
	err := s.client.Do(ctx, s.client.B().Del().Key(id).Build()).Error()
	if err != nil {
		return &v1.NotFoundByHashError{Hash: id}
	}

	return nil
}
