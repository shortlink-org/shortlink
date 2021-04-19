package profiling

import (
	"net/http"
	"net/http/pprof"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
)

type PprofEndpoint *http.ServeMux

func New(log logger.Logger) PprofEndpoint {
	// Create an "common" listener
	pprofMux := http.NewServeMux()

	// Registration pprof-handlers
	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	go http.ListenAndServe("0.0.0.0:7071", pprofMux) // nolint errcheck
	log.Info("Run profiling", field.Fields{
		"addr": "0.0.0.0:7071",
	})

	return pprofMux
}
