package profiling

import (
	"net/http"
	"net/http/pprof"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

type PprofEndpoint *http.ServeMux

func New(log logger.Logger) PprofEndpoint {
	// Create "common" listener
	pprofMux := http.NewServeMux()

	// Registration pprof-handlers
	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	go func() {
		err := http.ListenAndServe("0.0.0.0:7071", pprofMux)
		if err != nil {
			log.Error(err.Error())
		}
	}()
	log.Info("Run profiling", field.Fields{
		"addr": "0.0.0.0:7071",
	})

	return pprofMux
}
