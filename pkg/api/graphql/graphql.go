package graphql

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/markbates/pkger"
	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/pkg/api/graphql/resolver"
	api_type "github.com/batazor/shortlink/pkg/api/type"
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
			file, err := pkger.Open(path)
			if err != nil {
				return err
			}

			f := make([]byte, info.Size())
			_, err = file.Read(f)

			// Add a newline if the file does not end in a newline.
			if len(f) > 0 && f[len(f)-1] != '\n' {
				if errWriteByte := buf.WriteByte('\n'); err != nil {
					panic(errWriteByte)
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
func (api *API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer opentracing.Tracer) error { // nolint unparam
	api.ctx = ctx

	log.Info("Run GraphQL API")

	handler := api.GetHandler()

	http.Handle("/api/query", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

	return err
}
