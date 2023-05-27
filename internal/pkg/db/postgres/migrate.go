package postgres

import (
	"embed"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/johejo/golang-migrate-extra/source/iofs"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS
)

// migrate - apply migration to db
func (p *Store) migrate() error {
	uri := strings.Join([]string{viper.GetString("STORE_POSTGRES_URI"), "x-multi-statement=true"}, "&")

	// Create connect
	db, err := otelsql.Open("postgres", uri, otelsql.WithAttributes(semconv.DBSystemPostgreSQL), otelsql.WithDBName("PostgreSQL"))
	if err != nil {
		return err
	}

	driver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", driver, uri)
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
