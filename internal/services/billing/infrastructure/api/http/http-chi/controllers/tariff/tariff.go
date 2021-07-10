package tariff

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	billing_store "github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
)

type TariffAPI struct {
	jsonpb protojson.MarshalOptions
	store  *billing_store.TariffRepository
}

func New(store *billing_store.TariffRepository) (*TariffAPI, error) {
	return &TariffAPI{
		store: store,
	}, nil
}

// Routes creates a REST router
func (api *TariffAPI) Routes(r chi.Router) {
	r.Get("/tariffs", api.list)
	r.Get("/tariff/{hash}", api.get)
	r.Post("/tariff", api.add)
	r.Delete("/tariff/{hash}", api.delete)
}

// Add ...
func (api *TariffAPI) add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// Get ...
func (api *TariffAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// List ...
func (api *TariffAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// Delete ...
func (api *TariffAPI) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}
