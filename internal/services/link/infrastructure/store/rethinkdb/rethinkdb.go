// Deprecated: This database is no longer supported
package rethinkdb

import (
	"context"
	"fmt"

	"gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *rethinkdb.Session
}

type Link struct {
	*v1.Link
	Id string `gorethink:"id,omitempty"`
}

// Init ...
func (s *Store) Init(_ context.Context, db *db.Store) error {
	s.client = db.Store.GetConn().(*rethinkdb.Session)
	return nil
}

func (r *Store) Get(_ context.Context, id string) (*v1.Link, error) {
	c, err := rethinkdb.DB("shortlink").Table("link").Get(id).Run(r.client)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	if c.IsNil() {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var customlink Link
	err = c.One(&customlink)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return customlink.Link, nil
}

// TODO: How get all keys?
func (r *Store) List(_ context.Context, _ *query.Filter) (*v1.Links, error) {
	return &v1.Links{
		Link: make([]*v1.Link, 0),
	}, nil
}

func (r *Store) Add(_ context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	linkRethinkDB := &Link{
		source,
		source.Hash,
	}

	_, err = rethinkdb.DB("shortlink").Table("link").Insert(linkRethinkDB, rethinkdb.InsertOpts{}).RunWrite(r.client)
	if err != nil {
		return nil, err
	}

	return source, nil
}

func (s *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

func (r *Store) Delete(_ context.Context, id string) error {
	_, err := rethinkdb.DB("shortlink").Table("link").Get(id).Delete().Run(r.client)
	if err != nil {
		return err
	}

	return nil
}
