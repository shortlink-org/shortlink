package http

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"
	"google.golang.org/protobuf/encoding/protojson"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"

	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/controllers/cqrs"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/controllers/link"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/controllers/sitemap"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

// API HTTP-server
type Server struct {
	ctx    context.Context
	jsonpb protojson.MarshalOptions
}

// Config HTTP-server config
type Config struct {
	// Common
	Ctx    context.Context
	I18n   *message.Printer
	Log    logger.Logger
	Config *config.Config
	Http   http_server.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	RpcServer *rpc.Server

	// Infrastructure
	Link_rpc     link_rpc.LinkServiceClient
	Link_command link_cqrs.LinkCommandServiceClient
	Link_query   link_cqrs.LinkQueryServiceClient
	Sitemap_rpc  sitemap_rpc.SitemapServiceClient
}

type Controller struct {
	link.Controller
	cqrs.LinkCQRSController
	sitemap.SitemapController
}
