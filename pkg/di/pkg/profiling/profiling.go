package profiling

import (
	"context"
	"net/http"
	"net/http/pprof"
	"runtime"

	"github.com/grafana/pyroscope-go"
	pypprof "github.com/grafana/pyroscope-go/godeltaprof/http/pprof"
	"github.com/spf13/viper"

	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

type PprofEndpoint *http.ServeMux

func New(ctx context.Context, log logger.Logger) (PprofEndpoint, error) {
	viper.SetDefault("PROFILING_PORT", 7071) //nolint:revive,mnd // ignore
	viper.SetDefault("PROFILING_TIMEOUT", "30s")
	viper.SetDefault("PYROSCOPE_ADDRESS", "http://127.0.0.1:4040")

	// Create "common" listener
	pprofMux := http.NewServeMux()

	// Registration pprof-handlers
	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	pprofMux.HandleFunc("/debug/pprof/delta_heap", pypprof.Heap)
	pprofMux.HandleFunc("/debug/pprof/delta_block", pypprof.Block)
	pprofMux.HandleFunc("/debug/pprof/delta_mutex", pypprof.Mutex)

	go func() {
		config := http_server.Config{
			Port:    viper.GetInt("PROFILING_PORT"),
			Timeout: viper.GetDuration("PROFILING_TIMEOUT"),
		}
		server := http_server.New(ctx, pprofMux, config, nil)

		err := server.ListenAndServe()
		if err != nil {
			log.Error(err.Error())
		}
	}()
	log.Info("Run profiling", field.Fields{
		"addr": "0.0.0.0:7071",
	})

	// These 2 lines are only required if you're using mutex or block profiling
	// to read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: viper.GetString("SERVICE_NAME"),
		ServerAddress:   viper.GetString("PYROSCOPE_ADDRESS"),
		Logger:          nil,
		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	if err != nil {
		return nil, err
	}

	log.Info("Run pyroscope", field.Fields{
		"addr": viper.GetString("PYROSCOPE_ADDRESS"),
	})

	return pprofMux, nil
}
