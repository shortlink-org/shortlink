package api

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/logger"
	api_rpc "github.com/batazor/shortlink/internal/services/api/infrastructure/rpc"
	"github.com/batazor/shortlink/pkg/api/cloudevents"
	gokit "github.com/batazor/shortlink/pkg/api/go-kit"
	"github.com/batazor/shortlink/pkg/api/graphql"
	grpcweb "github.com/batazor/shortlink/pkg/api/grpc-web"
	httpchi "github.com/batazor/shortlink/pkg/api/http-chi"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/rpc"
)

// runAPIServer - start HTTP-server
func (*Server) RunAPIServer(ctx context.Context, log logger.Logger, tracer *opentracing.Tracer, rpcServer *rpc.RPCServer, rpcClient *grpc.ClientConn) {
	var server API

	viper.SetDefault("API_TYPE", "http-chi")        // Select: http-chi, gRPC-web, graphql, cloudevents, go-kit
	viper.SetDefault("API_PORT", 7070)              // API port
	viper.SetDefault("API_TIMEOUT", 60*time.Second) // Request Timeout

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

	// Register clients
	_, err := api_rpc.Use(ctx, rpcClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := server.Run(ctx, config, log, tracer); err != nil {
		log.Fatal(err.Error())
	}
}
