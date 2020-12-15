package rpc

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/logger/field"
)

// runGRPCServer ...
func RunGRPCServer(log logger.Logger, tracer opentracing.Tracer) (*RPCServer, func(), error) {
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

	// Initialize the gRPC server.
	rpc := grpc.NewServer(
		grpc.Creds(creds),

		// Initialize your gRPC server's interceptor.
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_prometheus.UnaryServerInterceptor,
		)),

		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_prometheus.StreamServerInterceptor,
		)),
	)

	r := &RPCServer{
		Server: rpc,
		Run: func() {
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

// runGRPCClient - set up a connection to the server.
func RunGRPCClient(log logger.Logger, tracer opentracing.Tracer) (*grpc.ClientConn, func(), error) {
	viper.SetDefault("GRPC_CLIENT_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_CLIENT_PORT")

	viper.SetDefault("GRPC_CLIENT_CERT_PATH", "ops/cert/intermediate_ca.pem") // gRPC client cert
	certFile := viper.GetString("GRPC_CLIENT_CERT_PATH")

	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return nil, nil, err
	}

	// Set up a connection to the server peer
	conn, err := grpc.Dial(
		fmt.Sprintf("0.0.0.0:%d", grpc_port),
		grpc.WithTransportCredentials(creds),

		// Initialize your gRPC server's interceptor.
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_prometheus.UnaryClientInterceptor,
		)),

		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			otgrpc.OpenTracingStreamClientInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_prometheus.StreamClientInterceptor,
		)),
	)
	if err != nil {
		return nil, nil, err
	}

	log.Info("Run gRPC client", field.Fields{"port": grpc_port})

	cleanup := func() {
		conn.Close()
	}

	return conn, cleanup, nil
}
