package store

import (
	"github.com/batazor/shortlink/pkg/link"
	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

// CassandraConfig ...
type CassandraConfig struct { // nolint unused
	URI string
}

// CassandraLinkList implementation of store interface
type CassandraLinkList struct { // nolint unused
	client *gocql.Session
	config CassandraConfig
}

// Init ...
func (c *CassandraLinkList) Init() error {
	var err error

	// Set configuration
	c.setConfig()

	// Connect to CassandraDB
	cluster := gocql.NewCluster(c.config.URI)
	cluster.Keyspace = "shortlink"

	c.client, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	return nil
}

// Close ...
func (c *CassandraLinkList) Close() error {
	c.client.Close()
	return nil
}

// Get ...
func (c *CassandraLinkList) Get(id string) (*link.Link, error) {
	panic("implement me")

	return nil, nil
}

// Add ...
func (c *CassandraLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	if err := c.client.Query(`INSERT INTO links (url hash describe) VALUES (?, ?, ?)`, data.Url, data.Hash, data.Describe).Exec(); err != nil {
		return nil, err
	}

	return &data, nil
}

// List ...
func (c *CassandraLinkList) List() ([]*link.Link, error) {
	iter := c.client.Query(`SELECT url, hash, describe FROM links`).Iter()

	// Here's an array in which you can store the decoded documents
	var response []*link.Link
	var link link.Link

	for iter.Scan(&link) {
		response = append(response, &link)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return response, nil
}

// Update ...
func (c *CassandraLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (c *CassandraLinkList) Delete(id string) error {
	panic("implement me")

	return nil
}

// setConfig - set configuration
func (c *CassandraLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_CASSANDRA_URI", "localhost")
	c.config = CassandraConfig{
		URI: viper.GetString("STORE_CASSANDRA_URI"),
	}
}
