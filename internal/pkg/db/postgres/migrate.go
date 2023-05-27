package postgres

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/johejo/golang-migrate-extra/source/iofs"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS
)

// migrate - apply migration to db
func (p *Store) migrate() error {
	driver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", driver, p.config.config.ConnString())
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
