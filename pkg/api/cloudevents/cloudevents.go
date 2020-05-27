package cloudevents

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/logger"
	api_type "github.com/batazor/shortlink/pkg/api/type"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Receive ...
func Receive(ctx context.Context, event cloudevents.Event) error { // nolint unused
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
func (api *API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer opentracing.Tracer) error { // nolint unparam
	api.ctx = ctx

	log.Info("Run Cloud-Events API")

	// New endpoint (HTTP)
	cloudevents.WithPort(config.Port)
	cloudevents.WithPath("/")

	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		return err
	}

	if err = c.StartReceiver(context.Background(), Receive); err != nil {
		return err
	}

	return nil
}
