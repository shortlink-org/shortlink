package api

import (
	"context"
	"github.com/batazor/shortlink/internal/store"
)

// API - general describe of API
type API interface { // nolint unused
	Run(ctx context.Context, st store.DB, config Config) error
}

// Config - base configuration for API
type Config struct { // nolint unused
	Port int
}
