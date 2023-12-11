package payment

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
	jsonpb protojson.MarshalOptions

	paymentService *payment_application.PaymentService
}

func New(paymentService *payment_application.PaymentService) (*API, error) {
	return &API{
		paymentService: paymentService,
	}, nil
}

// Routes create a REST router
func (api *API) Routes(r chi.Router) {
	r.Get("/payment/{id}", api.get)
	r.Get("/payments", api.list)
	r.Post("/payment", api.open)
	r.Delete("/payment/{id}", api.close)
}

// open a new payment
func (api *API) open(w http.ResponseWriter, r *http.Request) {
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

	newPayment, err := api.paymentService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	res, err := api.jsonpb.Marshal(newPayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) //nolint:errcheck // ignore
}

// get payment by identity
func (api *API) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace_id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	aggregateId := chi.URLParam(r, "id")
	if aggregateId == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set payment of identity"}`)) //nolint:errcheck // ignore

		return
	}

	getPayment, err := api.paymentService.Get(r.Context(), aggregateId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	res, err := api.jsonpb.Marshal(getPayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) //nolint:errcheck // ignore
}

// list - get all payments of users
func (api *API) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace_id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) //nolint:errcheck // ignore
}

// close a payment
func (api *API) close(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace_id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	aggregateId := chi.URLParam(r, "id")
	if aggregateId == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set payment of identity"}`)) //nolint:errcheck // ignore

		return
	}

	err := api.paymentService.Close(r.Context(), aggregateId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) //nolint:errcheck // ignore
}
