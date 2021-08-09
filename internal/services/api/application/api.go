/*
API
*/

package api_application

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/pkg/logger"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	v1 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
)

// API - general describe of API
type API interface { // nolint unused
	Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer *opentracing.Tracer) error
}

type Server struct {
	// Delivery
	MetadataClient metadata_rpc.MetadataServiceClient

	LinkServiceClient        link_rpc.LinkServiceClient
	LinkCommandServiceClient v1.LinkCommandServiceClient
	LinkQueryServiceClient   v1.LinkQueryServiceClient
}
