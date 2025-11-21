//go:generate sqlc generate -f ./schema/sqlc.yaml

package postgres

import (
	"context"
	"embed"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"

	"github.com/shortlink-org/go-sdk/batch"
	"github.com/shortlink-org/go-sdk/config"
	"github.com/shortlink-org/go-sdk/db"
	"github.com/shortlink-org/go-sdk/db/drivers/postgres/migrate"
	"github.com/shortlink-org/go-sdk/db/options"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/postgres/schema/crud"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS

	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

// New store
func New(ctx context.Context, store db.DB, cfg *config.Config) (*Store, error) {
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
		cb := func(items []*batch.Item[*domain.Link]) error {
			sources := &domain.Links{}

			for _, item := range items {
				sources.Push(item.Item)
			}

			dataList, errBatchWrite := s.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for _, item := range items {
					item.CallbackChannel <- nil
					close(item.CallbackChannel)
				}
				return errBatchWrite
			}

			for i, link := range dataList.GetLinks() {
				items[i].CallbackChannel <- link
				close(items[i].CallbackChannel)
			}

			return nil
		}

		var err error
		s.config.job, err = batch.NewSync[*domain.Link](ctx, cfg, cb)
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
