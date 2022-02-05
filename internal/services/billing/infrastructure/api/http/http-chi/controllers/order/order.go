package order

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/batazor/shortlink/internal/services/api/application/http-chi/helpers"
	order_application "github.com/batazor/shortlink/internal/services/billing/application/order"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/order/v1"
)

type OrderAPI struct {
	jsonpb protojson.MarshalOptions // nolint structcheck

	orderService *order_application.OrderService
}

func New(orderService *order_application.OrderService) (*OrderAPI, error) {
	return &OrderAPI{
		orderService: orderService,
	}, nil
}

// Routes creates a REST router
func (api *OrderAPI) Routes(r chi.Router) {
	r.Get("/order/{hash}", api.get)
	r.Get("/orders", api.list)
	r.Post("/order", api.add)
	r.Delete("/order/{hash}", api.delete)
}

// Add ...
func (api *OrderAPI) add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request billing.Order
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	newOrder, err := api.orderService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := api.jsonpb.Marshal(newOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint errcheck
}

// Get ...
func (api *OrderAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}

// List ...
func (api *OrderAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}

// Delete ...
func (api *OrderAPI) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}
