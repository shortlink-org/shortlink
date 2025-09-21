/*
Config package
*/
package config

import (
	"errors"
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	mu sync.RWMutex
}

// New - read .env and ENV variables.
func New() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".") // look for config in the working directory
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		var typeErr viper.ConfigFileNotFoundError
		if !errors.As(err, &typeErr) {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	config := &Config{
		mu: sync.RWMutex{},
	}

	// Enable feature toggle
	err = config.FeatureToogleRun()
	if err != nil {
		return nil, err
	}

	return config, nil
}