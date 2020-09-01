/*
Config package
*/
package config

import (
	"github.com/spf13/viper"
)

// Init - read .env and ENV variables
func Init() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// TODO: logger this fact
			//return errors.New("The .env file has not been found in the current directory")
			return nil
		} else {
			return err
		}
	}

	return nil
}
