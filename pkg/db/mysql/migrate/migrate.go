package migrate

import (
	"context"
	"database/sql"
	"embed"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/johejo/golang-migrate-extra/source/iofs"

	"github.com/shortlink-org/shortlink/pkg/db"
)

// Migration - apply migration to db
func Migration(_ context.Context, store db.DB, fs embed.FS, tableName string) error {
	client, ok := store.GetConn().(*sql.DB)
	if !ok {
		return db.ErrGetConnection
	}

	driverMigrations, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	driverDB, err := mysql.WithInstance(client, &mysql.Config{
		MigrationsTable: "schema_migrations_" + strings.ReplaceAll(tableName, "-", "_"),
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", driverMigrations, "mysql", driverDB)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
