//go:generate protoc -I../../../internal/api/domain/link --gotemplate_out=all=true,template_dir=template:. link.proto
//go:generate go-bindata -prefix migrations -pkg migrations -ignore migrations.go -o migrations/migrations.go migrations
package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq" // need for init PostgreSQL interface

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/db/options"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

// Get ...
func (p *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)

	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	if rows.Err() != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link
	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
		}
	}

	if response.Hash == "" {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (p *Store) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("links")
	links = p.buildFilter(links, filter)
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	var response []*link.Link

	for rows.Next() {
		var result link.Link
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		response = append(response, &result)
	}

	return response, nil
}

// Add ...
func (p *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	switch p.config.mode {
	case options.MODE_BATCH_WRITE:
		cb, err := p.config.job.Push(source)
		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, err
		case link.Link:
			return &data, nil
		default:
			return nil, nil
		}
	case options.MODE_SINGLE_WRITE:
		data, err := p.singleWrite(ctx, source)
		return data, err
	}

	return nil, nil
}

// Update ...
func (p *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (p *Store) Delete(ctx context.Context, id string) error {
	// query builder
	links := psql.Delete("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = p.client.Exec(ctx, q, args...)
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

func (p *Store) singleWrite(ctx context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	// save as JSON. it doesn't make sense
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// query builder
	links := psql.Insert("links").
		Columns("url", "hash", "describe", "json").
		Values(data.Url, data.Hash, data.Describe, dataJson)

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&data.Url, &data.Hash, &data.Describe).Error()
	if errScan == "no rows in result set" {
		return data, nil
	}
	if errScan != "" {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return data, nil
}
