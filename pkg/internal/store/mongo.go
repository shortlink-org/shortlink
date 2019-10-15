package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/batazor/shortlink/pkg/internal/link"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoLinkList struct {
	client *mongo.Client
}

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

func (m *MongoLinkList) Get(id string) (*link.Link, error) {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	val := collection.FindOne(ctx, bson.D{{"hash", id}})

	if val.Err() != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	var response link.Link

	if err := val.Decode(&response); err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Failed parse link: %s", id))}
	}

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	return &response, nil
}

func (m *MongoLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: errors.New(fmt.Sprintf("Failed marsharing link: %s", data.Url))}
	}

	return &data, nil
}

func (m *MongoLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

func (m *MongoLinkList) Delete(id string) error {
	collection := m.client.Database("shortlink").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{{"hash", id}})
	if err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Failed save link: %s", id))}
	}

	return nil
}
