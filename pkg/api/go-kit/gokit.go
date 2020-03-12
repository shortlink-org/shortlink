package gokit

import (
	"context"
	"encoding/json"
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

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
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
	return func(_ context.Context, request interface{}) (interface{}, error) {
		//req := request.(link.Link)
		//
		//responseCh := make(chan interface{})
		//
		//// TODO: send []byte format
		//go notify.Publish(api_type.METHOD_GET, req.Hash, responseCh, "RESPONSE_STORE_GET")
		//
		//c := <-responseCh
		//switch r := c.(type) {
		//case nil:
		//	return nil, fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		//case notify.Response:
		//	return r.Payload.(*link.Link), nil
		//}

		return nil, nil
	}
}

func (api API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer opentracing.Tracer) error {
	api.ctx = ctx

	svc := linkService{}

	log.Info("Run go-kit API")

	linkAddHandler := httptransport.NewServer(
		makeAddLinkEndpoint(svc),
		decodeLinkRequest,
		encodeResponse,
	)

	linkGetListHandler := httptransport.NewServer(
		makeGetLinkEndpoint(svc),
		decodeLinkRequest,
		encodeResponse,
	)

	// set-up router and initialize http endpoints
	r := mux.NewRouter()

	r.Methods("GET").Path("/links").Handler(linkGetListHandler)
	//r.Methods("GET").Path("/:id").Handler(linkAddHandler)
	r.Methods("POST").Path("/").Handler(linkAddHandler)
	//r.Methods("DELETE").Path("/").Handler(linkAddHandler)

	http.Handle("/", linkAddHandler)

	log.Info(fmt.Sprintf("Run on port %d", config.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

	return nil
}

func decodeLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request link.Link
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
