/*
Heath checker
*/

package healthcheck

import (
	"github.com/heptiolabs/healthcheck"
)

// Init ...
// Package healthcheck helps you implement Kubernetes liveness and readiness checks for your application.
// It supports synchronous and asynchronous (background) checks.
func Init() healthcheck.Handler { // nolint unused
	// Create a Handler that we can use to register liveness and readiness checks.
	health := healthcheck.NewHandler()

	// Add a liveness check to detect Goroutine leaks. If this fails we want
	// to be restarted/rescheduled.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	return health
}
