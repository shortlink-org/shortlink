package httpchi

import (
	"context"
)

// API ...
type API struct { // nolint unused
	ctx context.Context
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
