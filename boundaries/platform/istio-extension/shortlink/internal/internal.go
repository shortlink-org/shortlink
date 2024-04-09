package internal

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type RequestHandler struct {
	// Bring in the callback functions
	types.DefaultHttpContext
}

const (
	XRequestIdHeader = "x-request-id"
)

// OnHttpRequestHeaders is called on every request we intercept with this WASM filter
// Check out the types.HttpContext interface to see what other callbacks you can override
//
// Note: Parameters are not needed here, but a brief description:
//   - numHeaders = fairly self-explanatory, the number of request headers
//   - endOfStream = only set to false when there is a request body (e.g. in a POST/PATCH/PUT request)
func (r *RequestHandler) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogInfof("WASM plugin Handling request")

	// Get the actual request headers from the Envoy Sidecar
	requestHeaders, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get request headers: %v", err)
		// Allow Envoy Sidecar to forward this request to the upstream service
		return types.ActionContinue
	}

	// Convert the request headers to a map for easier access (more useful in subsequent sections)
	reqHeaderMap := headerArrayToMap(requestHeaders)

	// Get the x-request-id for grouping logs belonging to the same request
	xRequestID := reqHeaderMap[XRequestIdHeader]

	// Now we can take action on this request
	return r.doSomethingWithRequest(reqHeaderMap, xRequestID)
}

// headerArrayToMap is a simple function to convert from array of headers to a Map
func headerArrayToMap(requestHeaders [][2]string) map[string]string {
	headerMap := make(map[string]string)
	for _, header := range requestHeaders {
		headerMap[header[0]] = header[1]
	}
	return headerMap
}

func (r *RequestHandler) doSomethingWithRequest(reqHeaderMap map[string]string, xRequestID string) types.Action {
	// for now, let's just log all the request headers to we get an idea of what we have to work with
	for key, value := range reqHeaderMap {
		proxywasm.LogInfof("  %s: request header --> %s: %s", xRequestID, key, value)
	}

	// Forward request to upstream service, i.e. unblock request
	return types.ActionContinue
}
