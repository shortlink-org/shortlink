//go:generate go-bindata -prefix migrations -pkg migrations -ignore migrations.go -o migrations/migrations.go migrations

package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/db/mysql/migrations"
)

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
	if err != nil {
		return err
	}

	// wrap assets into Resource
	res := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})

	driver, err := bindata.WithInstance(res)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", driver, fmt.Sprintf("mysql://%s", s.config.URI))
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
