package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig ...
type MongoConfig struct { // nolint unused
	URI string
}

// MongoLinkList implementation of store interface
type MongoLinkList struct { // nolint unused
	ctx context.Context

	client *mongo.Client
	config MongoConfig
}

// Init ...
func (m *MongoLinkList) Init(ctx context.Context) error {
	var err error

	// Set context
	m.ctx = ctx

	// Set configuration
	m.setConfig()

	// Connect to MongoDB
	m.client, err = mongo.NewClient(options.Client().ApplyURI(m.config.URI))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = m.client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return nil
}

// Close ...
func (m *MongoLinkList) Close() error {
	return m.client.Disconnect(context.Background())
}

// Migrate ...
func (m *MongoLinkList) migrate() error { // nolint unused
	return nil
}

// Add ...
func (m *MongoLinkList) Add(source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, &data)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed marsharing link: %s", data.Url)}
	}

	return data, nil
}

// Get ...
func (m *MongoLinkList) Get(id string) (*link.Link, error) {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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
func (m *MongoLinkList) List(filter *query.Filter) ([]*link.Link, error) { // nolint unused
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(ctx, bson.D{})
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
func (m *MongoLinkList) Update(data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (m *MongoLinkList) Delete(id string) error {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{primitive.E{Key: "hash", Value: id}})
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}

// setConfig - set configuration
func (m *MongoLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MONGODB_URI", "mongodb://localhost:27017")
	m.config = MongoConfig{
		URI: viper.GetString("STORE_MONGODB_URI"),
	}
}
