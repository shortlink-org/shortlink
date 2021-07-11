package payment

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	payment_application "github.com/batazor/shortlink/internal/services/billing/application/payment"
)

type PaymentAPI struct {
	jsonpb protojson.MarshalOptions

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
	r.Post("/payment", api.add)
	r.Delete("/payment/{id}", api.delete)
}

// add ...
func (api *PaymentAPI) add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// get ...
func (api *PaymentAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// list ...
func (api *PaymentAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// delete ...
func (api *PaymentAPI) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}
