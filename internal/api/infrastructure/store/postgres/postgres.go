//go:generate protoc -I../../../internal/api/domain/link --gotemplate_out=all=true,template_dir=template:. link.proto
//go:generate go-bindata -prefix migrations -pkg migrations -ignore migrations.go -o migrations/migrations.go migrations
package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/postgres/migrations"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/batch"
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

	// Create batch job
	if p.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) interface{} {
			sources := make([]*link.Link, len(args))

			for key := range args {
				sources[key] = args[key].Item.(*link.Link)
			}

			dataList, errBatchWrite := p.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CB <- errors.New("Error write to PostgreSQL")
				}
				return errBatchWrite
			}

			for key, item := range dataList {
				args[key].CB <- item
			}

			return nil
		}
		p.config.job, err = batch.New(ctx, cb)
		if err != nil {
			return err
		}

		go p.config.job.Run(ctx)
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

// Get ...
func (p *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)

	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	if rows.Err() != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link
	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
		}
	}

	if response.Hash == "" {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (p *Store) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
	// query builder
	links := psql.Select("url, hash, describe").
		From("links")
	links = p.buildFilter(links, filter)
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	var response []*link.Link

	for rows.Next() {
		var result link.Link
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		response = append(response, &result)
	}

	return response, nil
}

// Add ...
func (p *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	switch p.config.mode {
	case options.MODE_BATCH_WRITE:
		cb, err := p.config.job.Push(source)
		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, err
		case link.Link:
			return &data, nil
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
func (p *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (p *Store) Delete(ctx context.Context, id string) error {
	// query builder
	links := psql.Delete("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = p.client.Exec(ctx, q, args...)
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

func (p *Store) singleWrite(ctx context.Context, source *link.Link) (*link.Link, error) { // nolint unused
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	// save as JSON. it doesn't make sense
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// query builder
	links := psql.Insert("links").
		Columns("url", "hash", "describe", "json").
		Values(data.Url, data.Hash, data.Describe, dataJson)

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&data.Url, &data.Hash, &data.Describe).Error()
	if errScan == "no rows in result set" {
		return data, nil
	}
	if errScan != "" {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return data, nil
}

func (p *Store) batchWrite(ctx context.Context, sources []*link.Link) ([]*link.Link, error) {
	docs := make([]interface{}, len(sources))

	// Create a new link
	for key := range sources {
		data, err := link.NewURL(sources[key].Url)
		if err != nil {
			return nil, err
		}

		docs[key] = data
	}

	links := psql.Insert("links").Columns("url", "hash", "describe", "json")

	// query builder
	for _, source := range sources {
		// save as JSON. it doesn't make sense
		dataJson, err := json.Marshal(source)
		if err != nil {
			return nil, err
		}

		links.Values(source.Url, source.Hash, source.Describe, dataJson)
	}

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	p.client.QueryRow(ctx, q, args...)

	return sources, nil
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
