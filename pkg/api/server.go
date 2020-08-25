package api

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/pkg/api/cloudevents"
	gokit "github.com/batazor/shortlink/pkg/api/go-kit"
	"github.com/batazor/shortlink/pkg/api/graphql"
	grpcweb "github.com/batazor/shortlink/pkg/api/grpc-web"
	httpchi "github.com/batazor/shortlink/pkg/api/http-chi"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// runAPIServer - start HTTP-server
func (*Server) RunAPIServer(ctx context.Context, log logger.Logger, tracer opentracing.Tracer, rpcServer *grpc.Server) {
	var server API

	viper.SetDefault("API_TYPE", "http-chi") // Select: http-chi, gRPC-web, graphql, cloudevents, go-kit
	viper.SetDefault("API_PORT", 7070)       // API port
	viper.SetDefault("API_TIMEOUT", 60)      // Request Timeout

	config := api_type.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	serverType := viper.GetString("API_TYPE")

	switch serverType {
	case "http-chi":
		server = &httpchi.API{}
	case "go-kit":
		server = &gokit.API{}
	case "gRPC-web":
		server = &grpcweb.API{
			RPC: rpcServer,
		}
	case "graphql":
		server = &graphql.API{}
	case "cloudevents":
		server = &cloudevents.API{}
	default:
		server = &httpchi.API{}
	}

	if err := server.Run(ctx, config, log, tracer); err != nil {
		log.Fatal(err.Error())
	}
}
