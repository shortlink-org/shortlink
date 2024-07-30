package temporal

import (
	"go.temporal.io/sdk/client"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

func New(log logger.Logger) (client.Client, error) {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		return nil, err
	}

	log.Info("Temporal client created")

	return c, nil
}
