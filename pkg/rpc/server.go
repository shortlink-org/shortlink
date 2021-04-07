package rpc

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
)

// InitServer ...
func InitServer(log logger.Logger, tracer *opentracing.Tracer) (*RPCServer, func(), error) {
	viper.SetDefault("GRPC_SERVER_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_SERVER_PORT")

	viper.SetDefault("GRPC_SERVER_CERT_PATH", "ops/cert/shortlink-server.pem") // gRPC server cert
	certFile := viper.GetString("GRPC_SERVER_CERT_PATH")
	viper.SetDefault("GRPC_SERVER_KEY_PATH", "ops/cert/shortlink-server-key.pem") // gRPC server key
	keyFile := viper.GetString("GRPC_SERVER_KEY_PATH")

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("0.0.0.0:%d", grpc_port)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, nil, err
	}

	// Define custom func to handle panic
	customFunc := func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	// Shared options for the logger, with a custom gRPC code to log level function.
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(customFunc),
	}

	// UnaryServer
	var incerceptorUnaryServerList = []grpc.UnaryServerInterceptor{
		grpc_prometheus.UnaryServerInterceptor,

		// Create a server. Recovery handlers should typically be last in the chain so that other middleware
		// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
		recovery.UnaryServerInterceptor(recoveryOpts...),
	}

	if tracer != nil {
		incerceptorUnaryServerList = append(incerceptorUnaryServerList, otgrpc.OpenTracingServerInterceptor(*tracer, otgrpc.LogPayloads()))
	}

	// StreamClient
	var incerceptorStreamServerList = []grpc.StreamServerInterceptor{
		grpc_prometheus.StreamServerInterceptor,

		// Create a server. Recovery handlers should typically be last in the chain so that other middleware
		// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
		recovery.StreamServerInterceptor(recoveryOpts...),
	}

	if tracer != nil {
		incerceptorStreamServerList = append(incerceptorStreamServerList, otgrpc.OpenTracingStreamServerInterceptor(*tracer, otgrpc.LogPayloads()))
	}

	// Initialize the gRPC server.
	rpc := grpc.NewServer(
		grpc.Creds(creds),

		// Initialize your gRPC server's interceptor.
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(incerceptorUnaryServerList...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(incerceptorStreamServerList...)),
	)

	r := &RPCServer{
		Server: rpc,
		Run: func() {
			// Register reflection service on gRPC server.
			reflection.Register(rpc)

			// After all your registrations, make sure all of the Prometheus metrics are initialized.
			grpc_prometheus.Register(rpc)

			log.Info("Run gRPC server", field.Fields{"port": grpc_port})
			err = rpc.Serve(lis)
			if err != nil {
				log.Error(err.Error())
			}
		},
		Endpoint: endpoint,
	}

	cleanup := func() {
		rpc.GracefulStop()
	}

	return r, cleanup, err
}
