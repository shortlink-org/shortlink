package rpc

import (
	"fmt"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	grpc_logger "github.com/batazor/shortlink/pkg/rpc/logger"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
)

// InitClient - set up a connection to the server.
func InitClient(log logger.Logger, tracer *opentracing.Tracer) (*grpc.ClientConn, func(), error) {
	viper.SetDefault("GRPC_CLIENT_TLS_ENABLED", false) // gRPC tls
	isEnableTLS := viper.GetBool("GRPC_CLIENT_TLS_ENABLED")

	viper.SetDefault("GRPC_CLIENT_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_CLIENT_PORT")

	viper.SetDefault("GRPC_CLIENT_HOST", "0.0.0.0") // gRPC host
	grpc_host := viper.GetString("GRPC_CLIENT_HOST")

	viper.SetDefault("GRPC_CLIENT_CERT_PATH", "ops/cert/intermediate_ca.pem") // gRPC client cert
	certFile := viper.GetString("GRPC_CLIENT_CERT_PATH")

	viper.SetDefault("GRPC_CLIENT_LOGGER_ENABLE", true) // Enable logging for gRPC-client
	isEnableLogger := viper.GetBool("GRPC_CLIENT_LOGGER_ENABLE")

	// UnaryClien
	var incerceptorUnaryClientList = []grpc.UnaryClientInterceptor{
		grpc_prometheus.UnaryClientInterceptor,
	}

	// StreamClient
	var incerceptorStreamClientList = []grpc.StreamClientInterceptor{
		grpc_prometheus.StreamClientInterceptor,
	}

	if tracer != nil {
		incerceptorUnaryClientList = append(incerceptorUnaryClientList, otgrpc.OpenTracingClientInterceptor(*tracer, otgrpc.LogPayloads()))
		incerceptorStreamClientList = append(incerceptorStreamClientList, otgrpc.OpenTracingStreamClientInterceptor(*tracer, otgrpc.LogPayloads()))
	}

	if isEnableLogger {
		incerceptorUnaryClientList = append(incerceptorUnaryClientList, grpc_logger.UnaryClientInterceptor(log))
		incerceptorStreamClientList = append(incerceptorStreamClientList, grpc_logger.StreamClientInterceptor(log))
	}

	optionsNewClient := []grpc.DialOption{
		// Initialize your gRPC server's interceptor.
		grpc.WithUnaryInterceptor(middleware.ChainUnaryClient(incerceptorUnaryClientList...)),
		grpc.WithStreamInterceptor(middleware.ChainStreamClient(incerceptorStreamClientList...)),
	}
	if isEnableTLS {
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			return nil, nil, err
		}

		optionsNewClient = append(optionsNewClient, grpc.WithTransportCredentials(creds))
	} else {
		optionsNewClient = append(optionsNewClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// Set up a connection to the server peer
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", grpc_host, grpc_port),
		optionsNewClient...,
	)
	if err != nil {
		return nil, nil, err
	}

	log.Info("Run gRPC client", field.Fields{"port": grpc_port, "host": grpc_host})

	cleanup := func() {
		_ = conn.Close()
	}

	return conn, cleanup, nil
}
