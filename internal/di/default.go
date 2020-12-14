//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"fmt"
	"net/http"

	"github.com/heptiolabs/healthcheck"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/batazor/shortlink/internal/traicing"

	"github.com/batazor/shortlink/internal/logger"
)

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

// Cobra/Flags =========================================================================================================
func InitFlags() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	return rootCmd, nil
}

// Health ==============================================================================================================
func NewHealthCheck() (healthcheck.Handler, error) {
	// create a new health instance
	endpoint := healthcheck.NewHandler()

	// Expose the /live and /ready endpoints over HTTP
	go http.ListenAndServe("0.0.0.0:9090", endpoint)

	return endpoint, nil
}

// AutoMaxProcs ========================================================================================================
type diAutoMaxPro *string

// InitAutoMaxProcs - Automatically set GOMAXPROCS to match Linux container CPU quota
func InitAutoMaxProcs(log logger.Logger) (diAutoMaxPro, func(), error) {
	undo, err := maxprocs.Set(maxprocs.Logger(func(s string, args ...interface{}) {
		log.Info(fmt.Sprintf(s, args))
	}))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		undo()
	}

	return nil, cleanup, nil
}

// Tracing =============================================================================================================
func InitTracer(ctx context.Context, log logger.Logger) (opentracing.Tracer, func(), error) {
	viper.SetDefault("TRACER_SERVICE_NAME", "ShortLink") // Service Name
	viper.SetDefault("TRACER_URI", "localhost:6831")     // Tracing addr:host

	config := traicing.Config{
		ServiceName: viper.GetString("TRACER_SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	tracer, tracerClose, err := traicing.Init(config)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := tracerClose.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return tracer, cleanup, nil
}
