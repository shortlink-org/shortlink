//go:generate protoc -I../../../pkg/link --gotemplate_out=all=true,template_dir=template:. link.proto
package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

// Init ...
func (p *PostgresLinkList) Init() error {
	var err error

	// Set configuration
	p.setConfig()

	// Connect to Postgres
	if p.client, err = pgxpool.Connect(context.Background(), p.config.URI); err != nil {
		return err
	}

	return nil
}

// Close ...
func (p *PostgresLinkList) Close() error { // nolint unparam
	p.client.Close()
	return nil
}

// Migrate ...
func (p *PostgresLinkList) migrate() error { // nolint unused
	return nil
}

// Get ...
func (p *PostgresLinkList) Get(id string) (*link.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(context.Background(), query, args...)

	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
		}
	}

	return &response, nil
}

// List ...
func (p *PostgresLinkList) List(filter *query.Filter) ([]*link.Link, error) { // nolint unused
	// query builder
	links := psql.Select("url, hash, describe").
		From("links")
	links = p.buildFilter(links, filter)
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(context.Background(), query, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	var response []*link.Link

	for rows.Next() {
		var result link.Link
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		response = append(response, &result)
	}

	return response, nil
}

// Add ...
func (p *PostgresLinkList) Add(source link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	// save as JSON. it doesn't make sense
	dataJson, err := json.Marshal(data)

	// query builder
	links := psql.Insert("links").
		Columns("url", "hash", "describe", "json").
		Values(data.Url, data.Hash, data.Describe, dataJson)

	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(context.Background(), query, args...)

	errScan := row.Scan(&data.Url, &data.Hash, &data.Describe).Error()
	if errScan == "no rows in result set" {
		return &data, nil
	}
	if errScan != "" {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return &data, nil
}

// Update ...
func (p *PostgresLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (p *PostgresLinkList) Delete(id string) error {
	// query builder
	links := psql.Delete("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = p.client.Exec(context.Background(), query, args...)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

// setConfig - set configuration
func (p *PostgresLinkList) setConfig() {
	dbinfo := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", "shortlink", "shortlink", "shortlink")

	viper.AutomaticEnv()
	viper.SetDefault("STORE_POSTGRES_URI", dbinfo)

	p.config = PostgresConfig{
		URI: viper.GetString("STORE_POSTGRES_URI"),
	}
}
