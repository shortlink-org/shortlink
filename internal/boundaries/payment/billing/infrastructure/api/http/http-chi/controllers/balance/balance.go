package balance

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/segmentio/encoding/json"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/encoding/protojson"

	payment_application "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/application/payment"
	billing "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/domain/billing/payment/v1"
)

type API struct {
	jsonpb protojson.MarshalOptions //nolint:structcheck // ignore

	paymentService *payment_application.PaymentService
}

func New(paymentService *payment_application.PaymentService) (*API, error) {
	return &API{
		paymentService: paymentService,
	}, nil
}

// Routes create a REST router
func (api *API) Routes(r chi.Router) {
	r.Put("/balance/{payment_id}", api.update)
}

// Update - update
func (api *API) update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace_id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request billing.Payment
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck,goconst // ignore

		return
	}

	updatePayment, err := api.paymentService.UpdateBalance(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	res, err := api.jsonpb.Marshal(updatePayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) //nolint:errcheck // ignore
}
