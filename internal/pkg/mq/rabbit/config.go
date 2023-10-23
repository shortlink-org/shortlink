package rabbit

import (
	"github.com/spf13/viper"
)

type Config struct {
	URI           string
	ReconnectTime int
}

// setConfig - Construct a new RabbitMQ configuration.
func loadConfig() *Config {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_RABBIT_URI", "amqp://localhost:5672") // RabbitMQ URI
	// RabbitMQ reconnects after delay seconds
	viper.SetDefault("MQ_RECONNECT_DELAY_SECONDS", 3) //nolint:gomnd

	return &Config{
		URI:           viper.GetString("MQ_RABBIT_URI"),
		ReconnectTime: viper.GetInt("MQ_RECONNECT_DELAY_SECONDS"),
	}
}
