//go:generate protoc -I../../../internal/api/domain/link --gotemplate_out=all=true,template_dir=template:. link.proto
//go:generate go-bindata -prefix migrations -pkg migrations -ignore migrations.go -o migrations/migrations.go migrations
package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/api/infrastructure/store/postgres/migrations"
	"github.com/batazor/shortlink/internal/db/options"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

// Init ...
func (p *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	p.setConfig()

	// Apply migration
	err = p.migrate()
	if err != nil {
		return err
	}

	// Connect to Postgres
	if p.client, err = pgxpool.Connect(ctx, p.config.URI); err != nil {
		return err
	}

	return nil
}

// Close ...
func (p *Store) Close() error { // nolint unparam
	p.client.Close()
	return nil
}

// Migrate ...
func (p *Store) migrate() error { // nolint unused
	// Create connect
	db, err := sql.Open("postgres", p.config.URI)
	if err != nil {
		return err
	}

	// wrap assets into Resource
	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})

	driver, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", driver, p.config.URI)
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
func (p *Store) setConfig() {
	dbinfo := fmt.Sprintf("postgres://%s:%s@localhost:5435/%s?sslmode=disable", "shortlink", "shortlink", "shortlink")

	viper.AutomaticEnv()
	viper.SetDefault("STORE_POSTGRES_URI", dbinfo)                           // Postgres URI
	viper.SetDefault("STORE_POSTGRES_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db

	p.config = Config{
		URI:  viper.GetString("STORE_POSTGRES_URI"),
		mode: viper.GetInt("STORE_POSTGRES_MODE_WRITE"),
	}
}
