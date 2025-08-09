package server

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelgraphql"
	"go.opentelemetry.io/otel/trace"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"

	"github.com/shortlink-org/shortlink/boundaries/api/api-gateway/gateways/graphql/infrastructure/server/resolver"
	"github.com/shortlink-org/shortlink/pkg/db"
	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

//go:embed schema/*.graphqls
var schema embed.FS //nolint:unused

// API ...
type API struct {
	store db.DB
	ctx   context.Context

	// delivery
	linkServiceClient link_rpc.LinkServiceClient
}

// GetHandler ...
func (api *API) GetHandler(traceProvider trace.TracerProvider) *relay.Handler {
	// tracing
	tracer := otelgraphql.NewTracer(otelgraphql.WithTracerProvider(traceProvider))

	buf := bytes.Buffer{}

	err := fs.WalkDir(schema, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".graphqls" {
			fileData, errReadFile := fs.ReadFile(schema, path)
			if errReadFile != nil {
				return fmt.Errorf("failed to read file: %w", errReadFile)
			}

			buf.Write(fileData)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	s := graphql.MustParseSchema(buf.String(), &resolver.Resolver{
		Store:             api.store,
		LinkServiceClient: api.linkServiceClient,
	}, graphql.Tracer(tracer))
	handler := relay.Handler{Schema: s}

	return &handler
}

// Run ...
func (api *API) Run(
	ctx context.Context,
	config http_server.Config,
	log logger.Logger,
	tracer trace.TracerProvider,

	// delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) error {

	api.ctx = ctx
	api.linkServiceClient = link_rpc

	handler := api.GetHandler(tracer)

	path := fmt.Sprintf("%s/query", viper.GetString("BASE_PATH"))
	log.Info("Run GraphQL API", field.Fields{"base_path": path})

	http.Handle(path, http.TimeoutHandler(handler, config.Timeout, http_server.TimeoutMessage))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

	return err
}
