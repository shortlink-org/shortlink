package tariff

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/encoding/protojson"

	tariff_application "github.com/shortlink-org/shortlink/internal/services/billing/application/tariff"
	billing "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/tariff/v1"
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
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request billing.Tariff
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	newTariff, err := api.tariffService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := api.jsonpb.Marshal(newTariff)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint:errcheck
}

// Get ...
func (api *TariffAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint:errcheck
}

// List ...
func (api *TariffAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	tariffs, err := api.tariffService.List(r.Context(), nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := api.jsonpb.Marshal(tariffs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint:errcheck
}

// Delete ...
func (api *TariffAPI) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint:errcheck
}
