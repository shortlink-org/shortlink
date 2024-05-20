package account

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"

	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1"
	account_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/account"
)

type AccoutAPI struct {
	jsonpb protojson.MarshalOptions

	accountService account_application.AccountService
}

func New(accountService *account_application.AccountService) (*AccoutAPI, error) {
	return &AccoutAPI{
		accountService: *accountService,
	}, nil
}

// Routes creates a REST router
func (api *AccoutAPI) Routes(r chi.Router) {
	r.Get("/accounts", api.list)
	r.Get("/account/{hash}", api.get)
	r.Post("/account", api.add)
	r.Delete("/account/{hash}", api.delete)
}

// Add - add
func (api *AccoutAPI) add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// Parse request
	var request billing.Account
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck,goconst // ignore

		return
	}

	newAccount, err := api.accountService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	res, err := api.jsonpb.Marshal(newAccount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) //nolint:errcheck // ignore
}

// Get - get
func (api *AccoutAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

// List - list
func (api *AccoutAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

// Delete - delete
func (api *AccoutAPI) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusNoContent)
}
