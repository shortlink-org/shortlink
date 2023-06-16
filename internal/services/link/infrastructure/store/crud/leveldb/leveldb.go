package leveldb

import (
	"context"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

// Store implementation of db interface
type Store struct {
	client *leveldb.DB
}

// New store
func New(ctx context.Context, db *db.Store) (*Store, error) {
	s := &Store{
		client: db.Store.GetConn().(*leveldb.DB),
	}

	return s, nil
}

// Add ...
func (l *Store) Add(_ context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	payload, err := protojson.Marshal(source)
	if err != nil {
		return nil, err
	}

	err = l.client.Put([]byte(source.Hash), payload, nil)
	if err != nil {
		return nil, err
	}

	return source, nil
}

// Get ...
func (l *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	value, err := l.client.Get([]byte(id), nil)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response v1.Link
	err = protojson.Unmarshal(value, &response)
	if err != nil {
		return nil, err
	}

	if response.Url == "" {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (l *Store) List(_ context.Context, _ *query.Filter) (*v1.Links, error) {
	links := &v1.Links{
		Link: []*v1.Link{},
	}
	iterator := l.client.NewIterator(nil, nil)

	for iterator.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		value := iterator.Value()

		var response v1.Link
		err := protojson.Unmarshal(value, &response)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: query.ErrNotFound}
		}

		links.Link = append(links.Link, &response)
	}

	iterator.Release()
	err := iterator.Error()
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: query.ErrNotFound}
	}

	return links, nil
}

// Update ...
func (l *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete ...
func (l *Store) Delete(ctx context.Context, id string) error {
	err := l.client.Delete([]byte(id), nil)
	return err
}
