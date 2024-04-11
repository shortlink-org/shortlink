package internal

import (
	"os"

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

func (r *RequestHandler) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	version, ok := os.LookupEnv("PLUGIN_VERSION")
	if !ok {
		version = "unknown"
	}

	_ = proxywasm.AddHttpResponseHeader("injected-by-istio-plugin-shortlink", version)

	// Ger response headers
	responseHeaders, _ := proxywasm.GetHttpResponseHeaders()

	// Print response headers
	proxywasm.LogWarnf("Response headers: %v \n", responseHeaders)

	proxywasm.LogWarn("WASM plugin Handling response")

	return types.ActionContinue
}

// headerArrayToMap is a simple function to convert from an array of headers to a Map
func headerArrayToMap(requestHeaders [][2]string) map[string]string {
	headerMap := make(map[string]string)

	for _, header := range requestHeaders {
		headerMap[header[0]] = header[1]
	}

	return headerMap
}
