package api

import "context"

// API - general describe of API
type API interface {
	Run(ctx context.Context) error
}
