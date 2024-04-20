package link

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	v1 "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/infrastructure/rpc/link/v1"

	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/api"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/controllers/link/dto"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

type Controller struct {
	log logger.Logger

	linkServiceClient link_rpc.LinkServiceClient
}

// NewController - create new link controller
func NewController(log logger.Logger, linkServiceClient link_rpc.LinkServiceClient) Controller {
	return Controller{
		log: log,

		linkServiceClient: linkServiceClient,
	}
}

// ErrMessages - helper for create error messages
func ErrMessages(err error) *api.ErrorResponse {
	messages := []string{err.Error()}

	return &api.ErrorResponse{
		Messages: &messages,
	}
}

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

// GetLink - get link by hash
func (c *Controller) GetLink(w http.ResponseWriter, r *http.Request, hash string) {
	// result, err := c.linkServiceClient.Get(r.Context(), &link_rpc.GetRequest{Hash: hash})
	// if err != nil {
	// 	var errorLink *v1.NotFoundError
	//
	// 	if errors.Is(err, errorLink) {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 		return
	// 	}
	//
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 	return
	// }

	// response := &api.Link{
	// 	Url:       result.GetLink().GetUrl(),
	// 	Hash:      result.GetLink().GetHash(),
	// 	Describe:  result.GetLink().GetDescribe(),
	// 	CreatedAt: result.GetLink().GetCreatedAt().AsTime(),
	// 	UpdatedAt: result.GetLink().GetUpdatedAt().AsTime(),
	// }
	//
	// w.WriteHeader(http.StatusOK)
	// err = json.NewEncoder(w).Encode(response)
	// if err != nil {
	// 	c.log.Error(err.Error())
	// }
}

// GetLinks - get links
func (c *Controller) GetLinks(w http.ResponseWriter, r *http.Request, params api.GetLinksParams) {
	// TODO: add mapper for filter

	// request := &link_rpc.ListRequest{}
	//
	// if params.Cursor != nil {
	// 	request.Cursor = *params.Cursor
	// }
	//
	// if params.Limit != nil {
	// 	request.Limit = uint32(*params.Limit)
	// }
	//
	// result, err := c.linkServiceClient.List(r.Context(), request)
	// if err != nil {
	// 	var errorLink *v1.NotFoundError
	//
	// 	if errors.Is(err, errorLink) {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 		return
	// 	}
	//
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 	return
	// }
	//
	// response := &api.GetLinks200Response{
	// 	Links:      make([]api.Link, 0, len(result.GetLinks().GetLink())),
	// 	NextCursor: result.GetCursor(),
	// }
	//
	// for _, link := range result.GetLinks().GetLink() {
	// 	response.Links = append(response.Links, api.Link{
	// 		Url:       link.GetUrl(),
	// 		Hash:      link.GetHash(),
	// 		Describe:  link.GetDescribe(),
	// 		CreatedAt: link.GetCreatedAt().AsTime(),
	// 		UpdatedAt: link.GetUpdatedAt().AsTime(),
	// 	})
	// }
	//
	// w.WriteHeader(http.StatusOK)
	// err = json.NewEncoder(w).Encode(response)
	// if err != nil {
	// 	c.log.Error(err.Error())
	// }
}

// DeleteLink - delete link
func (c *Controller) DeleteLink(w http.ResponseWriter, r *http.Request, hash string) {
	// _, err := c.linkServiceClient.Delete(r.Context(), &link_rpc.DeleteRequest{Hash: hash})
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
}
