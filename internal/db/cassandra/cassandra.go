package cassandra

import (
	"context"
	"net/url"
	"strconv"

	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

// Config ...
type Config struct { // nolint unused
	URI string
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *gocql.Session
	config Config
}

// Init ...
func (c *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	c.setConfig()

	uri, err := url.ParseRequestURI(c.config.URI)
	if err != nil {
		return err
	}

	// Connect to CassandraDB
	cluster := gocql.NewCluster(c.config.URI)
	cluster.ProtoVersion = 4
	cluster.Port, err = strconv.Atoi(uri.Opaque)

	if err != nil {
		return err
	}

	c.client, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	// Migration
	if err = c.migrate(); err != nil {
		panic(err)
	}

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (c *Store) Close() error { // nolint unparam
	c.client.Close()
	return nil
}

// Migrate ...
// TODO: ddd -> describe
func (c *Store) migrate() error { // nolint unused
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
		if err := c.client.Query(schema).Exec(); err != nil {
			return err
		}
	}

	return nil
}

// setConfig - set configuration
func (c *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_CASSANDRA_URI", "localhost:9042") // Cassandra URI
	c.config = Config{
		URI: viper.GetString("STORE_CASSANDRA_URI"),
	}
}
