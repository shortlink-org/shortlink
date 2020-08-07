package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/batazor/shortlink/internal/batch"
)

// MongoConfig ...
type MongoConfig struct { // nolint unused
	URI  string
	mode int
	job  *batch.Config
}

// MongoLinkList implementation of store interface
type MongoLinkList struct { // nolint unused
	client *mongo.Client
	config MongoConfig
}
