package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"

	"github.com/shortlink-org/shortlink/boundaries/platform/istio-extension/shortlink/internal"
)

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type filterContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

// NewPluginContext Override types.DefaultVMContext otherwise this plugin would do nothing :)
func (v *vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	proxywasm.LogInfof("NewPluginContext context:%v", contextID)

	return &filterContext{}
}

// OnPluginStart Override types.DefaultPluginContext.
func (h *filterContext) OnPluginStart(_ int) types.OnPluginStartStatus {
	return types.OnPluginStartStatusOK
}

// NewHttpContext Override types.DefaultPluginContext to allow us to declare a request handler for each
// intercepted request the Envoy Sidecar sends us
func (h *filterContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &internal.RequestHandler{}
}
