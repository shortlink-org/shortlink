package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
)

// Config ...
type Config struct {
	job  *batch.Config
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client *mongo.Client
	config Config
}
