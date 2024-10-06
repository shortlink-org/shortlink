//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Pricer DI-package
*/
package di

import (
	"log"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/application"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/infrastructure"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/interfaces/cli"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type PricerService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint

	// Application
	CartService *application.CartService

	// CLI
	CLIHandler *cli.CLIHandler
}

// PricerService =======================================================================================================
var PricerSet = wire.NewSet(
	// Common
	di.DefaultSet,
	rpc.InitServer,

	// Repository
	newDiscountPolicy,
	newTaxPolicy,
	newPolicyNames,

	// Application
	application.NewCartService,

	NewPricerService,
)

// newDiscountPolicy creates a new DiscountPolicy
func newDiscountPolicy() application.DiscountPolicy {
	discountPolicyPath := viper.GetString("policies.discounts")
	discountQuery := viper.GetString("queries.discounts")

	discountEvaluator, err := infrastructure.NewOPAEvaluator(discountPolicyPath, discountQuery)
	if err != nil {
		log.Fatalf("Failed to initialize Discount Policy Evaluator: %v", err)
	}

	return discountEvaluator
}

// newTaxPolicy creates a new TaxPolicy
func newTaxPolicy() application.TaxPolicy {
	taxPolicyPath := viper.GetString("policies.taxes")
	taxQuery := viper.GetString("queries.taxes")

	taxEvaluator, err := infrastructure.NewOPAEvaluator(taxPolicyPath, taxQuery)
	if err != nil {
		log.Fatalf("Failed to initialize Tax Policy Evaluator: %v", err)
	}

	return taxEvaluator
}

// newPolicyNames retrieves policy names
func newPolicyNames() ([]string, error) {
	discountPolicyPath := viper.GetString("policies.discounts")
	taxPolicyPath := viper.GetString("policies.taxes")

	return infrastructure.GetPolicyNames(discountPolicyPath, taxPolicyPath)
}

func NewPricerService(
	// Common
	log logger.Logger,
	config *config.Config,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	// Application
	cartService *application.CartService,
) (*PricerService, error) {
	return &PricerService{
		// Common
		Log:        log,
		Config:     config,
		AutoMaxPro: autoMaxProcsOption,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,

		// Application
		CartService: cartService,
	}, nil
}

func InitializePricerService() (*PricerService, func(), error) {
	panic(wire.Build(PricerSet))
}
