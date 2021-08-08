package link_api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/application/http-chi/helpers"
	"github.com/batazor/shortlink/internal/services/api/domain"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
)

// GetByCQRS ...
func GetByCQRS(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	var (
		response *v1.Link
		err      error
	)

	responseCh := make(chan interface{})

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	go notify.Publish(r.Context(), api_domain.METHOD_CQRS_GET, hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_RPC_CQRS_GET"})

	c := <-responseCh
	switch resp := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_CQRS_GET")
	case notify.Response:
		err = resp.Error
		if err == nil {
			response = resp.Payload.(*v1.Link) // nolint errcheck
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

	res, err := jsonpb.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}
