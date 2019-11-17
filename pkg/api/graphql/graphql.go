package graphql

import (
	"context"
	"github.com/batazor/shortlink/pkg/api/graphql/resolver"
	"github.com/batazor/shortlink/pkg/api/graphql/schema"
	"github.com/batazor/shortlink/pkg/logger"
	"github.com/batazor/shortlink/pkg/store"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
)

// API ...
type API struct { // nolint unused
	store store.DB
	ctx   context.Context
}

// GetHandler ...
func (api *API) GetHandler() *relay.Handler {
	s := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{Store: api.store})
	handler := relay.Handler{Schema: s}

	return &handler
}

// Run ...
func (api *API) Run(ctx context.Context, db store.DB) error {
	api.ctx = ctx
	api.store = db

	log := logger.GetLogger(ctx)
	log.Info("Run GraphQL API")

	handler := api.GetHandler()

	http.Handle("/api/query", handler)
	err := http.ListenAndServe(":7070", nil)

	return err
}
