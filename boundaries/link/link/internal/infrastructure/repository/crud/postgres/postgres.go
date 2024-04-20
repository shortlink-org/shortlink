//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate -f ./schema/sqlc.yaml

package postgres

import (
	"context"
	"embed"
	"encoding/json"
	"errors"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/spf13/viper"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres/filter"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres/schema/crud"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
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
		return nil, &v1.NotFoundByHashError{Hash: hash}
	}

	resp, err := domain.NewLinkBuilder().
		SetURL(link.Url).
		SetDescribe(link.Describe).
		Build()

	return resp, nil
}

// List - return list links
func (s *Store) List(ctx context.Context, params *v1.FilterLink) (*domain.Links, error) {
	request := psql.Select("url", "hash", "describe", "created_at", "updated_at").
		From("link.links")

	// Build filter
	request = filter.NewFilter(params).BuildFilter(request)

	q, args, err := request.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil || rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	links := domain.NewLinks()
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

		link, errBuilder := domain.NewLinkBuilder().
			SetURL(url).
			SetDescribe(describe).
			SetCreatedAt(createdAt.Time).
			SetUpdatedAt(updatedAt.Time).
			Build()

		if errBuilder != nil {
			return nil, errBuilder
		}

		links.Push(link)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return links, nil
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
	link := in.GetUrl()
	_, err := s.query.UpdateLink(ctx, crud.UpdateLinkParams{
		Url:      link.String(),
		Hash:     in.GetHash(),
		Describe: in.GetDescribe(),
		Json:     *in,
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
	// save as JSON. it doesn't make sense
	dataJson, err := json.Marshal(in)
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
		return nil, &v1.NotFoundByHashError{Hash: in.GetHash()}
	}

	return in, nil
}

func (s *Store) batchWrite(ctx context.Context, in *domain.Links) (*domain.Links, error) {
	links := make([]crud.CreateLinksParams, 0, len(in.GetLink()))

	// Create a new link
	list := in.GetLink()
	for key := range list {
		link := list[key].GetUrl()
		links = append(links, crud.CreateLinksParams{
			Url:      link.String(),
			Hash:     list[key].GetHash(),
			Describe: list[key].GetDescribe(),
			Json:     *list[key],
		})
	}

	_, err := s.query.CreateLinks(ctx, links)
	if err != nil {
		errs := make([]error, 0, len(list))
		for key := range list {
			errs = append(errs, &v1.CreateLinkError{Link: *list[key]})
		}

		return nil, errors.Join(errs...)
	}

	return in, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode writes to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
