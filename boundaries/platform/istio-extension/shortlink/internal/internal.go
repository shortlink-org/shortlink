package internal

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type RequestHandler struct {
	// Bring in the callback functions
	types.DefaultHttpContext

	ContextID uint32
}

func (r *RequestHandler) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogInfof("WASM plugin Handling request")

	// Forward request to upstream service, i.e. unblock request
	return types.ActionContinue
}

// Additional headers supposed to be injected to response headers.
var additionalHeaders = map[string]string{
	"who-am-i":    "wasm-extension",
	"injected-by": "istio-api!",
}

func (r *RequestHandler) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	for key, value := range additionalHeaders {
		proxywasm.AddHttpResponseHeader(key, value)
	}

	proxywasm.LogWarn("WASM plugin Handling response")

	return types.ActionContinue
}
