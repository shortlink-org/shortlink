package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/luna-duclos/instrumentedsql"
	"github.com/luna-duclos/instrumentedsql/opentracing"
	"github.com/mattn/go-sqlite3"
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
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	sql.Register("instrumented-sqlite", instrumentedsql.WrapDriver(&sqlite3.SQLiteDriver{}, instrumentedsql.WithTracer(opentracing.NewTracer(false))))
	if s.client, err = sql.Open("instrumented-sqlite", s.config.Path); err != nil {
		return err
	}

	s.client.SetMaxOpenConns(25)
	s.client.SetMaxIdleConns(2)
	s.client.SetConnMaxLifetime(time.Minute)

	sqlStmt := `
		CREATE TABLE IF NOT EXISTS links (
			id integer not null primary key,
			url      varchar(255) not null,
			hash     varchar(255) not null,
			describe text
		);
	`

	if _, err = s.client.Exec(sqlStmt); err != nil {
		panic(err)
	}

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (s *Store) Close() error {
	return s.client.Close()
}

// Migrate ...
func (s *Store) migrate() error { // nolint unused
	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_SQLITE_PATH", "/tmp/links.sqlite") // SQLite URI

	s.config = Config{
		Path: viper.GetString("STORE_SQLITE_PATH"),
	}
}
