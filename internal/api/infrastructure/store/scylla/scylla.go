package scylla

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/db"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client gocqlx.Session

	linksTable *table.Table
}

// Init ...
func (s *Store) Init(_ context.Context, db *db.Store) error {
	s.client = db.Store.GetConn().(gocqlx.Session)

	m := table.Metadata{
		Name:    "shortlink.links",
		Columns: []string{"url", "hash", "ddd"},
	}
	s.linksTable = table.New(m)

	return nil
}

// Get ...
func (s *Store) Get(_ context.Context, id string) (*link.Link, error) {
	stmt, values := qb.Select("shortlink.links").
		Columns(s.linksTable.Metadata().Columns...).
		Where(qb.EqNamed("hash", id)).
		ToCql()
	iter, err := s.client.Query(stmt, values).Bind(id).Consistency(gocql.One).Iter().SliceMap()
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
func (s *Store) List(_ context.Context, _ *query.Filter) ([]*link.Link, error) {
	stmt, values := qb.Select("shortlink.links").
		Columns(s.linksTable.Metadata().Columns...).
		ToCql()
	iter, err := s.client.Query(stmt, values).Iter().SliceMap()
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
func (s *Store) Add(_ context.Context, source *link.Link) (*link.Link, error) {
	err := link.NewURL(source)
	if err != nil {
		return nil, err
	}

	if err := s.client.Query(s.linksTable.Insert()).BindMap(map[string]interface{}{
		"url":  source.Url,
		"hash": source.Hash,
		"ddd":  source.Describe,
	}).Exec(); err != nil {
		return nil, err
	}

	return source, nil
}

// Update ...
func (s *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (s *Store) Delete(ctx context.Context, id string) error {
	stmt, values := s.linksTable.DeleteBuilder("url", "ddd").
		Where(qb.EqNamed("hash", id)).
		ToCql()

	err := s.client.Query(stmt, values).Bind(id).Exec()
	return err
}
