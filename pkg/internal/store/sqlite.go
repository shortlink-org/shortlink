package store

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/batazor/shortlink/pkg/link"
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteLinkList implementation of store interface
type SQLiteLinkList struct { // nolint unused
	client *sql.DB
}

// Init ...
func (lite *SQLiteLinkList) Init() error {
	var err error

	if lite.client, err = sql.Open("sqlite3", "/tmp/links.sqlite"); err != nil {
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
