package config

import (
	"github.com/Unleash/unleash-client-go/v4"
	"github.com/spf13/viper"
)

const REFRESH_INTERVAL = 10000

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
		unleash.WithRefreshInterval(REFRESH_INTERVAL),
	)
	if err != nil {
		return err
	}

	return nil
}
