package cloudevents

import (
	"context"
	"github.com/batazor/shortlink/internal/store"
)

// API ...
type API struct { // nolint unused
	store store.DB
	ctx   context.Context
}
