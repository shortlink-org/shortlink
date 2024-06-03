package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres/schema/crud"
	"github.com/shortlink-org/shortlink/pkg/batch"
)

// Config - config
type Config struct {
	job  *batch.Batch
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool
	query  *crud.Queries

	config Config
}

// ExampleJsonLink - example json link
// NOTE: we use this structure only for demonstration work with JSONb type in Postgres
type ExampleJsonLink struct {
	URI      string `json:"uri,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Describe string `json:"describe,omitempty"`
}

func NewExampleJsonLink(link domain.Link) *ExampleJsonLink {
	return &ExampleJsonLink{
		URI:      link.GetUrl().String(),
		Hash:     link.GetHash(),
		Describe: link.GetDescribe(),
	}
}
