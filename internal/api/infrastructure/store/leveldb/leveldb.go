package leveldb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *leveldb.DB
}

// Add ...
func (l *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = l.client.Put([]byte(data.Hash), payload, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Get ...
func (l *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	value, err := l.client.Get([]byte(id), nil)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	err = json.Unmarshal(value, &response)
	if err != nil {
		return nil, err
	}

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (l *Store) List(_ context.Context, _ *query.Filter) ([]*link.Link, error) {
	links := []*link.Link{}
	iterator := l.client.NewIterator(nil, nil)

	for iterator.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		value := iterator.Value()

		var response link.Link

		err := json.Unmarshal(value, &response)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		links = append(links, &response)
	}

	iterator.Release()
	err := iterator.Error()
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	return links, nil
}

// Update ...
func (l *Store) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (l *Store) Delete(ctx context.Context, id string) error {
	err := l.client.Delete([]byte(id), nil)
	return err
}
