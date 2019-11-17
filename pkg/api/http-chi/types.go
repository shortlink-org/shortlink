package httpchi

import (
	"context"
	"github.com/batazor/shortlink/pkg/store"
)

// API ...
type API struct { // nolint unused
	store store.DB
	ctx   context.Context
}

// addRequest ...
type addRequest struct { // nolint unused
	URL      string
	Describe string
}

// getRequest ...
type getRequest struct { // nolint unused
	Hash     string
	Describe string
}

// deleteRequest ...
type deleteRequest struct { // nolint unused
	Hash string
}
