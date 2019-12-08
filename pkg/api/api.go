package api

import (
	"context"

	"github.com/batazor/shortlink/internal/store"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// API - general describe of API
type API interface { // nolint unused
	Run(ctx context.Context, st store.DB, config api_type.Config) error
}

type Server struct{}
