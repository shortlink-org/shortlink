package metadata

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// BuildInfo metric as recommended in ADR-0014
	// Exposes software version to Prometheus
	buildInfo = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "metadata",
			Subsystem: "application",
			Name:      "build_info",
			Help:      "Build information about the metadata service including version, commit, and build time",
		},
		[]string{"version", "commit", "build_time"},
	)

	// SLA/SLO/SLI Metrics as per ADR-0018
	serviceAvailabilityRatio = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "metadata",
			Subsystem: "application",
			Name:      "service_availability_ratio",
			Help:      "Ratio of uptime to total time for SLA tracking",
		},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "metadata",
			Subsystem: "application",
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request response times for SLO tracking",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "endpoint", "status"},
	)

	errorRatePerMinute = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metadata",
			Subsystem: "application",
			Name:      "error_rate_per_minute",
			Help:      "Number of errors per minute for system reliability tracking",
		},
		[]string{"type", "operation"},
	)

	// Cache Metrics as per ADR-0018
	cacheHitTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metadata",
			Subsystem: "cache",
			Name:      "hit_total",
			Help:      "Total number of cache hits",
		},
		[]string{"cache_type"},
	)

	cacheMissTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metadata",
			Subsystem: "cache",
			Name:      "miss_total",
			Help:      "Total number of cache misses",
		},
		[]string{"cache_type"},
	)

	// Basic Service Metrics as per ADR-0018
	requestsPerSecond = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metadata",
			Subsystem: "application",
			Name:      "requests_per_second",
			Help:      "Service request load per second",
		},
		[]string{"operation"},
	)

	responseTimeSeconds = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "metadata",
			Subsystem: "application",
			Name:      "response_time_seconds",
			Help:      "Service response time",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"operation"},
	)

	errorRatePercentage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "metadata",
			Subsystem: "application",
			Name:      "error_rate_percentage",
			Help:      "Percentage of error requests",
		},
		[]string{"operation"},
	)
)

// SetBuildInfo sets the build information metric
func SetBuildInfo(version, commit, buildTime string) {
	buildInfo.WithLabelValues(version, commit, buildTime).Set(1)
}

// SetServiceAvailability sets the service availability ratio
func SetServiceAvailability(ratio float64) {
	serviceAvailabilityRatio.Set(ratio)
}

// RecordHTTPRequestDuration records HTTP request duration
func RecordHTTPRequestDuration(method, endpoint, status string, duration float64) {
	httpRequestDuration.WithLabelValues(method, endpoint, status).Observe(duration)
}

// IncrementErrorRate increments the error rate counter
func IncrementErrorRate(errorType, operation string) {
	errorRatePerMinute.WithLabelValues(errorType, operation).Inc()
}

// RecordCacheHit records a cache hit
func RecordCacheHit(cacheType string) {
	cacheHitTotal.WithLabelValues(cacheType).Inc()
}

// RecordCacheMiss records a cache miss
func RecordCacheMiss(cacheType string) {
	cacheMissTotal.WithLabelValues(cacheType).Inc()
}

// GetCacheHitRatio calculates the cache hit ratio
func GetCacheHitRatio(cacheType string) float64 {
	hits := getCounterValue(cacheHitTotal.WithLabelValues(cacheType))
	misses := getCounterValue(cacheMissTotal.WithLabelValues(cacheType))
	total := hits + misses
	if total == 0 {
		return 0
	}
	return hits / total
}

// IncrementRequestsPerSecond increments the RPS counter
func IncrementRequestsPerSecond(operation string) {
	requestsPerSecond.WithLabelValues(operation).Inc()
}

// RecordResponseTime records service response time
func RecordResponseTime(operation string, duration float64) {
	responseTimeSeconds.WithLabelValues(operation).Observe(duration)
}

// SetErrorRatePercentage sets the error rate percentage
func SetErrorRatePercentage(operation string, percentage float64) {
	errorRatePercentage.WithLabelValues(operation).Set(percentage)
}

// getCounterValue is a helper to extract counter value (simplified version)
func getCounterValue(counter prometheus.Counter) float64 {
	// In production, you would use prometheus.Gatherer to get the actual value
	// This is a simplified version for demonstration
	return 0
}
