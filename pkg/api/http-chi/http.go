package httpchi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/batazor/shortlink/internal/notify"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

// Routes creates a REST router
func (api *API) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", api.Add)
	r.Get("/links", api.List)
	r.Get("/{hash}", api.Get)
	r.Delete("/", api.Delete)

	return r
}

// Add ...
func (api *API) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request addRequest
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	newLink := &link.Link{
		Url:      request.URL,
		Describe: request.Describe,
	}

	responseCh := make(chan interface{})

	// TODO: send []byte format
	go notify.Publish(api_type.METHOD_ADD, *newLink, responseCh, "RESPONSE_STORE_ADD")

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
	case notify.Response:
		err = r.Error
		if err == nil {
			newLink = r.Payload.(*link.Link) // nolint errcheck
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := json.Marshal(newLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint errcheck
}

// Get ...
func (api *API) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	// Parse request
	var request = &getRequest{
		Hash: hash,
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
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	responseLink = ResponseLink{
		response,
	}
	res, err := json.Marshal(responseLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// List ...
func (api *API) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

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

	var errorLink *link.NotFoundError
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

	for l := range response {
		responseLink = append(responseLink, ResponseLink{response[l]})
	}

	res, err := json.Marshal(responseLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// Delete ...
func (api *API) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	b, err := ioutil.ReadAll(r.Body)
	defer func() {
		_ = r.Body.Close() // nolint errcheck
	}()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	var request deleteRequest
	err = json.Unmarshal(b, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	responseCh := make(chan interface{})

	go notify.Publish(api_type.METHOD_DELETE, request.Hash, responseCh, "RESPONSE_STORE_DELETE")

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
	case notify.Response:
		err = r.Error
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}
