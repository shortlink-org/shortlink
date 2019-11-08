package store

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/pkg/link"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
)

// PostgresLinkList implementation of store interface
type PostgresLinkList struct { // nolint unused
	client *pgx.Conn
}

// Init ...
func (p *PostgresLinkList) Init() error {
	const (
		DbUser     = "shortlink"
		DbPassword = "shortlink"
		DbName     = "shortlink"
	)

	var err error

	// Connect to Postgres
	dbinfo := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", DbUser, DbPassword, DbName)
	if p.client, err = pgx.Connect(context.Background(), dbinfo); err != nil {
		panic(err)
	}

	return nil
}

// Get ...
func (p *PostgresLinkList) Get(id string) (*link.Link, error) {
	rows, err := p.client.Query(context.Background(), "SELECT url, hash, describe FROM links WHERE hash=$1", id)

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

// Add ...
func (p *PostgresLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	err := p.client.QueryRow(context.Background(), "INSERT INTO links(url,hash,describe) VALUES($1,$2,$3) ON CONFLICT (hash) DO NOTHING;", data.Url, data.Hash, data.Describe)

	if err.Scan().Error() == "sql: no rows in result set" {
		return &data, nil
	}
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return &data, nil
}

// List ...
func (p *PostgresLinkList) List() ([]*link.Link, error) {
	rows, err := p.client.Query(context.Background(), "SELECT url, hash, describe describe FROM links")

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

// Update ...
func (p *PostgresLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (p *PostgresLinkList) Delete(id string) error {
	_, err := p.client.Exec(context.Background(), "", "delete from links where hash=$1", id)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}
