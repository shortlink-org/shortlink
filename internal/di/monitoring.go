//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"net/http"
	"net/http/pprof"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Monitoring ==========================================================================================================
func InitMonitoring(sentryHandler *sentryhttp.Handler) *http.ServeMux {
	// Create a new Prometheus registry
	registry := prometheus.NewRegistry()

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(registry, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	// Create an "common" listener
	commonMux := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	commonMux.Handle("/metrics", sentryHandler.Handle(promhttp.Handler()))

	// Expose a liveness check on /live
	commonMux.HandleFunc("/live", sentryHandler.HandleFunc(health.LiveEndpoint))

	// Expose a readiness check on /ready
	commonMux.HandleFunc("/ready", sentryHandler.HandleFunc(health.ReadyEndpoint))

	return commonMux
}

// Profiling ===========================================================================================================
type PprofEndpoint *http.ServeMux

func InitProfiling() PprofEndpoint {
	// Create an "common" listener
	pprofMux := http.NewServeMux()

	// Registration pprof-handlers
	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return pprofMux
}
