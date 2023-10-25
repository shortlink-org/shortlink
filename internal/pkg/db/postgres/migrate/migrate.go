package migrate

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/johejo/golang-migrate-extra/source/iofs"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
)

// Migration - apply migration to db
func Migration(_ context.Context, store *db.Store, fs embed.FS, tableName string) error {
	client, ok := store.Store.GetConn().(*pgxpool.Pool)
	if !ok {
		return errors.New("can't get db connection")
	}

	driver, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	uri := buildURI(client, tableName)

	m, err := migrate.NewWithSourceInstance("iofs", driver, uri)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}

func buildURI(client *pgxpool.Pool, tableName string) string {
	uri := strings.Builder{}
	connStr := client.Config().ConnString()
	if !strings.Contains(connStr, "?") {
		connStr += "?"
	}

	_, _ = uri.WriteString(connStr)                                                                      //nolint:errcheck // ignore error
	_, _ = uri.WriteString("&dbname=")                                                                   //nolint:errcheck // ignore error
	_, _ = uri.WriteString(client.Config().ConnConfig.Database)                                          //nolint:errcheck // ignore error
	_, _ = uri.WriteString("&x-migrations-table=")                                                       //nolint:errcheck // ignore error
	_, _ = uri.WriteString(fmt.Sprintf("schema_migrations_%s", strings.ReplaceAll(tableName, "-", "_"))) //nolint:errcheck // ignore error

	return uri.String()
}
