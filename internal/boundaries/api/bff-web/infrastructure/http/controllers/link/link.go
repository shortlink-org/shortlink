package link

import (
	"errors"
	"net/http"

	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/internal/boundaries/api/bff-web/infrastructure/http/api"
	v1 "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/link/v1"
)

var jsonpb protojson.MarshalOptions

type LinkController struct {
	LinkServiceClient link_rpc.LinkServiceClient
}

// AddLink - add link
func (c *LinkController) AddLink(w http.ResponseWriter, r *http.Request) {
	// Parse request
	var request v1.Link
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	// Save link
	response, err := c.LinkServiceClient.Add(r.Context(), &link_rpc.AddRequest{Link: &request})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint:errcheck
}

// UpdateLink - update link
func (c *LinkController) UpdateLinks(w http.ResponseWriter, r *http.Request) {
	// Parse request
	var request v1.Link
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	// Update link
	response, err := c.LinkServiceClient.Update(r.Context(), &link_rpc.UpdateRequest{Link: &request})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint:errcheck
}

// GetLink - get link by hash
func (c *LinkController) GetLink(w http.ResponseWriter, r *http.Request, hash api.HashParam) {
	response, err := c.LinkServiceClient.Get(r.Context(), &link_rpc.GetRequest{Hash: hash})
	if err != nil {
		var errorLink *v1.NotFoundError

		if errors.Is(err, errorLink) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

			return
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.GetLink())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint:errcheck
}

// GetLinks - get links
func (c *LinkController) GetLinks(w http.ResponseWriter, r *http.Request, params api.GetLinksParams) {
	// Get filter
	filter, err := json.Marshal(params.Filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := c.LinkServiceClient.List(r.Context(), &link_rpc.ListRequest{Filter: string(filter)})
	if err != nil {
		var errorLink *v1.NotFoundError

		if errors.Is(err, errorLink) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

			return
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.GetLinks())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint:errcheck
}

// DeleteLink - delete link
func (c *LinkController) DeleteLink(w http.ResponseWriter, r *http.Request, hash api.HashParam) {
	_, err := c.LinkServiceClient.Delete(r.Context(), &link_rpc.DeleteRequest{Hash: hash})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint:errcheck
}
