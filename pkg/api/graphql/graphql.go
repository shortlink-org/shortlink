package graphql

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/pkg/api"
	"github.com/batazor/shortlink/pkg/api/graphql/resolver"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/markbates/pkger"
)

// API ...
type API struct { // nolint unused
	store store.DB
	ctx   context.Context
}

// GetHandler ...
func (api *API) GetHandler() *relay.Handler {
	buf := bytes.Buffer{}

	err := pkger.Walk("/pkg/api/graphql/schema", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, _ := pkger.Open(path)
			f := make([]byte, info.Size())
			_, err := file.Read(f)

			// Add a newline if the file does not end in a newline.
			if len(f) > 0 && f[len(f)-1] != '\n' {
				if err := buf.WriteByte('\n'); err != nil {
					panic(err)
				}
			}

			if err != nil {
				panic(err)
			}

			if _, err := buf.Write(f); err != nil {
				panic(err)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	s := graphql.MustParseSchema(buf.String(), &resolver.Resolver{Store: api.store})
	handler := relay.Handler{Schema: s}

	return &handler
}

// Run ...
func (api *API) Run(ctx context.Context, db store.DB, config api.Config) error {
	api.ctx = ctx
	api.store = db

	log := logger.GetLogger(ctx)
	log.Info("Run GraphQL API")

	handler := api.GetHandler()

	http.Handle("/api/query", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

	return err
}
