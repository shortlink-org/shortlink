package gokit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/notify"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

type linkService struct {
	ctx context.Context
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeAddLinkEndpoint(svc linkService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(link.Link)

		responseCh := make(chan interface{})

		// TODO: send []byte format
		go notify.Publish(api_type.METHOD_ADD, req, responseCh, "RESPONSE_STORE_ADD")

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			return nil, fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		case notify.Response:
			return r.Payload.(*link.Link), nil
		}

		return nil, nil
	}
}

func makeGetLinkEndpoint(svc linkService) endpoint.Endpoint {
	return func(_ context.Context, r interface{}) (interface{}, error) {
		vars := mux.Vars(r.(*http.Request))
		if vars["id"] == "" {
			return nil, errors.New(`{"error": "need set hash URL"}`)
		}

		// Parse request
		var request = &getRequest{
			Hash: vars["id"],
		}

		var (
			response     *link.Link
			responseLink ResponseLink // for custom JSON parsing
			err          error
		)

		responseCh := make(chan interface{})

		go notify.Publish(api_type.METHOD_GET, request.Hash, responseCh, "RESPONSE_STORE_GET")

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			err = fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		case notify.Response:
			err = r.Error
			if err == nil {
				response = r.Payload.(*link.Link) // nolint errcheck
			}
		}

		var errorLink *link.NotFoundError
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

func makeGetListLinkEndpoint(svc linkService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var (
			response     []*link.Link
			responseLink []ResponseLink // for custom JSON parsing
			err          error
		)

		responseCh := make(chan interface{})

		go notify.Publish(api_type.METHOD_LIST, nil, responseCh, "RESPONSE_STORE_LIST")

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			err = fmt.Errorf("Not found subscribe to event %s", "METHOD_LIST")
		case notify.Response:
			err = r.Error
			if err == nil {
				response = r.Payload.([]*link.Link) // nolint errcheck
			}
		}

		for l := range response {
			responseLink = append(responseLink, ResponseLink{response[l]})
		}

		return responseLink, nil
	}
}

func makeDeleteLinkEndpoint(svc linkService) endpoint.Endpoint {
	return func(_ context.Context, r interface{}) (interface{}, error) {
		req := r.(link.Link)
		var err error

		responseCh := make(chan interface{})

		go notify.Publish(api_type.METHOD_DELETE, req.Hash, responseCh, "RESPONSE_STORE_DELETE")

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

func (api API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer opentracing.Tracer) error {
	api.ctx = ctx

	svc := linkService{}

	log.Info("Run go-kit API")

	linkAddHandler := httptransport.NewServer(
		makeAddLinkEndpoint(svc),
		decodeAddLinkRequest,
		encodeResponse,
	)

	linkGetByIdHandler := httptransport.NewServer(
		makeGetLinkEndpoint(svc),
		decodeGetLinkRequest,
		encodeResponse,
	)

	linkGetListHandler := httptransport.NewServer(
		makeGetListLinkEndpoint(svc),
		decodeGetLinkRequest,
		encodeResponse,
	)

	linkDeleteHandler := httptransport.NewServer(
		makeDeleteLinkEndpoint(svc),
		decodeAddLinkRequest,
		encodeResponse,
	)

	// set-up router and initialize http endpoints
	r := mux.NewRouter()

	r.Methods("GET").Path("/api/links").Handler(linkGetListHandler)
	r.Methods("GET").Path("/api/{id}").Handler(linkGetByIdHandler)
	r.Methods("POST").Path("/api").Handler(linkAddHandler)
	r.Methods("DELETE").Path("/api").Handler(linkDeleteHandler)

	log.Info(fmt.Sprintf("Run on port %d", config.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)

	return nil
}

func decodeAddLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request link.Link
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
