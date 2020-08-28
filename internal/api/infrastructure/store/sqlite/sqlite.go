package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
)

// Config ...
type Config struct { // nolint unused
	Path string
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *sql.DB
	config Config
}

// Get ...
func (lite *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := lite.client.Prepare(query)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	defer stmt.Close() // nolint errcheck

	var response link.Link
	err = stmt.QueryRow(args...).Scan(&response.Url, &response.Hash, &response.Describe)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (lite *Store) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links")
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := lite.client.Query(query, args...)
	if err != nil || rows.Err() != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}
	defer rows.Close() // nolint errcheck

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
func (lite *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	// query builder
	links := squirrel.Insert("links").
		Columns("url", "hash", "describe").
		Values(data.Url, data.Hash, data.Describe)

	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = lite.client.Exec(query, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return data, nil
}

// Update ...
func (lite *Store) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (lite *Store) Delete(ctx context.Context, id string) error {
	// query builder
	links := squirrel.Delete("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = lite.client.Exec(query, args...)
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

// setConfig - set configuration
func (lite *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_SQLITE_PATH", "/tmp/links.sqlite") // SQLite URI

	lite.config = Config{
		Path: viper.GetString("STORE_SQLITE_PATH"),
	}
}
