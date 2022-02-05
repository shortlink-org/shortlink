package payment

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/batazor/shortlink/internal/services/api/application/http-chi/helpers"
	payment_application "github.com/batazor/shortlink/internal/services/billing/application/payment"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

type PaymentAPI struct {
	jsonpb protojson.MarshalOptions // nolint structcheck

	paymentService *payment_application.PaymentService
}

func New(paymentService *payment_application.PaymentService) (*PaymentAPI, error) {
	return &PaymentAPI{
		paymentService: paymentService,
	}, nil
}

// Routes creates a REST router
func (api *PaymentAPI) Routes(r chi.Router) {
	r.Get("/payment/{id}", api.get)
	r.Get("/payments", api.list)
	r.Post("/payment", api.open)
	r.Delete("/payment/{id}", api.close)
}

// open a new payment
func (api *PaymentAPI) open(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request billing.Payment
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	newPayment, err := api.paymentService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := api.jsonpb.Marshal(newPayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint errcheck
}

// get payment by identity
func (api *PaymentAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	var aggregateId = chi.URLParam(r, "id")
	if aggregateId == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set payment of identity"}`)) // nolint errcheck
		return
	}

	getPayment, err := api.paymentService.Get(r.Context(), aggregateId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := api.jsonpb.Marshal(getPayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// list - get all payments of users
func (api *PaymentAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}

// close a payment
func (api *PaymentAPI) close(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	var aggregateId = chi.URLParam(r, "id")
	if aggregateId == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set payment of identity"}`)) // nolint errcheck
		return
	}

	err := api.paymentService.Close(r.Context(), aggregateId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}
