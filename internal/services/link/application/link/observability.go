package link

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	prometheus.MustRegister(newLinkHistogram)
}

var newLinkHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace: "link",
	Subsystem: "application",
	Name:      "new",
})

func NewLinkHistogramObserve(ctx context.Context) {
	traceID := trace.SpanContextFromContext(ctx).TraceID()

	if exemplarObserver, ok := newLinkHistogram.(prometheus.ExemplarObserver); ok && traceID.IsValid() {
		exemplarObserver.ObserveWithExemplar(1, prometheus.Labels{
			"trace-id": traceID.String(),
		})

		return
	}

	newLinkHistogram.Observe(1)
}
