//go:generate protoc -I../../../../../link/domain/link/domain --gotemplate_out=all=true,template_dir=template:. link.proto
package postgres

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	"github.com/shortlink-org/shortlink/internal/pkg/db/postgres/migrate"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS

	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

// New store
func New(ctx context.Context, db *db.Store) (*Store, error) {
	var ok bool
	s := &Store{}

	// Set configuration -----------------------------------------------------------------------------------------------
	s.setConfig()
	s.client, ok = db.Store.GetConn().(*pgxpool.Pool)
	if !ok {
		return nil, errors.New("Error get connection to PostgreSQL")
	}

	// Migration -------------------------------------------------------------------------------------------------------
	err := migrate.Migration(ctx, db, migrations, viper.GetString("SERVICE_NAME"))
	if err != nil {
		return nil, err
	}

	// Create a batch job ----------------------------------------------------------------------------------------------
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) interface{} { // nolint:errcheck
			sources := make([]*domain.Link, len(args))

			for key := range args {
				sources[key] = args[key].Item.(*domain.Link) // nolint:errcheck
			}

			dataList, errBatchWrite := s.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CallbackChannel <- errors.New("Error write to PostgreSQL")
				}

				return errBatchWrite
			}

			for key, item := range dataList.Link {
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

// Get - get link
func (p *Store) Get(ctx context.Context, id string) (*domain.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("shortlink.links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	if rows.Err() != nil {
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response domain.Link
	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &domain.NotFoundError{Link: &domain.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
		}
	}

	if response.Hash == "" {
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List - list links
func (p *Store) List(ctx context.Context, filter *query.Filter) (*domain.Links, error) {
	// query builder
	links := psql.Select("url, hash, describe, created_at, updated_at").
		From("shortlink.links")

	if filter != nil {
		links = links.
			Limit(uint64(filter.Pagination.Limit)).
			Offset(uint64(filter.Pagination.Page * filter.Pagination.Limit))
	}

	links = p.buildFilter(links, filter)
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)
	if err != nil {
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
		result.CreatedAt = &timestamp.Timestamp{Seconds: int64(created_ad.Time.Second()), Nanos: int32(created_ad.Time.Nanosecond())}
		result.UpdatedAt = &timestamp.Timestamp{Seconds: int64(updated_at.Time.Second()), Nanos: int32(updated_at.Time.Nanosecond())}

		response.Link = append(response.Link, &result)
	}

	return response, nil
}

// Add - add link
func (p *Store) Add(ctx context.Context, source *domain.Link) (*domain.Link, error) {
	switch p.config.mode {
	case options.MODE_BATCH_WRITE:
		cb := p.config.job.Push(source)

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
		data, err := p.singleWrite(ctx, source)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, nil
}

// Update - update link
func (p *Store) Update(_ context.Context, _ *domain.Link) (*domain.Link, error) {
	return nil, nil
}

// Delete - delete link
func (p *Store) Delete(ctx context.Context, id string) error {
	// query builder
	request := psql.Delete("shortlink.links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := request.ToSql()
	if err != nil {
		return err
	}

	_, err = p.client.Exec(ctx, q, args...)
	if err != nil {
		return &domain.NotFoundError{Link: &domain.Link{Hash: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

func (p *Store) singleWrite(ctx context.Context, source *domain.Link) (*domain.Link, error) {
	err := domain.NewURL(source)
	if err != nil {
		return nil, err
	}

	// save as JSON. it doesn't make sense
	dataJson, err := protojson.Marshal(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := psql.Insert("shortlink.links").
		Columns("url", "hash", "describe", "json").
		Values(source.Url, source.Hash, source.Describe, string(dataJson))

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)

	errScan := row.Scan()
	if errors.Is(errScan, pgx.ErrNoRows) {
		return source, nil
	}
	if errScan.Error() != "" {
		return nil, &domain.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

func (p *Store) batchWrite(ctx context.Context, sources []*domain.Link) (*domain.Links, error) {
	// Create a new link
	for key := range sources {
		err := domain.NewURL(sources[key])
		if err != nil {
			return nil, err
		}
	}

	links := psql.Insert("shortlink.links").Columns("url", "hash", "describe", "json")

	// query builder
	for _, source := range sources {
		// save as JSON. it doesn't make sense
		dataJson, err := protojson.Marshal(source)
		if err != nil {
			return nil, err
		}

		links = links.Values(source.Url, source.Hash, source.Describe, dataJson)
	}

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)
	errScan := row.Scan(&sources)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return &domain.Links{
			Link: []*domain.Link{},
		}, nil
	}
	if errScan != nil {
		return nil, fmt.Errorf("Error save link")
	}

	response := &domain.Links{
		Link: []*domain.Link{},
	}

	response.Link = append(response.Link, sources...)

	return response, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
