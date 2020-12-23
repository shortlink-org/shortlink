package scylla

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/spf13/viper"
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

// Migrate ...
// TODO: ddd -> describe
func (s *Store) migrate() error { // nolint unused
	infoSchemas := []string{`
CREATE KEYSPACE IF NOT EXISTS shortlink
	WITH REPLICATION = {
		'class' : 'SimpleStrategy',
		'replication_factor': 1
	};`, `
CREATE TABLE IF NOT EXISTS shortlink.links (
	url text,
	hash text,
	ddd text,
	PRIMARY KEY(hash)
)`}

	for _, schema := range infoSchemas {
		if err := s.client.ExecStmt(schema); err != nil {
			return err
		}
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
