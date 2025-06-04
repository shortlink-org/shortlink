package link

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/api"
)

// UpdateLinks - update link
func (c *Controller) UpdateLinks(w http.ResponseWriter, r *http.Request) {
	// Parse request
	var request api.UpdateLinkRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck

		return
	}

	// Create link
	// link, err := v1.NewLinkBuilder().
	// 	SetURL(request.Link.Url.String()).
	// 	SetDescribe(request.Link.Describe).
	// 	Build()
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 	return
	// }

	// // Update link
	// _, err = c.linkServiceClient.Update(r.Context(), &link_rpc.UpdateRequest{Link: link})
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 	return
	// }

	count := 0
	response := &api.UpdateLinks200Response{
		UpdatedCount: &count,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		c.log.Error(err.Error())
	}
}
