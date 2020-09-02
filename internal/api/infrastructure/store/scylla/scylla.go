package scylla

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes"
	"github.com/scylladb/gocqlx/qb"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *gocql.Session
}

// Get ...
func (c *Store) Get(ctx context.Context, id string) (*link.Link, error) {
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
func (c *Store) List(_ context.Context, _ *query.Filter) ([]*link.Link, error) {
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
func (c *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
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
func (c *Store) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (c *Store) Delete(ctx context.Context, id string) error {
	err := c.client.Query(`DELETE FROM shortlink.links WHERE hash = ?`, id).Exec()
	return err
}
