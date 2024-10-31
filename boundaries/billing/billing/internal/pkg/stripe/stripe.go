package provider

import (
	"context"

	"github.com/stripe/stripe-go/v81/client"
)

type Provider struct {
	client *client.API
}

func New(ctx context.Context) *Provider {
	// Default values
	viper.SetDefault("STRIPE_SECRET_KEY", "secret_key")

	config := &stripe.BackendConfig{
		MaxNetworkRetries: stripe.Int64(3),
	}

	// Setup
	sc := &client.API{}
	sc.Init(viper.GetString("STRIPE_SECRET_KEY"), &stripe.Backends{
		API:             stripe.GetBackendWithConfig(stripe.APIBackend, config),
		Uploads:         stripe.GetBackendWithConfig(stripe.UploadsBackend, config),
		EnableTelemetry: stripe.Bool(true),
	})

	return &Provider{
		client: sc,
	}
}
