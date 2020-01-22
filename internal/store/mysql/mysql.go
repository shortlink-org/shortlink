package mysql

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// Init ...
func (m *MySQLLinkList) Init() error {
	var err error

	// Set configuration
	m.setConfig()

	if m.client, err = sqlx.Connect("mysql", m.config.URI); err != nil {
		return err
	}

	sqlStmt := `
		CREATE TABLE IF NOT EXISTS links (
			id          int NOT NULL AUTO_INCREMENT,
			url         varchar(255) NOT NULL,
			hash        varchar(255) NOT NULL,
			description text NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
	`

	if _, err = m.client.Exec(sqlStmt); err != nil {
		panic(err)
	}

	return nil
}

// Close ...
func (m *MySQLLinkList) Close() error {
	return m.client.Close()
}

// Migrate ...
func (m *MySQLLinkList) migrate() error { // nolint unused
	return nil
}

// Get ...
func (m *MySQLLinkList) Get(id string) (*link.Link, error) {
	// query builder
	links := squirrel.Select("url, hash, description").
		From("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := m.client.Prepare(query)
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
func (m *MySQLLinkList) List(filter *query.Filter) ([]*link.Link, error) { // nolint unused
	// query builder
	links := squirrel.Select("url, hash, description").
		From("links")
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := m.client.Query(query, args...)
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
func (m *MySQLLinkList) Add(source link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	// query builder
	links := squirrel.Insert("links").
		Columns("url", "hash", "description").
		Values(data.Url, data.Hash, data.Describe)

	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = m.client.Exec(query, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return &data, nil
}

// Update ...
func (m *MySQLLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (m *MySQLLinkList) Delete(id string) error {
	// query builder
	links := squirrel.Delete("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = m.client.Exec(query, args...)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

// setConfig - set configuration
func (m *MySQLLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MYSQL_URI", "shortlink:shortlink@(localhost:3306)/shortlink?parseTime=true")

	m.config = MySQLConfig{
		URI: viper.GetString("STORE_MYSQL_URI"),
	}
}
