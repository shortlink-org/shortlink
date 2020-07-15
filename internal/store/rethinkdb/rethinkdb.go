package rethinkdb

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/internal/tool"
	"github.com/batazor/shortlink/pkg/link"
)

// RethinkDBConfig ...
type RethinkDBConfig struct { // nolint unused
	URI []string
}

// MongoLinkList implementation of store interface
type RethinkDBLinkList struct { // nolint unused
	client *rethinkdb.Session
	config RethinkDBConfig
}

type Link struct {
	*link.Link
	Id string `gorethink:"id,omitempty"`
}

func (r *RethinkDBLinkList) Init(ctx context.Context) error {
	var err error

	// Set configuration
	r.setConfig()

	// Connect to RethinkDB
	r.client, err = rethinkdb.Connect(rethinkdb.ConnectOpts{
		Addresses:  r.config.URI, // endpoint without http
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		return err
	}

	// Apply migration
	if err = r.migrate(); err != nil {
		return err
	}

	return nil
}

// Migrate ...
func (r *RethinkDBLinkList) migrate() error { // nolint unused
	// create database
	dbList, err := r.getDatabases()
	if err != nil {
		return err
	}

	if ok := tool.Contains(dbList, "shortlink"); !ok {
		if _, errDBCreate := rethinkdb.DBCreate("shortlink").Run(r.client); errDBCreate != nil {
			return errDBCreate
		}
	}

	// Create table
	tableList, err := r.getTables()
	if err != nil {
		return err
	}

	if ok := tool.Contains(tableList, "link"); !ok {
		if _, errTableCreate := rethinkdb.DB("shortlink").TableCreate("link").Run(r.client); errTableCreate != nil {
			return errTableCreate
		}
	}

	return nil
}

func (r *RethinkDBLinkList) Get(ctx context.Context, id string) (*link.Link, error) {
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
func (r *RethinkDBLinkList) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) {
	return nil, nil
}

func (r *RethinkDBLinkList) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
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

func (r *RethinkDBLinkList) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

func (r *RethinkDBLinkList) Delete(ctx context.Context, id string) error {
	_, err := rethinkdb.DB("shortlink").Table("link").Get(id).Delete().Run(r.client)
	if err != nil {
		return err
	}

	return nil
}

func (r *RethinkDBLinkList) Close() error {
	err := r.client.Close()
	return err
}

// setConfig - set configuration
func (r *RethinkDBLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_RETHINKDB_URI", "localhost:28015") // MongoDB URI

	r.config = RethinkDBConfig{
		URI: viper.GetStringSlice("STORE_RETHINKDB_URI"),
	}
}
