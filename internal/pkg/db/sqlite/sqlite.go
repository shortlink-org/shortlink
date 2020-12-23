package sqlite

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver
	"github.com/spf13/viper"
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

// Init ...
func (lite *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	lite.setConfig()

	if lite.client, err = sql.Open("sqlite3", lite.config.Path); err != nil {
		return err
	}

	lite.client.SetMaxOpenConns(25)
	lite.client.SetMaxIdleConns(2)
	lite.client.SetConnMaxLifetime(time.Minute)

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

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (lite *Store) Close() error {
	return lite.client.Close()
}

// Migrate ...
func (lite *Store) migrate() error { // nolint unused
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
