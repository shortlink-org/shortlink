//go:generate go-bindata -prefix migrations -pkg migrations -ignore migrations.go -o migrations/migrations.go migrations
package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/batch"
	"github.com/batazor/shortlink/internal/store/mongo/migrations"
	storeOptions "github.com/batazor/shortlink/internal/store/options"
	"github.com/batazor/shortlink/internal/store/query"
)

// Init ...
func (m *MongoLinkList) Init(ctx context.Context) error {
	var err error

	// Set configuration
	m.setConfig()

	// Connect to MongoDB
	m.client, err = mongo.NewClient(options.Client().ApplyURI(m.config.URI))
	if err != nil {
		return err
	}

	err = m.client.Connect(ctx)
	if err != nil {
		return err
	}

	// TODO: check correct ping
	// Check connect
	//ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	//err = m.client.Ping(ctx, readpref.Primary())
	//if err != nil {
	//	return err
	//}

	// Apply migration
	err = m.migrate()
	if err != nil {
		return err
	}

	// Create batch job
	if m.config.mode == storeOptions.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) interface{} {
			sources := make([]*link.Link, len(args))

			for key := range args {
				sources[key] = args[key].Item.(*link.Link)
			}

			dataList, errBatchWrite := m.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CB <- errors.New("Error write to MongoDB")
				}
				return errBatchWrite
			}

			for key, item := range dataList {
				args[key].CB <- item
			}

			return nil
		}
		m.config.job, err = batch.New(ctx, cb)
		if err != nil {
			return err
		}

		go m.config.job.Run(ctx)
	}

	return nil
}

// Close ...
func (m *MongoLinkList) Close() error {
	return m.client.Disconnect(context.Background())
}

// Migrate ...
func (m *MongoLinkList) migrate() error { // nolint unused
	// wrap assets into Resource
	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})

	driver, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}

	ms, err := migrate.NewWithSourceInstance("go-bindata", driver, m.config.URI)
	if err != nil {
		return err
	}

	err = ms.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}

// Add ...
func (m *MongoLinkList) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	switch m.config.mode {
	case storeOptions.MODE_BATCH_WRITE:
		cb, err := m.config.job.Push(source)
		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, err
		case link.Link:
			return &data, nil
		default:
			return nil, nil
		}
	case storeOptions.MODE_SINGLE_WRITE:
		data, err := m.singleWrite(ctx, source)
		return data, err
	}

	return nil, nil
}

// Get ...
func (m *MongoLinkList) Get(ctx context.Context, id string) (*link.Link, error) {
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
func (m *MongoLinkList) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
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

	// Here's an array in which you can store the decoded documents
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
func (m *MongoLinkList) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (m *MongoLinkList) Delete(ctx context.Context, id string) error {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{primitive.E{Key: "hash", Value: id}})
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}

func (m *MongoLinkList) singleWrite(ctx context.Context, source *link.Link) (*link.Link, error) { // nolint unused
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, &data)
	if err != nil {
		switch err.(mongo.WriteException).WriteErrors[0].Code {
		case 11000:
			return nil, &link.NotUniqError{Link: data, Err: fmt.Errorf("Duplicate URL: %s", data.Url)}
		default:
			return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed marsharing link: %s", data.Url)}
		}
	}

	return data, nil
}

func (m *MongoLinkList) batchWrite(ctx context.Context, sources []*link.Link) ([]*link.Link, error) { // nolint unused
	docs := make([]interface{}, len(sources))

	// Create a new link
	for key := range sources {
		data, err := link.NewURL(sources[key].Url)
		if err != nil {
			return nil, err
		}

		docs[key] = data
	}

	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		switch err.(mongo.WriteException).WriteErrors[0].Code {
		case 11000:
			return nil, &link.NotUniqError{Link: sources[0], Err: fmt.Errorf("Duplicate URL: %s", sources[0].Url)}
		default:
			return nil, &link.NotFoundError{Link: sources[0], Err: fmt.Errorf("Failed marsharing link: %s", sources[0].Url)}
		}
	}

	return sources, nil
}

// setConfig - set configuration
func (m *MongoLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MONGODB_URI", "mongodb://localhost:27017/shortlink") // MongoDB URI
	viper.SetDefault("STORE_MONGODB_MODE_WRITE", storeOptions.MODE_SINGLE_WRITE) // mode write to store

	m.config = MongoConfig{
		URI:  viper.GetString("STORE_MONGODB_URI"),
		mode: viper.GetInt("STORE_MONGODB_MODE_WRITE"),
	}
}
