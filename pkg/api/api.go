package api

import (
	"context"

	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// API - general describe of API
type API interface { // nolint unused
	Run(ctx context.Context, config api_type.Config) error
}

type Server struct{}
