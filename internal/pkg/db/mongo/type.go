package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Config - config
type Config struct {
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client *mongo.Client
	config Config
}
