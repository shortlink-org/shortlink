package balance

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	balance_application "github.com/batazor/shortlink/internal/services/billing/application/balance"
)

type BalanceAPI struct {
	jsonpb protojson.MarshalOptions

	balanceService *balance_application.BalanceService
}

func New(balanceService *balance_application.BalanceService) (*BalanceAPI, error) {
	return &BalanceAPI{
		balanceService: balanceService,
	}, nil
}

// Routes creates a REST router
func (api *BalanceAPI) Routes(r chi.Router) {
	r.Get("/balances", api.get)
	r.Put("/balance/{account_id}", api.update)
}

// Get ...
func (api *BalanceAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// Update ...
func (api *BalanceAPI) update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}
