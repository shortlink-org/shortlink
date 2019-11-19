package store

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/batazor/shortlink/pkg/link"
	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver
	"github.com/spf13/viper"
)

// RedisConfig ...
type SQLiteConfig struct { // nolint unused
	Path string
}

// SQLiteLinkList implementation of store interface
type SQLiteLinkList struct { // nolint unused
	client *sql.DB
	config SQLiteConfig
}

// Init ...
func (lite *SQLiteLinkList) Init() error {
	var err error

	// Set configuration
	lite.setConfig()

	if lite.client, err = sql.Open("sqlite3", lite.config.Path); err != nil {
		return err
	}

	sqlStmt := `
		CREATE TABLE IF NOT EXISTS links (
			id integer not null primary key,
			url      varchar(255) not null,
			hash     varchar(255) not null,
			describe text
		);
	`

	if _, err = lite.client.Exec(sqlStmt); err != nil {
		panic(err)
	}

	return nil
}

// Close ...
func (lite *SQLiteLinkList) Close() error {
	return lite.client.Close()
}

// Get ...
func (lite *SQLiteLinkList) Get(id string) (*link.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := lite.client.Prepare(query)
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	defer stmt.Close() // nolint errcheck

	var response link.Link
	err = stmt.QueryRow(args...).Scan(&response.Url, &response.Hash, &response.Describe)
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (lite *SQLiteLinkList) List() ([]*link.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("links")
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := lite.client.Query(query, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{}, Err: fmt.Errorf("Not found links")}
	}
	defer rows.Close() // nolint errcheck

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
func (lite *SQLiteLinkList) Add(data link.Link) (*link.Link, error) {
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

	_, err = lite.client.Exec(query, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return &data, nil
}

// Update ...
func (lite *SQLiteLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (lite *SQLiteLinkList) Delete(id string) error {
	// query builder
	links := psql.Delete("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = lite.client.Exec(query, args...)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

// setConfig - set configuration
func (lite *SQLiteLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_SQLITE_PATH", "/tmp/links.sqlite")

	lite.config = SQLiteConfig{
		Path: viper.GetString("STORE_SQLITE_PATH"),
	}
}
