package api

import (
	"encoding/json"
	"errors"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

// Routes creates a REST router
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", Add)
	r.Get("/{url}", Get)
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

	newLink, err := link.NewURL(request.Url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	linkList, err := link.Init()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	err = linkList.Add(newLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"url": "` + newLink.Url + `"}`))
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var url = chi.URLParam(r, "url")

	// Parse request
	var request = &getRequest{
		Url: url,
	}

	getLink, err := link.NewURL(request.Url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	linkList, err := link.Init()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	response, err := linkList.Get(getLink)
	var errorLink *link.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	if response.Url == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"url": "` + response.Url + `"}`))
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

	newLink, err := link.NewURL(request.Url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	linkList, err := link.Init()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	err = linkList.Delete(newLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{}`))
}
