/*
Config package
*/
package config

import (
	"errors"

	"github.com/spf13/viper"

	error_di "github.com/shortlink-org/shortlink/pkg/di/pkg/error"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

type Config struct{}

// Init - read .env and ENV variables
func New(log logger.Logger) (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".") // look for config in the working directory
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var typeErr viper.ConfigFileNotFoundError
		if !errors.As(err, &typeErr) {
			return nil, &error_di.BaseError{Err: err}
		}

		log.Warn("The .env file has not been found in the current directory")
	}

	config := &Config{}

	// Enable feature toggle
	err := config.FeatureToogleRun()
	if err != nil {
		return nil, &error_di.BaseError{Err: err}
	}

	return config, nil
}
