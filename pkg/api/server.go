package api

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/store"
	httpchi "github.com/batazor/shortlink/pkg/api/http-chi"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// runAPIServer - start HTTP-server
func (*Server) RunAPIServer(ctx context.Context, db store.DB) {
	var API API

	viper.SetDefault("API_TYPE", "http-chi")
	viper.SetDefault("API_PORT", 7070)

	// Logger
	log := logger.GetLogger(ctx)

	config := api_type.Config{
		Port: viper.GetInt("API_PORT"),
	}

	serverType := viper.GetString("API_TYPE")

	switch serverType {
	case "http-chi":
		API = &httpchi.API{}
	// case "gRPC-web":
	// 	API = &grpcweb.API{}
	// case "graphql":
	// 	API = &graphql.API{}
	// case "cloudevents":
	// 	API = &cloudevents.API{}
	default:
		API = &httpchi.API{}
	}

	if err := API.Run(ctx, db, config); err != nil {
		log.Fatal(err.Error())
	}
}
