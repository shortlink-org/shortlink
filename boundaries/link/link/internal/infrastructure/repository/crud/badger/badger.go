package badger

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/badger/v4"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
)

// Store implementation of db interface
type Store struct {
	client *badger.DB
}

// New store
func New(_ context.Context) (*Store, error) {
	return &Store{}, nil
}

// Get - get
func (b *Store) Get(_ context.Context, id string) (*domain.Link, error) {
	var valCopy []byte

	err := b.client.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			return nil
		})
		if err != nil {
			return err
		}

		// Alternatively, you could also use item.ValueCopy().
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	var response domain.Link
	err = json.Unmarshal(valCopy, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// List - list
func (b *Store) List(_ context.Context, _ *v1.FilterLink) (*domain.Links, error) {
	var list [][]byte

	err := b.client.View(func(txn *badger.Txn) error {
		iterator := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iterator.Close()

		for iterator.Rewind(); iterator.Valid(); iterator.Next() {
			var valCopy []byte
			item := iterator.Item()

			err := item.Value(func(val []byte) error {
				// Copying or parsing val is valid.
				valCopy = append([]byte{}, val...)

				return nil
			})
			if err != nil {
				return err
			}

			// Alternatively, you could also use item.ValueCopy().
			valCopy, err = item.ValueCopy(nil)
			if err != nil {
				return err
			}

			list = append(list, valCopy)
		}

		return nil
	})
	if err != nil {
		return nil, &domain.NotFoundError{Link: &domain.Link{}}
	}

	response := domain.NewLinks()

	for _, item := range list {
		l := &domain.Link{}
		err = json.Unmarshal(item, l)
		if err != nil {
			return nil, err
		}

		response.Push(l)
	}

	return response, nil
}

// Add - add
func (b *Store) Add(_ context.Context, source *domain.Link) (*domain.Link, error) {
	payload, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	err = b.client.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(source.GetHash()), payload)
	})
	if err != nil {
		return nil, err
	}

	return source, nil
}

// Update - update
func (b *Store) Update(_ context.Context, _ *domain.Link) (*domain.Link, error) {
	return nil, nil
}

// Delete - delete
func (b *Store) Delete(ctx context.Context, id string) error {
	err := b.client.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(id))
		return err
	})

	return err
}
