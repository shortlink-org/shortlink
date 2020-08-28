package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/batazor/shortlink/internal/batch"
)

// Config ...
type Config struct { // nolint unused
	URI  string
	mode int
	job  *batch.Config
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *mongo.Client
	config Config
}
