package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Config ...
type Config struct { // nolint unused
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *mongo.Client
	config Config
}
