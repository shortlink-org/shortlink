package api

import (
	"context"
	"github.com/batazor/shortlink/internal/store"
)

// API - general describe of API
type API interface {
	Run(ctx context.Context, st store.DB) error
}
