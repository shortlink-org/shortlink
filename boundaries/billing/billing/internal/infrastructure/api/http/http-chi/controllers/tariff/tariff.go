package tariff

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"

	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/tariff/v1"
	tariff_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/tariff"
)

type API struct {
	jsonpb protojson.MarshalOptions

	tariffService *tariff_application.TariffService
}

func New(tariffService *tariff_application.TariffService) (*API, error) {
	return &API{
		tariffService: tariffService,
	}, nil
}

// Routes create a REST router
func (api *API) Routes(r chi.Router) {
	r.Get("/tariffs", api.list)
	r.Get("/tariff/{hash}", api.get)
	r.Post("/tariff", api.add)
	r.Delete("/tariff/{hash}", api.delete)
}

// Add - add
func (api *API) add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// Parse request
	var request billing.Tariff
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck,goconst // ignore

		return
	}

	newTariff, err := api.tariffService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	res, err := api.jsonpb.Marshal(newTariff)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) //nolint:errcheck // ignore
}

// Get - get
func (api *API) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

// list - list
func (api *API) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	tariffs, err := api.tariffService.List(r.Context(), nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	res, err := api.jsonpb.Marshal(tariffs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) //nolint:errcheck // ignore
}

// Delete - delete
func (api *API) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusNoContent)
}
