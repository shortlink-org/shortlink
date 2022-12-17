package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cloudevents/sdk-go/observability/opencensus/v2/client"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"knative.dev/pkg/tracing"
	"knative.dev/pkg/tracing/config"
)

type Service struct {
	elastic *elasticsearch.Client
}

func main() {
	ctx := context.Background()

	// Init elastic client
	elasticClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://shortlink-master.elasticsearch:9200",
		},
	})

	service := &Service{
		elastic: elasticClient,
	}

	c, err := client.NewClientHTTP(
		[]cehttp.Option{cehttp.WithMiddleware(healthzMiddleware)}, nil,
	)
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	conf, err := config.JSONToTracingConfig(os.Getenv("K_CONFIG_TRACING"))
	if err != nil {
		log.Fatal("Failed to read tracing config, using the on-op default: %v", err)
	}

	tracer, err := tracing.SetupPublishingWithStaticConfig(zap.L().Sugar(), "", conf)
	if err != nil {
		log.Fatal("Failed to initialize tracing: %v ", err)
	}
	defer func(ctx context.Context) {
		_ = tracer.Shutdown(ctx)
	}(ctx)

	err = c.StartReceiver(ctx, service.display)
	if err != nil {
		log.Fatal("Error during receiver's runtime: ", err)
	}
}

// display prints the given Event in a human-readable format.
func (s *Service) display(event cloudevents.Event) {
	fmt.Printf("☁️  cloudevents.Event\n%s", event)

	// send event to elastic
	_, err := s.elastic.Index("shortlink.event.link.new", nil)
	if err != nil {
		// TODO: use logger
		log.Fatal("Error indexing document: %s", err)
	}
}

// HTTP path of the health endpoint used for probing the service.
const healthzPath = "/healthz"

// healthzMiddleware is a cehttp.Middleware which exposes a health endpoint.
func healthzMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.RequestURI == healthzPath {
			w.WriteHeader(http.StatusNoContent)
		} else {
			next.ServeHTTP(w, req)
		}
	})
}
