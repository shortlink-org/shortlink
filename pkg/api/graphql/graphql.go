package graphql

import (
	"context"
	"github.com/batazor/shortlink/pkg/api/graphql/resolver"
	"github.com/batazor/shortlink/pkg/api/graphql/schema"
	"github.com/batazor/shortlink/pkg/internal/store"
	"github.com/batazor/shortlink/pkg/logger"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

type API struct {
	store store.DB
	ctx   context.Context
}

func (api *API) Run(ctx context.Context) error {
	var st store.Store

	api.ctx = ctx
	api.store = st.Use()

	logger := logger.GetLogger(ctx)
	logger.Info("Run GraphQL API")

	s := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{Store: api.store})
	http.Handle("/api/query", &relay.Handler{Schema: s})
	log.Fatal(http.ListenAndServe(":7070", nil))

	return nil
}
