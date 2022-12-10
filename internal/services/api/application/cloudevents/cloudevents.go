package cloudevents

import (
	"context"
	"fmt"

	http_server "github.com/batazor/shortlink/pkg/http/server"
	"golang.org/x/text/message"

	"github.com/batazor/shortlink/internal/pkg/logger"
	link_cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Receive ...
func Receive(_ context.Context, event cloudevents.Event) error {
	// do something with event.Context and event.Data (via event.DataAs(foo)
	data := &Example{}

	if err := event.DataAs(data); err != nil {
		fmt.Printf("Got Data Error: %s\n", err.Error())
	}

	fmt.Printf("Got Data: %+v\n", data)

	fmt.Printf("----------------------------\n")

	return nil
}

// Run ...
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

	api.ctx = ctx

	log.Info("Run Cloud-Events API")

	// New endpoint (HTTP)
	cloudevents.WithPort(config.Port)
	cloudevents.WithPath("/")

	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		return err
	}

	if err = c.StartReceiver(ctx, Receive); err != nil {
		return err
	}

	return nil
}
