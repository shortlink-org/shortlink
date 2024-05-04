package balance

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"

	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/billing/payment/v1"
	payment_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/payment"
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
	w.Header().Add("Content-Type", "application/json")

	// Parse request
	var request billing.Payment
	err := json.NewDecoder(r.Body).Decode(&request)
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
