package store

import (
	"context"
	"fmt"
	"time"

	"github.com/batazor/shortlink/pkg/link"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoLinkList implementation of store interface
type MongoLinkList struct { // nolint unused
	client *mongo.Client
}

// Init ...
func (m *MongoLinkList) Init() error {
	var err error

	// Connect to MongoDB
	m.client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
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

// Get ...
func (m *MongoLinkList) Get(id string) (*link.Link, error) {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	val := collection.FindOne(ctx, bson.D{primitive.E{Key: "hash", Value: id}})

	if val.Err() != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	if err := val.Decode(&response); err != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	if response.URL == "" {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// Add ...
func (m *MongoLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.URL), []byte("secret"))
	data.Hash = hash[:7]

	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed marsharing link: %s", data.URL)}
	}

	return &data, nil
}

// List ...
func (m *MongoLinkList) List() ([]*link.Link, error) {
	panic("implement me")
}

// Update ...
func (m *MongoLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (m *MongoLinkList) Delete(id string) error {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{primitive.E{Key: "hash", Value: id}})
	if err != nil {
		return &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}
