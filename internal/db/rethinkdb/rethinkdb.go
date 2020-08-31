// Deprecated: This database is no longer supported
package rethinkdb

import (
	"context"

	"github.com/spf13/viper"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/tool"
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

func (r *Store) Init(ctx context.Context) error {
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

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Migrate ...
func (r *Store) migrate() error { // nolint unused
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

func (r *Store) Close() error {
	err := r.client.Close()
	return err
}

// setConfig - set configuration
func (r *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_RETHINKDB_URI", "localhost:28015") // RethinkDB URI

	r.config = Config{
		URI: viper.GetStringSlice("STORE_RETHINKDB_URI"),
	}
}
