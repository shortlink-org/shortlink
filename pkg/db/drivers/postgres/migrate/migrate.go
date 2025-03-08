package migrate

import (
	"context"
	"embed"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/johejo/golang-migrate-extra/source/iofs"

	"github.com/shortlink-org/shortlink/pkg/db"
)

// Migration - apply migration to db
func Migration(_ context.Context, store db.DB, fs embed.FS, tableName string) error {
	client, ok := store.GetConn().(*pgxpool.Pool)
	if !ok {
		return db.ErrGetConnection
	}

	driverMigrations, err := iofs.New(fs, "migrations")
	if err != nil {
		return &MigrationError{
			Err:         err,
			Description: "failed to create migration source",
		}
	}

	conn := stdlib.OpenDBFromPool(client)

	driverDB, err := pgx.WithInstance(conn, &pgx.Config{
		MigrationsTable: "schema_migrations_" + strings.ReplaceAll(tableName, "-", "_"),
	})
	if err != nil {
		return &MigrationError{
			Err:         err,
			Description: "failed to create migration driver",
		}
	}

	migration, err := migrate.NewWithInstance("iofs", driverMigrations, "postgres", driverDB)
	if err != nil {
		return &MigrationError{
			Err:         err,
			Description: "failed to create migration instance",
		}
	}

	err = migration.Up()
	if err != nil && err.Error() != "no change" {
		return &MigrationError{
			Err:         err,
			Description: "failed to apply migration",
		}
	}

	return nil
}
