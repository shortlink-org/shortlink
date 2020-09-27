//go:generate go-bindata -prefix migrations -pkg migrations -ignore migrations.go -o migrations/migrations.go migrations
package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/batch"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/db/options"
)

// Init ...
func (s *Store) Init(ctx context.Context, db *db.Store) error {
	// Set configuration
	s.setConfig()
	s.client = db.Store.GetConn().(*mongo.Client)

	// Create batch job
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) interface{} {
			sources := make([]*link.Link, len(args))

			for key := range args {
				sources[key] = args[key].Item.(*link.Link)
			}

			dataList, errBatchWrite := s.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CB <- query.StoreError{Value: "Error write to MongoDB"}
				}
				return errBatchWrite
			}

			for key, item := range dataList {
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

// Add ...
func (m *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	switch m.config.mode {
	case options.MODE_BATCH_WRITE:
		cb, err := m.config.job.Push(source)
		if err != nil {
			return nil, err
		}

		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, err
		case query.StoreError:
			return nil, errors.New(data.Value)
		case *link.Link:
			return data, nil
		default:
			return nil, nil
		}
	case options.MODE_SINGLE_WRITE:
		data, err := m.singleWrite(ctx, source)
		return data, err
	}

	return nil, nil
}

// Get ...
func (m *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	val := collection.FindOne(ctx, bson.D{primitive.E{Key: "hash", Value: id}})

	if val.Err() != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	if err := val.Decode(&response); err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	return &response, nil
}

// List ...
func (m *Store) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// build Filter
	filterQuery := bson.D{}
	if filter != nil {
		filterQuery = getFilter(filter)
	}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(ctx, filterQuery)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	if cur.Err() != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	// Here's an array in which you can db the decoded documents
	var response []*link.Link

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem link.Link
		if errDecode := cur.Decode(&elem); errDecode != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		response = append(response, &elem)
	}

	// Close the cursor once finished
	err = cur.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Update ...
func (m *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (m *Store) Delete(ctx context.Context, id string) error {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{primitive.E{Key: "hash", Value: id}})
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}

func (m *Store) singleWrite(ctx context.Context, source *link.Link) (*link.Link, error) { // nolint unused
	err := link.NewURL(source)
	if err != nil {
		return nil, err
	}

	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, &source)
	if err != nil {
		switch err.(mongo.WriteException).WriteErrors[0].Code {
		case 11000:
			return nil, &link.NotUniqError{Link: source, Err: fmt.Errorf("Duplicate URL: %s", source.Url)}
		default:
			return nil, &link.NotFoundError{Link: source, Err: fmt.Errorf("Failed marsharing link: %s", source.Url)}
		}
	}

	return source, nil
}

func (m *Store) batchWrite(ctx context.Context, sources []*link.Link) ([]*link.Link, error) { // nolint unused
	docs := make([]interface{}, len(sources))

	// Create a new link
	for key := range sources {
		err := link.NewURL(sources[key])
		if err != nil {
			return nil, err
		}

		docs[key] = sources[key]
	}

	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}

	return sources, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
