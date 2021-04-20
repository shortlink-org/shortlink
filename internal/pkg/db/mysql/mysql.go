package mysql

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/johejo/golang-migrate-extra/source/file"
	"github.com/johejo/golang-migrate-extra/source/iofs"
	"github.com/spf13/viper"
)

//go:embed migrations/*.sql
var migrations embed.FS

// Init ...
func (s *Store) Init(_ context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	if s.client, err = sqlx.Connect("mysql", s.config.URI); err != nil {
		return err
	}

	// Apply migration
	err = s.migrate()
	if err != nil {
		return err
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
	// Create connect
	db, err := sql.Open("mysql", s.config.URI)

	driver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", driver, fmt.Sprintf("mysql://%s", s.config.URI))
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MYSQL_URI", "shortlink:shortlink@(localhost:3306)/shortlink?parseTime=true") // MySQL URI

	s.config = Config{
		URI: viper.GetString("STORE_MYSQL_URI"),
	}
}
