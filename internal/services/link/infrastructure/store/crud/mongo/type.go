package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
)

// Config - config
type Config struct {
	job  *batch.Batch
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client *mongo.Client
	config Config
}
