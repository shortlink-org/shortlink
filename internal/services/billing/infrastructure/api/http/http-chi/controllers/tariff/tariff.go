package tariff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/batazor/shortlink/internal/services/api/application/http-chi/helpers"
	tariff_application "github.com/batazor/shortlink/internal/services/billing/application/tariff"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/tariff/v1"
)

type TariffAPI struct {
	jsonpb protojson.MarshalOptions

	tariffService *tariff_application.TariffService
}

func New(tariffService *tariff_application.TariffService) (*TariffAPI, error) {
	return &TariffAPI{
		tariffService: tariffService,
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

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request billing.Tariff
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	newTariff, err := api.tariffService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := api.jsonpb.Marshal(newTariff)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint errcheck
}

// Get ...
func (api *TariffAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}

// List ...
func (api *TariffAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	tariffs, err := api.tariffService.List(r.Context(), nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := api.jsonpb.Marshal(tariffs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// Delete ...
func (api *TariffAPI) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}
