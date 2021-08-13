package api_application

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/services/api/application/cloudevents"
	gokit "github.com/batazor/shortlink/internal/services/api/application/go-kit"
	"github.com/batazor/shortlink/internal/services/api/application/graphql"
	grpcweb "github.com/batazor/shortlink/internal/services/api/application/grpc-web"
	http_chi "github.com/batazor/shortlink/internal/services/api/application/http-chi"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	link_cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	"github.com/batazor/shortlink/pkg/rpc"
)

// runAPIServer - start HTTP-server
func RunAPIServer(
	ctx context.Context,
	log logger.Logger,
	tracer *opentracing.Tracer,
	rpcServer *rpc.RPCServer,

	// delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) (*API, error) {
	var server API

	viper.SetDefault("API_TYPE", "http-chi") // Select: http-chi, gRPC-web, graphql, cloudevents, go-kit
	viper.SetDefault("API_PORT", 7070)       // API port
	viper.SetDefault("API_TIMEOUT", 60)      // Request Timeout (seconds)

	config := api_type.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT") * time.Second, // nolint durationcheck
	}

	serverType := viper.GetString("API_TYPE")

	switch serverType {
	case "http-chi":
		server = &http_chi.API{}
	case "go-kit":
		server = &gokit.API{}
	case "gRPC-web":
		server = &grpcweb.API{
			RPC: rpcServer,
		}
	case "graphql":
		server = &graphql.API{}
	case "cloudevents":
		server = &cloudevents.API{}
	default:
		server = &http_chi.API{}
	}

	if err := server.Run(ctx, config, log, tracer, link_rpc, link_command, link_query, sitemap_rpc); err != nil {
		return nil, err
	}

	return &server, nil
}
