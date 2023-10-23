package grpc_logger

import (
	"context"
	"path"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

// UnaryClientInterceptor returns a new unary client interceptor that optionally logs the execution of external gRPC calls.
func UnaryClientInterceptor(logger logger.Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(startTime)

		fields := field.Fields{
			"grpc.service":   path.Dir(method)[1:],
			"grpc.method":    path.Base(method),
			"code":           status.Code(err).String(),
			"duration (mks)": duration.Microseconds(),
		}

		if err != nil {
			printLog(ctx, logger, err, fields)
		}

		return err
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamClientInterceptor(logger logger.Logger) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {

		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		fields := field.Fields{
			"grpc.service": path.Dir(method)[1:],
			"grpc.method":  path.Base(method),
			"code":         status.Code(err).String(),
		}

		if err != nil {
			printLog(ctx, logger, err, fields)
		}

		return clientStream, err
	}
}
