package temporal

import (
	"fmt"
	"time"

	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	sdktally "go.temporal.io/sdk/contrib/tally"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
)

func newPrometheusScope(c *prometheus.Configuration, monitor *monitoring.Monitoring, log logger.Logger) (tally.Scope, error) {
	reporter, err := c.NewReporter(
		prometheus.ConfigurationOptions{
			Registry: monitor.Prometheus,
			OnError: func(err error) {
				log.Error(fmt.Sprintf("error in prometheus reporter: %v", err))
			},
		},
	)
	if err != nil {
		return nil, err
	}

	scopeOpts := tally.ScopeOptions{
		CachedReporter:  reporter,
		Separator:       prometheus.DefaultSeparator,
		SanitizeOptions: &sdktally.PrometheusSanitizeOptions,
	}
	scope, _ := tally.NewRootScope(scopeOpts, time.Second)
	scope = sdktally.NewPrometheusNamingScope(scope)

	return scope, nil
}
