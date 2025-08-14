/*
Main DI-package
*/
package di

import (
	"github.com/google/wire"

	"github.com/shortlink-org/shortlink/pkg/cache"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	shortctx "github.com/shortlink-org/shortlink/pkg/di/pkg/context"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/flags"
	logger_di "github.com/shortlink-org/shortlink/pkg/di/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/permission"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	traicing_di "github.com/shortlink-org/shortlink/pkg/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
)

// DefaultSet ==========================================================================================================
var DefaultSet = wire.NewSet(
	shortctx.New,
	autoMaxPro.New,
	flags.New,
	config.New,
	logger_di.New,
	traicing_di.New,
	metrics.New,
	cache.New,
	profiling.New,
	permission.New,
)
