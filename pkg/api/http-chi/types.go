package httpchi

import (
	"context"
	"github.com/batazor/shortlink/pkg/internal/store"
)

// API ...
type API struct {
	store store.DB
	ctx   context.Context
}

// addRequest ...
type addRequest struct {
	URL      string
	Describe string
}

// getRequest ...
type getRequest struct {
	Hash     string
	Describe string
}

// deleteRequest ...
type deleteRequest struct {
	Hash string
}
