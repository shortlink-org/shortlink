package link

import (
	"net/http"
)

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
