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

	"github.com/batazor/shortlink/internal/api/infrastructure/store"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/notify"
)

//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

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

	// Init Store
	var st db.Store
	_, err = st.Use(ctx, log)
	assert.Nil(t, err)

	store := store.LinkStore{}
	_, err = store.Use(ctx, log, nil)
	assert.Nil(t, err)

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
		_, body := testRequest(t, ts, "POST", "/", bytes.NewReader(payload)) // nolint bodyclose

		// Parse response
		var resp map[string]interface{}
		err = json.Unmarshal([]byte(body), &resp)
		assert.Nil(t, err)

		assert.Equal(t, resp["hash"], "92c9c679c")
	})

	t.Run("with db", func(t *testing.T) {
		payload, err := json.Marshal(addRequest{
			Describe: "",
		})
		assert.Nil(t, err)

		_, body := testRequest(t, ts, "POST", "/", bytes.NewReader(payload)) // nolint bodyclose
		assert.NotNil(t, body)

		// clean db subscribe
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

	t.Run("with db", func(t *testing.T) {
		// Init Store
		var st db.Store
		_, err = st.Use(ctx, log)
		assert.Nil(t, err)

		store := store.LinkStore{}
		_, err = store.Use(ctx, log, nil)
		assert.Nil(t, err)

		response := `{"error": "Not found link: hash"}`
		_, body := testRequest(t, ts, "GET", "/hash", nil) // nolint bodyclose
		assert.Equal(t, body, response)

		// clean db subscribe
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

	t.Run("with db", func(t *testing.T) {
		// Init Store
		var st db.Store
		_, err = st.Use(ctx, log)
		assert.Nil(t, err)

		store := store.LinkStore{}
		_, err = store.Use(ctx, log, nil)
		assert.Nil(t, err)

		response := `null`
		_, body := testRequest(t, ts, "GET", "/links", nil) // nolint bodyclose
		assert.Equal(t, body, response)

		// clean db subscribe
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
		payload, errJsonMarshal := json.Marshal(deleteRequest{
			Hash: "hash",
		})
		assert.Nil(t, errJsonMarshal)
		response := `{"error": "Not found subscribe to event METHOD_DELETE"}`
		_, body := testRequest(t, ts, "DELETE", "/", bytes.NewReader(payload)) // nolint bodyclose
		assert.Equal(t, body, response)
	})

	t.Run("with db", func(t *testing.T) {
		// Init Store
		var st db.Store
		_, err = st.Use(ctx, log)
		assert.Nil(t, err)

		store := store.LinkStore{}
		_, err = store.Use(ctx, log, nil)
		assert.Nil(t, err)

		payload, err := json.Marshal(deleteRequest{
			Hash: "hash",
		})
		assert.Nil(t, err)
		response := `{}`
		_, body := testRequest(t, ts, "DELETE", "/", bytes.NewReader(payload)) // nolint bodyclose
		assert.Equal(t, body, response)

		// clean db subscribe
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
