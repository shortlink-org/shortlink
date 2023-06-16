package account

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/encoding/protojson"

	account_application "github.com/shortlink-org/shortlink/internal/services/billing/application/account"
	billing "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/account/v1"
)

type AccoutAPI struct {
	jsonpb protojson.MarshalOptions

	accountService account_application.AccountService
}

func New(accountService *account_application.AccountService) (*AccoutAPI, error) {
	return &AccoutAPI{
		accountService: *accountService,
	}, nil
}

// Routes creates a REST router
func (api *AccoutAPI) Routes(r chi.Router) {
	r.Get("/accounts", api.list)
	r.Get("/account/{hash}", api.get)
	r.Post("/account", api.add)
	r.Delete("/account/{hash}", api.delete)
}

// add ...
func (api *AccoutAPI) add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request billing.Account
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	newAccount, err := api.accountService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := api.jsonpb.Marshal(newAccount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint:errcheck
}

// get ...
func (api *AccoutAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint:errcheck
}

// list ...
func (api *AccoutAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint:errcheck
}

// delete ...
func (api *AccoutAPI) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint:errcheck
}
