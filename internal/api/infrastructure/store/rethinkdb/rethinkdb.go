// Deprecated: This database is no longer supported
package rethinkdb

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
)

// Config ...
type Config struct { // nolint unused
	URI []string
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *rethinkdb.Session
	config Config
}

type Link struct {
	*link.Link
	Id string `gorethink:"id,omitempty"`
}

func (r *Store) Get(ctx context.Context, id string) (*link.Link, error) {
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
func (r *Store) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) {
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

func (r *Store) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

func (r *Store) Delete(ctx context.Context, id string) error {
	_, err := rethinkdb.DB("shortlink").Table("link").Get(id).Delete().Run(r.client)
	if err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (r *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_RETHINKDB_URI", "localhost:28015") // RethinkDB URI

	r.config = Config{
		URI: viper.GetStringSlice("STORE_RETHINKDB_URI"),
	}
}
