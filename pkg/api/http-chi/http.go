package httpchi

import (
	"encoding/json"
	"errors"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

// Routes creates a REST router
func (api *API) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", api.Add)
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
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	newLink := &link.Link{
		URL:      request.URL,
		Describe: request.Describe,
	}

	newLink, err = api.store.Add(*newLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	res, err := json.Marshal(newLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

// Get ...
func (api *API) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")

	// Parse request
	var request = &getRequest{
		Hash: hash,
	}

	response, err := api.store.Get(request.Hash)
	var errorLink *link.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

// Delete ...
func (api *API) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	b, err := ioutil.ReadAll(r.Body)
	defer func() {
		_ = r.Body.Close()
	}()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	var request deleteRequest
	err = json.Unmarshal(b, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	err = api.store.Delete(request.Hash)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}
