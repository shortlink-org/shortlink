package mongo

import (
	"context"
	"embed"
	"errors"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/shortlink-org/go-sdk/batch"
	"github.com/shortlink-org/go-sdk/config"
	"github.com/shortlink-org/go-sdk/db"
	"github.com/shortlink-org/go-sdk/db/drivers/mongo/migrate"
	"github.com/shortlink-org/go-sdk/db/options"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/mongo/dto"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/mongo/filter"
	types "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/types/v1"
)

//go:embed migrations/*.json
var migrations embed.FS

// New store
func New(ctx context.Context, store db.DB, cfg *config.Config) (*Store, error) {
	var ok bool
	s := &Store{}

	// Set configuration -----------------------------------------------------------------------------------------------
	s.setConfig()
	s.client, ok = store.GetConn().(*mongo.Client)
	if !ok {
		return nil, db.ErrGetConnection
	}

	// Migration -------------------------------------------------------------------------------------------------------
	err := migrate.Migration(ctx, store, migrations, "repository_link")
	if err != nil {
		return nil, err
	}

	// Create a batch job
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(items []*batch.Item[*v1.Link]) error {
			sources := make([]*v1.Link, len(items))

			for i, item := range items {
				sources[i] = item.Item
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
		s.config.job, err = batch.NewSync[*v1.Link](ctx, cb)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

// Add - add
func (s *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	switch s.config.mode {
	case options.MODE_BATCH_WRITE:
		resCh := s.config.job.Push(source)

		select {
		case res, ok := <-resCh:
			if !ok || res == nil {
				return nil, ErrWrite
			}
			return res, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	case options.MODE_SINGLE_WRITE:
		data, err := s.singleWrite(ctx, source)
		return data, err
	default:
		return nil, nil
	}
}

// Get - get
func (s *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	collection := s.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	val := collection.FindOne(ctx, bson.D{bson.E{Key: "hash", Value: id}})

	if val.Err() != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	var response dto.Link

	if err := val.Decode(&response); err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	link, err := response.ToDomain()
	if err != nil {
		return nil, err
	}

	return link, nil
}

// List - list
func (s *Store) List(ctx context.Context, params *types.FilterLink) (*v1.Links, error) {
	collection := s.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// Build filter
	filterQuery := filter.NewFilter(params).BuildMongoFilter()

	cur, err := collection.Find(ctx, filterQuery)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}

	if cur.Err() != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}

	links := v1.NewLinks()

	for cur.Next(ctx) {
		var elem dto.Link
		if errDecode := cur.Decode(&elem); errDecode != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}}
		}

		// convert to domain
		link, errToDomain := elem.ToDomain()
		if errToDomain != nil {
			return nil, errToDomain
		}

		links.Push(link)
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return links, nil
}

// Update - update
func (s *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete - delete
func (s *Store) Delete(ctx context.Context, id string) error {
	collection := s.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{bson.E{Key: "hash", Value: id}})
	if err != nil {
		return &v1.NotFoundByHashError{Hash: id}
	}

	return nil
}

func (s *Store) singleWrite(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	collection := s.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// convert to DTO
	link, err := dto.FromDomain(source)
	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(ctx, link)
	if err != nil {
		var writeErr mongo.WriteException
		if errors.As(err, &writeErr) {
			if writeErr.HasErrorCode(11000) {
				return nil, &v1.NotUniqError{Link: source}
			}
		}

		return nil, err
	}

	return source, nil
}

func (s *Store) batchWrite(ctx context.Context, sources []*v1.Link) (*v1.Links, error) {
	docs := make([]interface{}, len(sources))
	for i, source := range sources {
		// convert to DTO
		link, err := dto.FromDomain(source)
		if err != nil {
			return nil, err
		}

		docs[i] = link
	}

	collection := s.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}

	links := v1.NewLinks()
	links.Push(sources...)

	return links, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode writes to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
