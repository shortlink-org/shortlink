package cassandra

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"

	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes"
	"github.com/scylladb/gocqlx/qb"
	"github.com/spf13/viper"
)

// Config ...
type Config struct { // nolint unused
	URI string
}

// CassandraLinkList implementation of db interface
type CassandraLinkList struct { // nolint unused
	client *gocql.Session
	config Config
}

// Init ...
func (c *CassandraLinkList) Init(ctx context.Context) error {
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

// Close ...
func (c *CassandraLinkList) Close() error { // nolint unparam
	c.client.Close()
	return nil
}

// Migrate ...
// TODO: ddd -> describe
func (c *CassandraLinkList) migrate() error { // nolint unused
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

// Get ...
func (c *CassandraLinkList) Get(ctx context.Context, id string) (*link.Link, error) {
	stmt, values := qb.Select("shortlink.links").Columns("url", "hash", "ddd").Where(qb.EqNamed("hash", id)).ToCql()
	iter, err := c.client.Query(stmt, values[0]).Consistency(gocql.One).Iter().SliceMap()
	if err != nil {
		return nil, err
	}

	if len(iter) == 0 {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	// Here's an array in which you can db the decoded documents
	response := &link.Link{
		Url:      iter[0]["url"].(string),
		Hash:     iter[0]["hash"].(string),
		Describe: iter[0]["ddd"].(string),
	}

	return response, nil
}

// List ...
func (c *CassandraLinkList) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
	iter, err := c.client.Query(`SELECT url, hash, ddd FROM shortlink.links`).Iter().SliceMap()
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can db the decoded documents
	var response []*link.Link

	for index := range iter {
		response = append(response, &link.Link{
			Url:      iter[index]["url"].(string),
			Hash:     iter[index]["hash"].(string),
			Describe: iter[index]["ddd"].(string),
		})
	}

	return response, nil
}

// Add ...
func (c *CassandraLinkList) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	// Add timestamp
	data.CreatedAt = ptypes.TimestampNow()
	data.UpdatedAt = ptypes.TimestampNow()

	if err := c.client.Query(`INSERT INTO shortlink.links (url, hash, ddd) VALUES (?, ?, ?)`, data.Url, data.Hash, data.Describe).Exec(); err != nil {
		return nil, err
	}

	return data, nil
}

// Update ...
func (c *CassandraLinkList) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (c *CassandraLinkList) Delete(ctx context.Context, id string) error {
	err := c.client.Query(`DELETE FROM shortlink.links WHERE hash = ?`, id).Exec()
	return err
}

// setConfig - set configuration
func (c *CassandraLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_CASSANDRA_URI", "localhost:9042") // Cassandra URI
	c.config = Config{
		URI: viper.GetString("STORE_CASSANDRA_URI"),
	}
}
