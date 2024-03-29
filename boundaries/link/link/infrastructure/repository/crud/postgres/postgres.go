//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate -f ./schema/sqlc.yaml

package postgres

import (
	"context"
	"embed"
	"errors"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/infrastructure/repository/crud/postgres/schema/crud"
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
			sources := make([]*domain.Link, len(args))

			for key := range args {
				sources[key], ok = args[key].Item.(*domain.Link)
				if !ok {
					args[key].CallbackChannel <- batch.ErrInvalidType
				}
			}

			dataList, errBatchWrite := s.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CallbackChannel <- ErrWrite
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
	link, err := s.query.GetLinkByHash(ctx, hash)
	if err != nil {
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: hash}}
	}

	return &domain.Link{
		Url:       link.Url,
		Hash:      link.Hash,
		Describe:  link.Describe,
		CreatedAt: timestamppb.New(link.CreatedAt.Time),
		UpdatedAt: timestamppb.New(link.UpdatedAt.Time),
	}, nil
}

// List - return list links
func (s *Store) List(ctx context.Context, filter *domain.FilterLink) (*domain.Links, error) {
	// Set default filter
	if filter == nil {
		filter = &domain.FilterLink{}
	}

	request := psql.Select("url", "hash", "describe", "created_at", "updated_at").
		From("link.links")

	request = filter.BuildFilter(request)

	q, args, err := request.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil || rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]*domain.Link, 0)
	for rows.Next() {
		var (
			url       string
			hash      string
			describe  string
			createdAt pq.NullTime
			updatedAt pq.NullTime
		)

		err = rows.Scan(&url, &hash, &describe, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		links = append(links, &domain.Link{
			Url:       url,
			Hash:      hash,
			Describe:  describe,
			CreatedAt: timestamppb.New(createdAt.Time),
			UpdatedAt: timestamppb.New(updatedAt.Time),
		})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &domain.Links{
		Link: links,
	}, nil
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
	_, err := s.query.UpdateLink(ctx, crud.UpdateLinkParams{
		Url:      in.GetUrl(),
		Hash:     in.GetHash(),
		Describe: in.GetDescribe(),
		Json: domain.Link{
			Url:       in.GetUrl(),
			Hash:      in.GetHash(),
			Describe:  in.GetDescribe(),
			CreatedAt: in.GetCreatedAt(),
		},
	})
	if err != nil {
		return nil, err
	}

	return in, nil
}

// Delete - delete link
func (s *Store) Delete(ctx context.Context, hash string) error {
	err := s.query.DeleteLink(ctx, hash)
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

	links := psql.Insert("link.links").
		Columns("url", "hash", "describe", "json").
		Values(in.GetUrl(), in.GetHash(), in.GetDescribe(), string(dataJson))

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.client.Exec(ctx, q, args...)
	if err != nil {
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: in.GetHash()}}
	}

	return in, nil
}

func (s *Store) batchWrite(ctx context.Context, in []*domain.Link) (*domain.Links, error) {
	links := make([]crud.CreateLinksParams, 0, len(in))

	// Create a new link
	for key := range in {
		err := domain.NewURL(in[key])
		if err != nil {
			return nil, err
		}

		links = append(links, crud.CreateLinksParams{
			Url:      in[key].GetUrl(),
			Hash:     in[key].GetHash(),
			Describe: in[key].GetDescribe(),
			Json: domain.Link{
				Url:      in[key].GetUrl(),
				Hash:     in[key].GetHash(),
				Describe: in[key].GetDescribe(),
			},
		})
	}

	_, err := s.query.CreateLinks(ctx, links)
	if err != nil {
		errs := make([]error, 0, len(in))
		for key := range in {
			errs = append(errs, &domain.CreateLinkError{Link: &domain.Link{Hash: in[key].GetHash()}})
		}

		return nil, errors.Join(errs...)
	}

	return &domain.Links{
		Link: in,
	}, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode writes to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
