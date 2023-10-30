package migrate

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/johejo/golang-migrate-extra/source/iofs"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
)

// Migration - apply migration to db
func Migration(_ context.Context, store db.DB, fs embed.FS, tableName string) error {
	client, ok := store.GetConn().(*pgxpool.Pool)
	if !ok {
		return errors.New("can't get db connection")
	}

	driverMigrations, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	conn := stdlib.OpenDBFromPool(client)

	driverDB, err := postgres.WithInstance(conn, &postgres.Config{
		MigrationsTable: fmt.Sprintf("schema_migrations_%s", strings.ReplaceAll(tableName, "-", "_")),
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", driverMigrations, "postgres", driverDB)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
