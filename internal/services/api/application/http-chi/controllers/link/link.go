package link_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/application/http-chi/helpers"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	v12 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
)

var (
	jsonpb protojson.MarshalOptions
)

// Routes creates a REST router
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/link/{hash}", Get)
	r.Get("/links", List)
	r.Post("/link", Add)
	r.Delete("/link/{hash}", Delete)

	return r
}

// Add ...
func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request v1.Link
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	newLink := &v1.Link{
		Url:      request.Url,
		Describe: request.Describe,
	}
	var response *v12.AddResponse

	responseCh := make(chan interface{})

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	// TODO: send []byte format
	go notify.Publish(r.Context(), api_type.METHOD_ADD, newLink, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_RPC_ADD"})

	c := <-responseCh
	switch resp := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
	case notify.Response:
		err = resp.Error
		if err == nil {
			response = resp.Payload.(*v12.AddResponse) // nolint errcheck
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint errcheck
}

// Get ...
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	var (
		response *v12.GetResponse
		err      error
	)

	responseCh := make(chan interface{})

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	go notify.Publish(r.Context(), api_type.METHOD_GET, hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_RPC_GET"})

	c := <-responseCh
	switch resp := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
	case notify.Response:
		err = resp.Error
		if err == nil {
			response = resp.Payload.(*v12.GetResponse) // nolint errcheck
		}
	}

	var errorLink *v1.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// List ...
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Get filter
	filter := r.URL.Query().Get("filter")

	var (
		response *v12.ListResponse
		err      error
	)

	responseCh := make(chan interface{})

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	go notify.Publish(r.Context(), api_type.METHOD_LIST, filter, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_RPC_LIST"})

	c := <-responseCh
	switch resp := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_LIST")
	case notify.Response:
		err = resp.Error
		if err == nil {
			response = resp.Payload.(*v12.ListResponse) // nolint errcheck
		}
	}

	var errorLink *v1.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := jsonpb.Marshal(response.GetLinks())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	responseCh := make(chan interface{})

	go notify.Publish(r.Context(), api_type.METHOD_DELETE, hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_RPC_DELETE"})

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	var err error
	c := <-responseCh
	switch resp := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
	case notify.Response:
		err = resp.Error
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}
