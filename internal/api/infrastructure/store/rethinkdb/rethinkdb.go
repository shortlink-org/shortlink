// Deprecated: This database is no longer supported
package rethinkdb

import (
	"context"
	"fmt"

	"gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/db"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *rethinkdb.Session
}

type Link struct {
	*link.Link
	Id string `gorethink:"id,omitempty"`
}

// Init ...
func (_ *Store) Init(_ context.Context, _ *db.Store) error {
	return nil
}

func (r *Store) Get(_ context.Context, id string) (*link.Link, error) {
	c, err := rethinkdb.DB("shortlink").Table("link").Get(id).Run(r.client)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	if c.IsNil() {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var customlink Link
	err = c.One(&customlink)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return customlink.Link, nil
}

// TODO: How get all keys?
func (r *Store) List(_ context.Context, _ *query.Filter) ([]*link.Link, error) {
	return nil, nil
}

func (r *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	linkRethinkDB := &Link{
		data,
		data.Hash,
	}

	_, err = rethinkdb.DB("shortlink").Table("link").Insert(linkRethinkDB, rethinkdb.InsertOpts{}).RunWrite(r.client)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

func (r *Store) Delete(_ context.Context, id string) error {
	_, err := rethinkdb.DB("shortlink").Table("link").Get(id).Delete().Run(r.client)
	if err != nil {
		return err
	}

	return nil
}
