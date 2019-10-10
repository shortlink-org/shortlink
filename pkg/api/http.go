package api

import (
	"encoding/json"
	"errors"
	"github.com/batazor/shortlink/pkg/internal/link"
	"github.com/batazor/shortlink/pkg/internal/store"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

var (
	//TODO: refactoring
	st store.Store
	s  store.DB
)

func init() {
	s = st.Use()
}

// Routes creates a REST router
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", Add)
	r.Get("/{hash}", Get)
	r.Get("/s/{hash}", Redirect)
	r.Delete("/", Delete)

	return r
}

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request addRequest
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	newLink := &link.Link{
		Url:      request.Url,
		Describe: request.Describe,
	}

	newLink, err = s.Add(*newLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	res, err := json.Marshal(newLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")

	// Parse request
	var request = &getRequest{
		Hash: hash,
	}

	response, err := s.Get(request.Hash)
	var errorLink *link.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	var request deleteRequest
	err = json.Unmarshal(b, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	err = s.Delete(request.Hash)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{}`))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")

	// Parse request
	var request = &getRequest{
		Hash: hash,
	}

	response, err := s.Get(request.Hash)
	var errorLink *link.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	http.Redirect(w, r, response.Url, http.StatusMovedPermanently)
}
