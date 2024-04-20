package leveldb

import (
	"context"
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

// Store implementation of db interface
type Store struct {
	client *leveldb.DB
}

// New store
func New(ctx context.Context, store db.DB, log logger.Logger) (*Store, error) {
	conn, ok := store.GetConn().(*leveldb.DB)
	if !ok {
		return nil, db.ErrGetConnection
	}

	s := &Store{
		client: conn,
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		err := s.client.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	return s, nil
}

// Add - add new link
func (l *Store) Add(_ context.Context, source *v1.Link) (*v1.Link, error) {
	payload, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	err = l.client.Put([]byte(source.GetHash()), payload, nil)
	if err != nil {
		return nil, err
	}

	return source, nil
}

// Get - a get link
func (l *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	value, err := l.client.Get([]byte(id), nil)
	if err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	var response v1.Link
	err = json.Unmarshal(value, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// List - list links
func (l *Store) List(_ context.Context, _ *types.FilterLink) (*v1.Links, error) {
	links := v1.NewLinks()
	iterator := l.client.NewIterator(nil, nil)

	for iterator.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		value := iterator.Value()

		var response v1.Link
		err := json.Unmarshal(value, &response)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}}
		}

		links.Push(&response)
	}

	iterator.Release()
	err := iterator.Error()
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}

	return links, nil
}

// Update - update link
func (l *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete - delete link
func (l *Store) Delete(ctx context.Context, id string) error {
	err := l.client.Delete([]byte(id), nil)
	return err
}
