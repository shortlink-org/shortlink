package scylla

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db/scylla/migration"
)

// Config ...
type Config struct { // nolint unused
	URI string
}

// Store implementation of db interface
type Store struct { // nolint unused
	client gocqlx.Session
	config Config
}

// Init ...
func (s *Store) Init(_ context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	// Create gocql cluster.
	cluster := gocql.NewCluster(s.config.URI)

	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	s.client, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return err
	}

	// Migration
	if err = s.migrate(); err != nil {
		return err
	}

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (s *Store) Close() error { // nolint unparam
	s.client.Close()
	return nil
}

// migrate ...
// TODO: ddd -> describe
func (s *Store) migrate() error {
	err := migrate.FromFS(context.Background(), s.client, migration.Files)
	if err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_SCYLLA_URI", "localhost:9042") // Scylla URI
	s.config = Config{
		URI: viper.GetString("STORE_SCYLLA_URI"),
	}
}
