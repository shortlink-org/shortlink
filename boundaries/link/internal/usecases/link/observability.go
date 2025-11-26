package link

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"
)

var newLinkHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace: "link",
	Subsystem: "application",
	Name:      "new",
	Help:      "New link event",
})

var _ = registerLinkHistogram()

func registerLinkHistogram() struct{} {
	prometheus.MustRegister(newLinkHistogram)

	return struct{}{}
}

func NewLinkHistogramObserve(ctx context.Context) {
	now := time.Now()
	spanContext := trace.SpanContextFromContext(ctx)

	if spanContext.IsValid() {
		if exemplarObserver, ok := newLinkHistogram.(prometheus.ExemplarObserver); ok && spanContext.TraceID().IsValid() {
			exemplarObserver.ObserveWithExemplar(time.Since(now).Seconds(), prometheus.Labels{
				"traceID": spanContext.TraceID().String(),
			})

			return
		}
	}

	newLinkHistogram.Observe(time.Since(now).Seconds())
}
