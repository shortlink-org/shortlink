package graphql

import (
	"context"
	"github.com/batazor/shortlink/pkg/api/graphql/resolver"
	"github.com/batazor/shortlink/pkg/api/graphql/schema"
	"github.com/batazor/shortlink/pkg/internal/store"
	"github.com/batazor/shortlink/pkg/logger"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
)

type API struct {
	store store.DB
	ctx   context.Context
}

func (api *API) GetHandler() (*relay.Handler, error) {
	s := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{Store: api.store})
	handler := relay.Handler{Schema: s}

	return &handler, nil
}

func (api *API) Run(ctx context.Context) error {
	var st store.Store

	api.ctx = ctx
	api.store = st.Use()

	log := logger.GetLogger(ctx)
	log.Info("Run GraphQL API")

	handler, err := api.GetHandler()
	if err != nil {
		return err
	}

	http.Handle("/api/query", handler)
	err = http.ListenAndServe(":7070", nil)

	return err
}
