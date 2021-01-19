package config

import (
	"github.com/Unleash/unleash-client-go/v3"
	"github.com/spf13/viper"
)

func (c *Config) FeatureToogleRun() error {
	viper.SetDefault("FEATURE_TOGGLE_API", "http://localhost:4242/api/")

	err := unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName(viper.GetString("SERVICE_NAME")),
		unleash.WithUrl(viper.GetString("FEATURE_TOGGLE_API")),
		unleash.WithRefreshInterval(10000),
	)
	if err != nil {
		return err
	}

	return nil
}
