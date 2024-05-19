/*
Config package
*/
package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct{}

// Init - read .env and ENV variables
func New() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var typeErr viper.ConfigFileNotFoundError
		if errors.As(err, &typeErr) {
			// TODO: logger this fact
			// return errors.New("The .env file has not been found in the current directory")
			return nil, nil
		}

		return nil, err
	}

	config := &Config{}

	// Enable feature toggle
	err := config.FeatureToogleRun()
	if err != nil {
		return nil, err
	}

	return config, nil
}
