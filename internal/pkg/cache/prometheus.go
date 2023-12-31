package cache

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cacheOperationDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cache_operation_duration_seconds",
		Help:    "Duration of cache operations in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"operation", "key"})

	cacheOperationErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "cache_operation_errors_total",
		Help: "Total number of errors in cache operations",
	}, []string{"operation", "key"})

	cacheOperations = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "cache_operations_total",
		Help: "Total number of successful cache operations",
	}, []string{"operation", "status", "key"})
)
