package mongo

import (
	"go.mongodb.org/mongo-driver/v2/mongo"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/pkg/batch"
)

// Config - config
type Config struct {
	job  *batch.Batch[*v1.Link]
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client *mongo.Client
	config Config
}
