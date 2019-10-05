package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/batazor/shortlink/pkg/link"
	_ "github.com/lib/pq"
)

type PostgresLinkList struct {
	client *sql.DB
}

func (p *PostgresLinkList) Init() error {
	const (
		DB_USER     = "shortlink"
		DB_PASSWORD = "shortlink"
		DB_NAME     = "shortlink"
	)

	var err error

	// Connect to Postgres
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	p.client, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *PostgresLinkList) Get(id string) (*link.Link, error) {
	rows, err := p.client.Query("SELECT url, hash, describe FROM links LIMIT 1")

	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	var response link.Link

	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
		}
	}

	return &response, nil
}

func (p *PostgresLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.GetHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	err := p.client.QueryRow("INSERT INTO links(url,hash,describe) VALUES($1,$2,$3) ON CONFLICT (hash) DO NOTHING;", data.Url, data.Hash, data.Describe)

	if err.Scan().Error() == "sql: no rows in result set" {
		return &data, nil
	}
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: errors.New(fmt.Sprintf("Failed save link: %s", data.Url))}
	}

	return &data, nil
}

func (p *PostgresLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

func (p *PostgresLinkList) Delete(id string) error {
	stmt, err := p.client.Prepare("delete from links where hash=$1")
	if err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Failed save link: %s", id))}
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Failed save link: %s", id))}
	}

	return nil
}
