package httpchi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestAdd(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("empty payload", func(t *testing.T) {
		response := `{"error": "EOF"}`
		if _, body := testRequest(t, ts, "POST", "/", nil); body != response { // nolint bodyclose
			t.Errorf(`Assert: %s. Got %s`, response, body)
		}
	})

	t.Run("correct payload", func(t *testing.T) {
		payload, err := json.Marshal(addRequest{
			URL:      "http://test.com",
			Describe: "",
		})
		if err != nil {
			t.Error(err)
		}
		response := `{"error": "Not found subscribe to event METHOD_ADD"}`
		if _, body := testRequest(t, ts, "POST", "/", bytes.NewReader(payload)); body != response { // nolint bodyclose
			t.Errorf(`Assert: %s. Got %s`, response, body)
		}
	})

}

func TestGet(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("correct payload", func(t *testing.T) {
		response := `{"error": "Not found subscribe to event METHOD_GET"}`
		if _, body := testRequest(t, ts, "GET", "/hash", nil); body != response { // nolint bodyclose
			t.Errorf(`Assert: %s. Got %s`, response, body)
		}
	})
}

func TestList(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("correct payload", func(t *testing.T) {
		response := `{"error": "Not found subscribe to event METHOD_LIST"}`
		if _, body := testRequest(t, ts, "GET", "/links", nil); body != response { // nolint bodyclose
			t.Errorf(`Assert: %s. Got %s`, response, body)
		}
	})
}

func TestDelete(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("correct payload", func(t *testing.T) {
		payload, err := json.Marshal(deleteRequest{
			Hash: "hash",
		})
		if err != nil {
			t.Error(err)
		}
		response := `{"error": "Not found subscribe to event METHOD_DELETE"}`
		if _, body := testRequest(t, ts, "DELETE", "/", bytes.NewReader(payload)); body != response { // nolint bodyclose
			t.Errorf(`Assert: %s. Got %s`, response, body)
		}
	})
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) { // nolint unparam
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	defer resp.Body.Close()
	return resp, string(respBody)
}
