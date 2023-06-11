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
func Migration(ctx context.Context, db *db.Store, fs embed.FS, tableName string) error {
	client, ok := db.Store.GetConn().(*pgxpool.Pool)
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

	uri.WriteString(connStr)
	uri.WriteString("&x-migrations-table=")
	uri.WriteString(fmt.Sprintf("schema_migrations_%s", strings.Replace(tableName, "-", "_", -1)))
	return uri.String()
}
