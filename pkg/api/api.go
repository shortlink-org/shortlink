package api

import (
	"context"
	"github.com/batazor/shortlink/pkg/store"
)

// API - general describe of API
type API interface {
	Run(ctx context.Context, st store.DB) error
}
