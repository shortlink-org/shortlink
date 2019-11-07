package store

import (
	"database/sql"
	"fmt"

	"github.com/batazor/shortlink/pkg/link"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
)

// PostgresLinkList implementation of store interface
type PostgresLinkList struct { // nolint unused
	client *sql.DB
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
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)
	p.client, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	return nil
}

// Get ...
func (p *PostgresLinkList) Get(id string) (*link.Link, error) {
	rows, err := p.client.Query("SELECT url, hash, describe FROM links LIMIT 1")

	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	for rows.Next() {
		err = rows.Scan(&response.URL, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
		}
	}

	return &response, nil
}

// Add ...
func (p *PostgresLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.URL), []byte("secret"))
	data.Hash = hash[:7]

	err := p.client.QueryRow("INSERT INTO links(url,hash,describe) VALUES($1,$2,$3) ON CONFLICT (hash) DO NOTHING;", data.URL, data.Hash, data.Describe)

	if err.Scan().Error() == "sql: no rows in result set" {
		return &data, nil
	}
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.URL)}
	}

	return &data, nil
}

// List ...
func (b *PostgresLinkList) List() ([]*link.Link, error) {
	panic("implement me")
}

// Update ...
func (p *PostgresLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (p *PostgresLinkList) Delete(id string) error {
	stmt, err := p.client.Prepare("delete from links where hash=$1")
	if err != nil {
		return &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}
