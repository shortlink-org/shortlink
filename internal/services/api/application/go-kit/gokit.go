package gokit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/notify"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	additionalMiddleware "github.com/batazor/shortlink/pkg/http/middleware"
)

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeAddLinkEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*v1.Link)
		if !ok {
			return nil, nil
		}

		responseCh := make(chan interface{})

		// TODO: send []byte format
		go notify.Publish(ctx, v1.METHOD_ADD, req, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_ADD"})

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			return nil, fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		case notify.Response:
			return r.Payload.(*v1.Link), nil
		}

		return nil, nil
	}
}

func makeGetLinkEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		vars := mux.Vars(r.(*http.Request))
		if vars["id"] == "" {
			return nil, errors.New(`{"error": "need set hash URL"}`)
		}

		// Parse request
		var request = &getRequest{
			Hash: vars["id"],
		}

		var (
			response     *v1.Link
			responseLink ResponseLink // for custom JSON parsing
			err          error
		)

		responseCh := make(chan interface{})

		go notify.Publish(ctx, v1.METHOD_GET, request.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_GET"})

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			err = fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		case notify.Response:
			err = r.Error
			if err == nil {
				response = r.Payload.(*v1.Link) // nolint errcheck
			}
		}

		var errorLink *v1.NotFoundError
		if errors.As(err, &errorLink) {
			return nil, errors.New(`{"error": "` + err.Error() + `"}`)
		}
		if err != nil {
			return nil, errors.New(`{"error": "` + err.Error() + `"}`)
		}

		responseLink = ResponseLink{
			response,
		}

		return responseLink, nil
	}
}

func makeGetListLinkEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			response     []*v1.Link
			responseLink []ResponseLink // for custom JSON parsing
			err          error
		)

		responseCh := make(chan interface{})

		go notify.Publish(ctx, v1.METHOD_LIST, nil, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_LIST"})

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			err = fmt.Errorf("Not found subscribe to event %s", "METHOD_LIST")
			return nil, err
		case notify.Response:
			err = r.Error
			if err == nil {
				response = r.Payload.([]*v1.Link) // nolint errcheck
			}
		}

		for l := range response {
			responseLink = append(responseLink, ResponseLink{response[l]})
		}

		return responseLink, nil
	}
}

func makeDeleteLinkEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		var err error
		req, ok := r.(*v1.Link)
		if !ok {
			return nil, nil
		}

		responseCh := make(chan interface{})

		go notify.Publish(ctx, v1.METHOD_DELETE, req.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_DELETE"})

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			err = fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		case notify.Response:
			err = r.Error
		}

		if err != nil {
			return nil, errors.New(`{"error": "` + err.Error() + `"}`)
		}

		return `{}`, nil
	}
}

func (api API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer *opentracing.Tracer) error { // nolint unparam
	log.Info("Run go-kit API")

	linkAddHandler := httptransport.NewServer(
		makeAddLinkEndpoint(),
		decodeAddListRequest,
		encodeResponse,
	)

	linkGetByIdHandler := httptransport.NewServer(
		makeGetLinkEndpoint(),
		decodeGetListRequest,
		encodeResponse,
	)

	linkGetListHandler := httptransport.NewServer(
		makeGetListLinkEndpoint(),
		decodeGetListRequest,
		encodeResponse,
	)

	linkDeleteHandler := httptransport.NewServer(
		makeDeleteLinkEndpoint(),
		decodeAddListRequest,
		encodeResponse,
	)

	// set-up router and initialize http endpoints
	r := mux.NewRouter()

	r.Methods("GET").Path("/api/links").Handler(linkGetListHandler)
	r.Methods("GET").Path("/api/{id}").Handler(linkGetByIdHandler)
	r.Methods("POST").Path("/api").Handler(linkAddHandler)
	r.Methods("DELETE").Path("/api").Handler(linkDeleteHandler)

	// Additional middleware
	r.Use(additionalMiddleware.NewTracing(tracer))
	r.Use(additionalMiddleware.Logger(log))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(config.Timeout))

	log.Info(fmt.Sprintf("Run on port %d", config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)

	return err
}

func decodeAddListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request v1.Link
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func decodeGetListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
