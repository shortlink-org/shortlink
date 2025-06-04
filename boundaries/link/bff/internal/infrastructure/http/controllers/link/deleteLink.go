package link

import (
	"net/http"
)

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
