package migrate

import (
	"context"
	"embed"
	"errors"
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

	uri := strings.Builder{}
	uri.WriteString(client.Config().ConnString())
	uri.WriteString("&x-migrations-table=")
	uri.WriteString(tableName)

	m, err := migrate.NewWithSourceInstance("iofs", driver, uri.String())
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
