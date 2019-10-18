package graphql

import (
	"context"
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

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func (api *API) Run(ctx context.Context) error {
	var st store.Store

	api.ctx = ctx
	api.store = st.Use()

	logger := logger.GetLogger(ctx)
	logger.Info("Run GraphQL API")

	s := `
                schema {
                        query: Query
                }
                type Query {
                        hello: String!
                }
        `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":7070", nil))

	return nil
}
