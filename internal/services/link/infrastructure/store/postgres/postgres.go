//go:generate protoc -I../../../../../services/link/domain/link/v1 --gotemplate_out=all=true,template_dir=template:. link.proto
package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/batch"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/db/options"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

// Init ...
func (s *Store) Init(ctx context.Context, db *db.Store) error {
	// Set configuration
	s.setConfig()
	s.client = db.Store.GetConn().(*pgxpool.Pool)

	// Create batch job
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) interface{} {
			sources := make([]*v1.Link, len(args))

			for key := range args {
				sources[key] = args[key].Item.(*v1.Link)
			}

			dataList, errBatchWrite := s.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CB <- errors.New("Error write to PostgreSQL")
				}
				return errBatchWrite
			}

			for key, item := range dataList.Link {
				args[key].CB <- item
			}

			return nil
		}

		var err error
		s.config.job, err = batch.New(ctx, cb)
		if err != nil {
			return err
		}

		go s.config.job.Run(ctx)
	}

	return nil
}

// Get ...
func (p *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
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
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	if rows.Err() != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response v1.Link
	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
		}
	}

	if response.Hash == "" {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (p *Store) List(ctx context.Context, filter *query.Filter) (*v1.Links, error) {
	// query builder
	links := psql.Select("url, hash, describe, created_at, updated_at").
		From("shortlink.links")
	links = p.buildFilter(links, filter)
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
	}

	response := &v1.Links{
		Link: []*v1.Link{},
	}

	for rows.Next() {
		var result v1.Link
		var (
			created_ad sql.NullTime
			updated_at sql.NullTime
		)
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe, &created_ad, &updated_at)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
		}
		result.CreatedAt = &timestamp.Timestamp{Seconds: int64(created_ad.Time.Second()), Nanos: int32(created_ad.Time.Nanosecond())}
		result.UpdatedAt = &timestamp.Timestamp{Seconds: int64(updated_at.Time.Second()), Nanos: int32(updated_at.Time.Nanosecond())}

		response.Link = append(response.Link, &result)
	}

	return response, nil
}

// Add ...
func (p *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	switch p.config.mode {
	case options.MODE_BATCH_WRITE:
		cb, err := p.config.job.Push(source)
		if err != nil {
			return nil, err
		}

		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, data
		case *v1.Link:
			return data, nil
		default:
			return nil, nil
		}
	case options.MODE_SINGLE_WRITE:
		data, err := p.singleWrite(ctx, source)
		return data, err
	}

	return nil, nil
}

// Update ...
func (p *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete ...
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
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

func (p *Store) singleWrite(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	// save as JSON. it doesn't make sense
	dataJson, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := psql.Insert("shortlink.links").
		Columns("url", "hash", "describe", "json").
		Values(source.Url, source.Hash, source.Describe, dataJson)

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&source.Url, &source.Hash, &source.Describe)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return source, nil
	}
	if errScan.Error() != "" {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

func (p *Store) batchWrite(ctx context.Context, sources []*v1.Link) (*v1.Links, error) {
	// Create a new link
	for key := range sources {
		err := v1.NewURL(sources[key])
		if err != nil {
			return nil, err
		}
	}

	links := psql.Insert("shortlink.links").Columns("url", "hash", "describe", "json")

	// query builder
	for _, source := range sources {
		// save as JSON. it doesn't make sense
		dataJson, err := json.Marshal(source)
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
		return &v1.Links{
			Link: []*v1.Link{},
		}, nil
	}
	if errScan != nil {
		return nil, fmt.Errorf("Error save link")
	}

	response := &v1.Links{
		Link: []*v1.Link{},
	}

	for item := range sources {
		response.Link = append(response.Link, sources[item])
	}

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
