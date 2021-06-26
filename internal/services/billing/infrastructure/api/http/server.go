package api

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/logger"
	http_chi "github.com/batazor/shortlink/internal/services/api/application/http-chi"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
)

// API - general describe of API
type API interface { // nolint unused
	Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer *opentracing.Tracer) error
}

type Server struct{}

func (s *Server) Use(ctx context.Context, log logger.Logger, tracer *opentracing.Tracer) (*Server, error) {
	var server API

	viper.SetDefault("API_TYPE", "http-chi") // Select: http-chi
	viper.SetDefault("API_PORT", 7070)       // API port
	viper.SetDefault("API_TIMEOUT", 60)      // Request Timeout (seconds)

	config := api_type.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT") * time.Second, // nolint durationcheck
	}

	serverType := viper.GetString("API_TYPE")

	switch serverType {
	case "http-chi":
		server = &http_chi.API{}
	default:
		server = &http_chi.API{}
	}

	if err := server.Run(ctx, config, log, tracer); err != nil {
		return nil, err
	}

	return s, nil
}
