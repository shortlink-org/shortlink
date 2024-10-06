package pkg_di

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct{}

// ReadConfig reads the configuration
func ReadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of a config file (without an extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // look for config in the working directory
	viper.AddConfigPath("..")     // optionally look for config in the parent directory
	viper.AddConfigPath("./cmd")  // or the cmd directory
	viper.AutomaticEnv()          // read in environment variables that match
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	return &Config{}, nil
}
