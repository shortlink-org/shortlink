package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/di"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/interfaces/cli"
	"github.com/shortlink-org/shortlink/pkg/graceful_shutdown"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

func validateConfig() error {
	// Print the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	log.Printf("Current Working Directory: %s", wd)

	// Example validations
	if viper.GetString("policies.discounts") == "" {
		return errors.New("policies.discounts path is required")
	}
	if viper.GetString("policies.taxes") == "" {
		return errors.New("policies.taxes path is required")
	}
	if len(viper.GetStringSlice("cart_files")) == 0 {
		return errors.New("at least one cart file must be specified")
	}
	if viper.GetString("output_dir") == "" {
		return errors.New("output_dir is required")
	}
	return nil
}

func main() {
	viper.SetDefault("SERVICE_NAME", "shop-pricer")

	// Initialize Viper
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Validate configuration
	if err := validateConfig(); err != nil {
		log.Fatalf("Configuration validation error: %v", err)
	}

	// Init a new service
	service, cleanup, err := di.InitializePricerService()
	if err != nil {
		panic(err)
	}
	service.Log.Info("Service initialized")

	cartFiles := viper.GetStringSlice("cart_files")
	outputDir := viper.GetString("output_dir")

	discountParams := viper.GetStringMap("params.discount")
	taxParams := viper.GetStringMap("params.tax")

	// Initialize CLI Handler
	cliHandler := cli.CLIHandler{
		CartService: service.CartService,
		OutputDir:   outputDir,
	}

	// Process each cart file
	for _, cartFile := range cartFiles {
		fmt.Printf("Processing cart file: %s\n", cartFile)
		if err := cliHandler.Run(cartFile, discountParams, taxParams); err != nil {
			log.Printf("Error processing cart %s: %v", cartFile, err)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			service.Log.Error(r.(string)) //nolint:forcetypeassert // simple type assertion
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	signal := graceful_shutdown.GracefulShutdown()

	cleanup()

	service.Log.Info("Service stopped", field.Fields{
		"signal": signal.String(),
	})

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // exit code 143 is used to indicate graceful termination
}

// initConfig initializes Viper and reads the configuration file.
func initConfig() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // look for config in the working directory
	viper.AddConfigPath("..")     // optionally look for config in the parent directory
	viper.AddConfigPath("./cmd")  // or the cmd directory
	viper.AutomaticEnv()          // read in environment variables that match
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		return fmt.Errorf("fatal error config file: %w", err)
	}
	return nil
}
