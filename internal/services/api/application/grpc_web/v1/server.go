package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/text/message"
	"google.golang.org/grpc/status"

	"github.com/batazor/shortlink/internal/pkg/logger"
	link_cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	"github.com/batazor/shortlink/pkg/http/server"
	"github.com/batazor/shortlink/pkg/rpc"
)

// API ...
type API struct {
	LinkServiceServer

	http *http.Server
	RPC  *rpc.RPCServer
}

// Run HTTP-server
func (api *API) Run(
	ctx context.Context,
	i18n *message.Printer,
	config http_server.Config,
	log logger.Logger,

	// Delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) error {
	// Rug gRPC
	RegisterLinkServiceServer(api.RPC.Server, api)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(
		// Register custom error handler
		runtime.WithErrorHandler(api.CustomHTTPError),
	)

	// DefaultContextTimeout is used for gRPC call context.WithTimeout whenever a Grpc-Timeout inbound
	// header isn't present. If the value is 0 the sent `context` will not have a timeout.
	runtime.DefaultContextTimeout = config.Timeout

	err := RegisterLinkServiceHandlerServer(ctx, mux, api)
	if err != nil {
		return err
	}

	api.http = http_server.New(ctx, mux, config)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Info(fmt.Sprintf("API run on port %d", config.Port))
	err = api.http.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// Close ...
func (api *API) Close() error {
	if err := api.http.Close(); err != nil {
		return err
	}

	return nil
}

func (api *API) CustomHTTPError(_ context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s := status.Convert(err)
	pb := s.Proto()

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaler.ContentType(pb)
	w.Header().Set("Content-Type", contentType)

	w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))

	jErr := json.NewEncoder(w).Encode(customError{
		Error: status.Convert(err).Message(),
	})

	if jErr != nil {
		_, _ = w.Write([]byte(fallback)) // #nosec
	}
}
