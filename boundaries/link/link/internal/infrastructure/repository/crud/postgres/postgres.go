//go:generate sqlc generate -f ./schema/sqlc.yaml

package postgres

import (
	"context"
	"embed"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres/schema/crud"
	"github.com/shortlink-org/shortlink/pkg/batch"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/db/options"
	"github.com/shortlink-org/shortlink/pkg/db/postgres/migrate"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS

	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

// New store
func New(ctx context.Context, store db.DB) (*Store, error) {
	var ok bool
	s := &Store{}

	// Set configuration -----------------------------------------------------------------------------------------------
	s.setConfig()
	s.client, ok = store.GetConn().(*pgxpool.Pool)
	if !ok {
		return nil, db.ErrGetConnection
	}

	s.query = crud.New(s.client)

	// Migration -------------------------------------------------------------------------------------------------------
	err := migrate.Migration(ctx, store, migrations, "repository_link")
	if err != nil {
		return nil, err
	}

	// Create a batch job ----------------------------------------------------------------------------------------------
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) any { //nolint:errcheck // ignore
			sources := domain.NewLinks()

			for key := range args {
				link, ok := args[key].Item.(*domain.Link)
				if !ok {
					args[key].CallbackChannel <- batch.ErrInvalidType
				}

				sources.Push(link)
			}

			dataList, errBatchWrite := s.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CallbackChannel <- ErrWrite
				}

				return errBatchWrite
			}

			for key, item := range dataList.GetLinks() {
				args[key].CallbackChannel <- item
			}

			return nil
		}

		var err error
		s.config.job, err = batch.New(ctx, cb)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode writes to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
