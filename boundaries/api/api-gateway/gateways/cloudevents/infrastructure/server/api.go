package server

import (
	"context"

	otelObs "github.com/cloudevents/sdk-go/observability/opentelemetry/v2/client"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"

	"github.com/shortlink-org/shortlink/boundaries/api/api-gateway/gateways/cloudevents/infrastructure/server/handlers"
	http_client "github.com/shortlink-org/shortlink/pkg/http/client"
	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

// API ...
type API struct {
	ctx context.Context
}

// Run ...
func (api *API) Run(
	ctx context.Context,
	config http_server.Config,
	log logger.Logger,
	tracer trace.TracerProvider,

	// Delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) error {

	api.ctx = ctx

	log.Info("Run Cloud-Events API")

	// New endpoint (HTTP)
	p, err := cloudevents.NewHTTP(
		cloudevents.WithRoundTripper(otelhttp.NewTransport(http_client.New().Transport)),
		cloudevents.WithPort(config.Port),
		cloudevents.WithPath(viper.GetString("BASE_PATH")),
	)

	c, err := cloudevents.NewClient(p, client.WithObservabilityService(otelObs.NewOTelObservabilityService()))
	if err != nil {
		return err
	}

	if err = c.StartReceiver(ctx, handlers.Receive); err != nil {
		return err
	}

	return nil
}
