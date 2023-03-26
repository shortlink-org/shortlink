package graphql

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/uptrace/opentelemetry-go-extra/otelgraphql"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/services/api/application/graphql/resolver"
	link_cqrs "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
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
func (api *API) GetHandler(traceProvider *trace.TracerProvider) *relay.Handler {
	// tracing
	tracer := otelgraphql.NewTracer(otelgraphql.WithTracerProvider(*traceProvider))

	buf := bytes.Buffer{}

	err := filepath.Walk("./internal/services/api/application/graphql/schema", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() { // nolint:nestif
			file, errReadFile := os.ReadFile(filepath.Clean(path))
			if errReadFile != nil {
				return errReadFile
			}

			// Add a newline if the file does not end in a newline.
			if len(file) > 0 && file[len(file)-1] != '\n' {
				if errWriteByte := buf.WriteByte('\n'); errWriteByte != nil {
					panic(errWriteByte)
				}
			}

			_, errWrite := buf.Write(file)
			if errWrite != nil {
				panic(errWrite)
			}
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
	i18n *message.Printer,
	config http_server.Config,
	log logger.Logger,
	tracer *trace.TracerProvider,

	// delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) error {

	api.ctx = ctx
	api.linkServiceClient = link_rpc

	log.Info("Run GraphQL API")

	handler := api.GetHandler(tracer)

	http.Handle("/api/query", http.TimeoutHandler(handler, config.Timeout, http_server.TimeoutMessage))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

	return err
}
