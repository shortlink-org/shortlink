package sqlite

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// Config ...
type Config struct {
	Path string
}

// Store implementation of db interface
type Store struct {
	client *sql.DB
	config Config
}

// Init ...
func (s *Store) Init(ctx context.Context) error {
	const SET_MAX_OPEN_CONNS = 25
	const SET_MAX_IDLE_CONNS = 2

	var err error

	// Set configuration
	s.setConfig()

	s.client, err = otelsql.Open("sqlite3", s.config.Path, otelsql.WithAttributes(semconv.DBSystemSqlite), otelsql.WithDBName("SQLite"))
	if err != nil {
		return err
	}

	s.client.SetMaxOpenConns(SET_MAX_OPEN_CONNS)
	s.client.SetMaxIdleConns(SET_MAX_IDLE_CONNS)
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

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_SQLITE_PATH", "/tmp/links.sqlite") // SQLite URI

	s.config = Config{
		Path: viper.GetString("STORE_SQLITE_PATH"),
	}
}
