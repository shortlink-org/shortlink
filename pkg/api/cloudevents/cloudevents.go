package cloudevents

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/logger"
	api_type "github.com/batazor/shortlink/pkg/api/type"

	cloudevents "github.com/cloudevents/sdk-go"
)

// Receive ...
func Receive(ctx context.Context, event cloudevents.Event) error { // nolint unused
	// do something with event.Context and event.Data (via event.DataAs(foo)
	data := &Example{}

	if err := event.DataAs(data); err != nil {
		fmt.Printf("Got Data Error: %s\n", err.Error())
	}

	fmt.Printf("Got Data: %+v\n", data)

	fmt.Printf("Got Transport Context: %+v\n", cloudevents.HTTPTransportContextFrom(ctx))

	fmt.Printf("----------------------------\n")
	return nil
}

// Run ...
func (api *API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer opentracing.Tracer) error {
	api.ctx = ctx

	log.Info("Run Cloud-Events API")

	// New endpoint (HTTP)
	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(config.Port),
		cloudevents.WithPath("/"),
	)
	if err != nil {
		return err
	}

	c, err := cloudevents.NewClient(t)
	if err != nil {
		return err
	}

	if err = c.StartReceiver(context.Background(), Receive); err != nil {
		return err
	}

	return nil
}
