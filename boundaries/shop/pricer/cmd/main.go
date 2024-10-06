package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/application"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/infrastructure"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/interfaces/cli"
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
	// Initialize Viper
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Validate configuration
	if err := validateConfig(); err != nil {
		log.Fatalf("Configuration validation error: %v", err)
	}

	// Retrieve configuration values
	discountPolicyPath := viper.GetString("policies.discounts")
	discountQuery := viper.GetString("queries.discounts")

	taxPolicyPath := viper.GetString("policies.taxes")
	taxQuery := viper.GetString("queries.taxes")

	discountParams := viper.GetStringMap("params.discount")
	taxParams := viper.GetStringMap("params.tax")

	cartFiles := viper.GetStringSlice("cart_files")
	outputDir := viper.GetString("output_dir")

	// Initialize Policy Evaluators
	discountEvaluator, err := infrastructure.NewOPAEvaluator(discountPolicyPath, discountQuery)
	if err != nil {
		log.Fatalf("Failed to initialize Discount Policy Evaluator: %v", err)
	}

	taxEvaluator, err := infrastructure.NewOPAEvaluator(taxPolicyPath, taxQuery)
	if err != nil {
		log.Fatalf("Failed to initialize Tax Policy Evaluator: %v", err)
	}

	// Retrieve policy names
	policyNames, err := infrastructure.GetPolicyNames(discountPolicyPath, taxPolicyPath)
	if err != nil {
		log.Fatalf("Failed to retrieve policy names: %v", err)
	}

	// Initialize Cart Service
	cartService := application.NewCartService(discountEvaluator, taxEvaluator, policyNames)

	// Initialize CLI Handler
	cliHandler := cli.CLIHandler{
		CartService: cartService,
		OutputDir:   outputDir,
	}

	// Process each cart file
	for _, cartFile := range cartFiles {
		fmt.Printf("Processing cart file: %s\n", cartFile)
		if err := cliHandler.Run(cartFile, discountParams, taxParams); err != nil {
			log.Printf("Error processing cart %s: %v", cartFile, err)
		}
	}
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
