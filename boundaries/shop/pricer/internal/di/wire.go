//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Pricer DI-package
*/
package di

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/application"
	pkg_di "github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/di/pkg"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/infrastructure/cli"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/infrastructure/policy_evaluator"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/infrastructure/rpc/run"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
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

	// Delivery
	run *run.Response

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
	pkg_di.ReadConfig,

	// Repository
	newDiscountPolicy,
	newTaxPolicy,
	newPolicyNames,

	// Delivery
	NewRunRPCServer,

	// Application
	application.NewCartService,
	newCLIHandler,

	NewPricerService,
)

// TODO: refactoring. maybe drop this function
func NewRunRPCServer(runRPCServer *rpc.Server) (*run.Response, error) {
	return run.Run(runRPCServer)
}

// newDiscountPolicy creates a new DiscountPolicy
func newDiscountPolicy(ctx context.Context, log logger.Logger, cfg *pkg_di.Config) application.DiscountPolicy {
	discountPolicyPath := viper.GetString("policies.discounts")
	discountQuery := viper.GetString("queries.discounts")

	discountEvaluator, err := policy_evaluator.NewOPAEvaluator(log, discountPolicyPath, discountQuery)
	if err != nil {
		log.ErrorWithContext(ctx, "Failed to initialize Discount Policy Evaluator: %v", field.Fields{"error": err})
	}

	return discountEvaluator
}

// newTaxPolicy creates a new TaxPolicy
func newTaxPolicy(ctx context.Context, log logger.Logger, cfg *pkg_di.Config) application.TaxPolicy {
	taxPolicyPath := viper.GetString("policies.taxes")
	taxQuery := viper.GetString("queries.taxes")

	taxEvaluator, err := policy_evaluator.NewOPAEvaluator(log, taxPolicyPath, taxQuery)
	if err != nil {
		log.ErrorWithContext(ctx, "Failed to initialize Tax Policy Evaluator: %v", field.Fields{"error": err})
	}

	return taxEvaluator
}

// newPolicyNames retrieves policy names
func newPolicyNames(cfg *pkg_di.Config) ([]string, error) {
	discountPolicyPath := viper.GetString("policies.discounts")
	taxPolicyPath := viper.GetString("policies.taxes")

	return policy_evaluator.GetPolicyNames(discountPolicyPath, taxPolicyPath)
}

// newCLIHandler creates a new CLIHandler
func newCLIHandler(ctx context.Context, log logger.Logger, cartService *application.CartService, cfg *pkg_di.Config) *cli.CLIHandler {
	cartFiles := viper.GetStringSlice("cart_files")
	outputDir := viper.GetString("output_dir")

	discountParams := viper.GetStringMap("params.discount")
	taxParams := viper.GetStringMap("params.tax")

	cliHandler := &cli.CLIHandler{
		CartService: cartService,
		OutputDir:   outputDir,
	}

	// Process each cart file
	for _, cartFile := range cartFiles {
		fmt.Printf("Processing cart file: %s\n", cartFile)
		if err := cliHandler.Run(cartFile, discountParams, taxParams); err != nil {
			log.ErrorWithContext(ctx, "Error processing cart", field.Fields{"cart": cartFile, "error": err})
		}
	}

	return cliHandler
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

	// Delivery
	run *run.Response,

	// Application
	cartService *application.CartService,

	// CLI
	cliHandler *cli.CLIHandler,
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

		// Delivery
		run: run,

		// Application
		CartService: cartService,

		// CLI
		CLIHandler: cliHandler,
	}, nil
}

func InitializePricerService() (*PricerService, func(), error) {
	panic(wire.Build(PricerSet))
}
