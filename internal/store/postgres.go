package store

import (
	"context"
	"fmt"
	"github.com/spf13/viper"

	"github.com/Masterminds/squirrel"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

// PostgresConfig ...
type PostgresConfig struct { // nolint unused
	URI string
}

// PostgresLinkList implementation of store interface
type PostgresLinkList struct { // nolint unused
	client *pgxpool.Pool
	config PostgresConfig
}

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
func (p *PostgresLinkList) Close() error {
	p.client.Close()
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
func (p *PostgresLinkList) List() ([]*link.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("links")
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
func (p *PostgresLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	// query builder
	links := psql.Insert("links").
		Columns("url", "hash", "describe").
		Values(data.Url, data.Hash, data.Describe)

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
	dbinfo := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", "postgres", "postgres", "shortlink")

	viper.AutomaticEnv()
	viper.SetDefault("STORE_POSTGRES_URI", dbinfo)

	p.config = PostgresConfig{
		URI: viper.GetString("STORE_POSTGRES_URI"),
	}
}
