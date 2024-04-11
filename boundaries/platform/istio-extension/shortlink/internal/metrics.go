package internal

import (
	"fmt"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

type Metrics struct {
	// Our memoized map of metric names to ID
	counters map[string]proxywasm.MetricCounter
}

// We want to group all our metrics under a common prefix
const MetricPrefix = "envoy_wasm_shortlink_plugin"

func NewMetrics() *Metrics {
	return &Metrics{
		counters: make(map[string]proxywasm.MetricCounter),
	}
}

// Increment function increments the specific metric name by 1
func (m *Metrics) Increment(name string) {
	fullName := metricName(name)

	if _, exists := m.counters[fullName]; !exists {
		// We haven't seen this metric name before, so define it in the Envoy Host
		m.counters[fullName] = proxywasm.DefineCounterMetric(fullName)
	}

	m.counters[fullName].Increment(1)
}

func metricName(name string) string {
	return fmt.Sprintf("%s_%s", MetricPrefix, name)
}
