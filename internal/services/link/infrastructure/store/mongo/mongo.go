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

	"github.com/batazor/shortlink/internal/pkg/batch"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/db/options"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
)

// Init ...
func (s *Store) Init(ctx context.Context, db *db.Store) error {
	// Set configuration
	s.setConfig()
	s.client = db.Store.GetConn().(*mongo.Client)

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
					args[index].CB <- query.StoreError{Value: "Error write to MongoDB"}
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

// Add ...
func (m *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
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
		case *v1.Link:
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
func (m *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	val := collection.FindOne(ctx, bson.D{primitive.E{Key: "hash", Value: id}})

	if val.Err() != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response v1.Link

	if err := val.Decode(&response); err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	return &response, nil
}

// List ...
func (m *Store) List(ctx context.Context, filter *query.Filter) (*v1.Links, error) { // nolint unused
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// build Filter
	filterQuery := bson.D{}
	if filter != nil {
		filterQuery = getFilter(filter)
	}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(ctx, filterQuery)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
	}

	if cur.Err() != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
	}

	// Here's an array in which you can db the decoded documents
	links := &v1.Links{
		Link: []*v1.Link{},
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem v1.Link
		if errDecode := cur.Decode(&elem); errDecode != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
		}

		links.Link = append(links.Link, &elem)
	}

	// Close the cursor once finished
	err = cur.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	return links, nil
}

// Update ...
func (m *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete ...
func (m *Store) Delete(ctx context.Context, id string) error {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{primitive.E{Key: "hash", Value: id}})
	if err != nil {
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}

func (m *Store) singleWrite(ctx context.Context, source *v1.Link) (*v1.Link, error) { // nolint unused
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, &source)
	if err != nil {
		var typeErr mongo.WriteException
		errors.As(err, &typeErr)

		if errors.As(err, &typeErr) {
			switch typeErr.WriteErrors[0].Code {
			case 11000:
				return nil, &v1.NotUniqError{Link: source, Err: fmt.Errorf("Duplicate URL: %s", source.Url)}
			default:
				return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed marsharing link: %s", source.Url)}
			}
		}

		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed marsharing link: %s", source.Url)}
	}

	return source, nil
}

func (m *Store) batchWrite(ctx context.Context, sources []*v1.Link) (*v1.Links, error) { // nolint unused
	docs := make([]interface{}, len(sources))

	// Create a new link
	for key := range sources {
		err := v1.NewURL(sources[key])
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

	links := &v1.Links{
		Link: []*v1.Link{},
	}

	for item := range sources {
		links.Link = append(links.Link, sources[item])
	}

	return links, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
