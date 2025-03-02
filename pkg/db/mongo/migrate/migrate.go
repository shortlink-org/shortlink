package migrate

import (
	"context"
	"embed"

	"github.com/shortlink-org/shortlink/pkg/db"
)

// Migration - apply migration to db
// TODO: not support v2
// issue: https://github.com/golang-migrate/migrate/issues/1145
func Migration(_ context.Context, store db.DB, fs embed.FS, tableName string) error {
	// client, ok := store.GetConn().(*mongo.Client)
	// if !ok {
	// 	return db.ErrGetConnection
	// }
	//
	// driverMigrations, err := iofs.New(fs, "migrations")
	// if err != nil {
	// 	return err
	// }
	//
	// driverDB, err := mongodb.WithInstance(client, &mongodb.Config{
	// 	DatabaseName:         "shortlink",
	// 	MigrationsCollection: fmt.Sprintf("schema_migrations_%s", strings.ReplaceAll(tableName, "-", "_")),
	// })
	// if err != nil {
	// 	return err
	// }
	//
	// m, err := migrate.NewWithInstance("iofs", driverMigrations, "mongodb", driverDB)
	// if err != nil {
	// 	return err
	// }
	//
	// err = m.Up()
	// if err != nil && err.Error() != "no change" {
	// 	return err
	// }

	return nil
}
