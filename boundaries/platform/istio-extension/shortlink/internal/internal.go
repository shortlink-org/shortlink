package internal

import (
	"os"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type RequestHandler struct {
	// Bring in the callback functions
	types.DefaultHttpContext

	Metrics *Metrics
}

func (r *RequestHandler) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogInfof("WASM plugin Handling request")

	// Forward request to upstream service, i.e. unblock request
	return types.ActionContinue
}

func (r *RequestHandler) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	version, ok := os.LookupEnv("PLUGIN_VERSION")
	if !ok {
		version = "unknown"
	}

	_ = proxywasm.AddHttpResponseHeader("injected-by-istio-plugin-shortlink", version)
	r.Metrics.Increment("requests_intercepted")

	proxywasm.LogInfo("WASM plugin Handling response")

	return types.ActionContinue
}
