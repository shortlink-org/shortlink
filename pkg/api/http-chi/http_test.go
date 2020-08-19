package httpchi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/internal/store"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestAdd(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	t.Run("empty payload", func(t *testing.T) {
		response := `{"error": "EOF"}`
		_, body := testRequest(t, ts, "POST", "/", nil) // nolint bodyclose
		assert.Equal(t, body, response)
	})

	t.Run("correct payload", func(t *testing.T) {
		payload, err := json.Marshal(addRequest{
			URL:      "http://test.com",
			Describe: "",
		})
		assert.Nil(t, err)
		response := `{"error": "Not found subscribe to event METHOD_ADD"}`
		_, body := testRequest(t, ts, "POST", "/", bytes.NewReader(payload)) // nolint bodyclose
		assert.Equal(t, body, response)
	})

	t.Run("with store", func(t *testing.T) {
		// add store
		var st store.Store
		st.Use(ctx, log)

		payload, err := json.Marshal(addRequest{
			Describe: "",
		})
		assert.Nil(t, err)

		_, body := testRequest(t, ts, "POST", "/", bytes.NewReader(payload)) // nolint bodyclose
		assert.NotNil(t, body)

		// clean store subscribe
		notify.Clean()
	})
}

func TestGet(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	t.Run("correct payload", func(t *testing.T) {
		response := `{"error": "Not found subscribe to event METHOD_GET"}`
		_, body := testRequest(t, ts, "GET", "/hash", nil) // nolint bodyclose
		assert.Equal(t, body, response)
	})

	t.Run("with store", func(t *testing.T) {
		// add store
		var st store.Store
		st.Use(ctx, log)

		response := `{"error": "Not found link: hash"}`
		_, body := testRequest(t, ts, "GET", "/hash", nil) // nolint bodyclose
		assert.Equal(t, body, response)

		// clean store subscribe
		notify.Clean()
	})
}

func TestList(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	t.Run("correct payload", func(t *testing.T) {
		response := `{"error": "Not found subscribe to event METHOD_LIST"}`
		_, body := testRequest(t, ts, "GET", "/links", nil) // nolint bodyclose
		assert.Equal(t, body, response)
	})

	t.Run("with store", func(t *testing.T) {
		// add store
		var st store.Store
		st.Use(ctx, log)

		response := `null`
		_, body := testRequest(t, ts, "GET", "/links", nil) // nolint bodyclose
		assert.Equal(t, body, response)

		// clean store subscribe
		notify.Clean()
	})
}

func TestDelete(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	t.Run("correct payload", func(t *testing.T) {
		payload, err := json.Marshal(deleteRequest{
			Hash: "hash",
		})
		assert.Nil(t, err)
		response := `{"error": "Not found subscribe to event METHOD_DELETE"}`
		_, body := testRequest(t, ts, "DELETE", "/", bytes.NewReader(payload)) // nolint bodyclose
		assert.Equal(t, body, response)
	})

	t.Run("with store", func(t *testing.T) {
		// add store
		var st store.Store
		st.Use(ctx, log)

		payload, err := json.Marshal(deleteRequest{
			Hash: "hash",
		})
		assert.Nil(t, err)
		response := `{}`
		_, body := testRequest(t, ts, "DELETE", "/", bytes.NewReader(payload)) // nolint bodyclose
		assert.Equal(t, body, response)

		// clean store subscribe
		notify.Clean()
	})
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) { // nolint unparam
	req, err := http.NewRequestWithContext(context.TODO(), method, ts.URL+path, body)
	if err != nil {
		assert.Nil(t, err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		assert.Nil(t, err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Nil(t, err)
		return nil, ""
	}

	defer resp.Body.Close()
	return resp, string(respBody)
}
