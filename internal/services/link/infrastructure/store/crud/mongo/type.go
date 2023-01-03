package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
)

// Config ...
type Config struct {
	URI  string
	mode int
	job  *batch.Config
}

// Store implementation of db interface
type Store struct {
	client *mongo.Client
	config Config
}
