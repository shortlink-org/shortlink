package config

import (
	"github.com/Unleash/unleash-client-go/v3"
	"github.com/spf13/viper"
)

func (c *Config) FeatureToogleRun() error {
	viper.SetDefault("FEATURE_TOGGLE_ENABLE", false)
	viper.SetDefault("FEATURE_TOGGLE_API", "http://localhost:4242/api/")

	isEnableFeatureToggle := viper.GetBool("FEATURE_TOGGLE_ENABLE")
	if !isEnableFeatureToggle {
		return nil
	}

	err := unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName(viper.GetString("SERVICE_NAME")),
		unleash.WithUrl(viper.GetString("FEATURE_TOGGLE_API")),
		unleash.WithRefreshInterval(10000), // nolint:gomnd
	)
	if err != nil {
		return err
	}

	return nil
}
