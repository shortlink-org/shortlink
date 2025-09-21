package ws

import (
	"context"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

type WS struct{}

func (ws *WS) Run(ctx context.Context, log logger.Logger) (*WS, error) {
	viper.SetDefault("BASE_PATH", "/ws") // Base path for WS endpoints

	server := &WS{}

	g := errgroup.Group{}

	g.Go(func() error {
		http.HandleFunc(viper.GetString("BASE_PATH"), Handler)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			return err
		}

		return nil
	})

	log.Info("WS server started", field.Fields{
		"port":      8080,
		"base_path": viper.GetString("BASE_PATH"),
	})

	return server, nil
}
