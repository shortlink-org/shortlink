package sentry

import (
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/spf13/viper"
)

func New() (*sentryhttp.Handler, func(), error) {
	viper.SetDefault("SENTRY_DSN", "") // key for sentry
	DSN := viper.GetString("SENTRY_DSN")

	if DSN == "" {
		return nil, func() {}, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: DSN,
	})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		// Since sentry emits events in the background we need to make sure
		// they are sent before we shut down
		sentry.Flush(time.Second * 5) // nolint:gomnd
		sentry.Recover()
	}

	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	return sentryHandler, cleanup, nil
}
