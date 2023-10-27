//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate -f ./schema/sqlc.yaml

package postgres

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	"github.com/shortlink-org/shortlink/internal/pkg/db/postgres/migrate"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/postgres/schema/crud"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
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
		return nil, errors.New("error get connection to PostgreSQL")
	}

	s.newClient = crud.New(s.client)

	// Migration -------------------------------------------------------------------------------------------------------
	err := migrate.Migration(ctx, store, migrations, "repository_link")
	if err != nil {
		return nil, err
	}

	// Create a batch job ----------------------------------------------------------------------------------------------
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) any { //nolint:errcheck // ignore
			sources := make([]*domain.Link, len(args))

			for key := range args {
				sources[key] = args[key].Item.(*domain.Link) //nolint:errcheck // ignore
			}

			dataList, errBatchWrite := s.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CallbackChannel <- errors.New("error write to PostgreSQL")
				}

				return errBatchWrite
			}

			for key, item := range dataList.GetLink() {
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

// Get - a get link
func (s *Store) Get(ctx context.Context, hash string) (*domain.Link, error) {
	link, err := s.newClient.GetLinkByHash(ctx, hash)
	if err != nil {
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: hash}, Err: fmt.Errorf("failed get link: %s", hash)}
	}

	return &domain.Link{
		Url:       link.Url,
		Hash:      link.Hash,
		Describe:  link.Describe.String,
		CreatedAt: timestamppb.New(link.CreatedAt.Time),
		UpdatedAt: timestamppb.New(link.UpdatedAt.Time),
	}, nil
}

// List - list links
func (s *Store) List(ctx context.Context, filter *query.Filter) (*domain.Links, error) {
	// query builder
	links := psql.Select("url, hash, describe, created_at, updated_at").
		From("link.links")

	if filter != nil {
		links = links.
			Limit(uint64(filter.Pagination.Limit)).
			Offset(uint64(filter.Pagination.Page * filter.Pagination.Limit))
	}

	links = s.buildFilter(links, filter)
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &domain.Links{
				Link: []*domain.Link{},
			}, nil
		}

		return nil, &domain.NotFoundError{Link: &domain.Link{}, Err: query.ErrNotFound}
	}

	response := &domain.Links{
		Link: []*domain.Link{},
	}

	for rows.Next() {
		var result domain.Link
		var (
			created_ad sql.NullTime
			updated_at sql.NullTime
		)
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe, &created_ad, &updated_at)
		if err != nil {
			return nil, &domain.NotFoundError{Link: &domain.Link{}, Err: query.ErrNotFound}
		}
		result.CreatedAt = &timestamppb.Timestamp{Seconds: created_ad.Time.Unix(), Nanos: int32(created_ad.Time.Nanosecond())}
		result.UpdatedAt = &timestamppb.Timestamp{Seconds: updated_at.Time.Unix(), Nanos: int32(updated_at.Time.Nanosecond())}

		response.Link = append(response.GetLink(), &result)
	}

	return response, nil
}

// Add - an add link
func (s *Store) Add(ctx context.Context, source *domain.Link) (*domain.Link, error) {
	switch s.config.mode {
	case options.MODE_BATCH_WRITE:
		cb := s.config.job.Push(source)

		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, data
		case *domain.Link:
			return data, nil
		default:
			return nil, nil
		}
	case options.MODE_SINGLE_WRITE:
		data, err := s.singleWrite(ctx, source)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, nil
}

// Update - update link
func (s *Store) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	_, err := s.newClient.UpdateLink(ctx, crud.UpdateLinkParams{
		Url:  in.Url,
		Hash: in.Hash,
		Describe: pgtype.Text{
			String: in.Describe,
			Valid:  true,
		},
		Json: *in,
	})
	if err != nil {
		return nil, err
	}

	return in, nil
}

// Delete - delete link
func (s *Store) Delete(ctx context.Context, hash string) error {
	err := s.newClient.DeleteLink(ctx, hash)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) singleWrite(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	err := domain.NewURL(in)
	if err != nil {
		return nil, err
	}

	// save as JSON. it doesn't make sense
	dataJson, err := protojson.Marshal(in)
	if err != nil {
		return nil, err
	}

	// query builder
	links := psql.Insert("link.links").
		Columns("url", "hash", "describe", "json").
		Values(in.GetUrl(), in.GetHash(), in.GetDescribe(), string(dataJson))

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.client.Exec(ctx, q, args...)
	if err != nil {
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: in.GetHash()}, Err: fmt.Errorf("failed create link: %s", in.GetHash())}
	}

	return in, nil
}

func (s *Store) batchWrite(ctx context.Context, sources []*domain.Link) (*domain.Links, error) {
	// Create a new link
	for key := range sources {
		err := domain.NewURL(sources[key])
		if err != nil {
			return nil, err
		}
	}

	links := psql.Insert("link.links").Columns("url", "hash", "describe", "json")

	// query builder
	for _, source := range sources {
		// save as JSON. it doesn't make sense
		dataJson, err := protojson.Marshal(source)
		if err != nil {
			return nil, err
		}

		links = links.Values(source.GetUrl(), source.GetHash(), source.GetDescribe(), dataJson)
	}

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := s.client.QueryRow(ctx, q, args...)
	errScan := row.Scan(&sources)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return &domain.Links{
			Link: []*domain.Link{},
		}, nil
	}
	if errScan != nil {
		return nil, fmt.Errorf("error save link")
	}

	response := &domain.Links{
		Link: []*domain.Link{},
	}

	response.Link = append(response.GetLink(), sources...)

	return response, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode writes to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
