package link

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	v1 "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/infrastructure/rpc/link/v1"

	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/api"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/controllers/link/dto"
)

// AddLink - add link
func (c *Controller) AddLink(w http.ResponseWriter, r *http.Request) {
	// Parse request
	var request api.AddLink
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck

		return
	}

	// Save link
	result, err := c.linkServiceClient.Add(r.Context(), &v1.AddRequest{Link: dto.MakeAddLinkRequest(request)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck

		return
	}

	response := &api.Link{
		Url:       result.GetLink().GetUrl(),
		Hash:      result.GetLink().GetHash(),
		Describe:  result.GetLink().GetDescribe(),
		CreatedAt: result.GetLink().GetCreatedAt().AsTime(),
		UpdatedAt: result.GetLink().GetUpdatedAt().AsTime(),
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		c.log.Error(err.Error())
	}
}
