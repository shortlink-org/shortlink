package grpc_logger

import (
	"context"
	"path"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

// UnaryServerInterceptor returns a new unary server interceptors that adds zap.Logger to the context.
func UnaryServerInterceptor(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		startTime := time.Now()
		resp, err := handler(ctx, req)
		duration := time.Since(startTime)

		fields := field.Fields{
			"grpc.service":   path.Dir(info.FullMethod)[1:],
			"grpc.method":    path.Base(info.FullMethod),
			"code":           status.Code(err).String(),
			"duration (mks)": duration.Microseconds(),
		}

		if err != nil {
			printLog(ctx, log, err, fields)
		}

		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds zap.Logger to the context.
func StreamServerInterceptor(log logger.Logger) grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now()
		wrapped := grpc_middleware.WrapServerStream(stream)

		err := handler(srv, wrapped)
		duration := time.Since(startTime)

		fields := field.Fields{
			"grpc.service":   path.Dir(info.FullMethod)[1:],
			"grpc.method":    path.Base(info.FullMethod),
			"code":           status.Code(err).String(),
			"duration (mks)": duration.Microseconds(),
		}

		if err != nil {
			printLog(wrapped.Context(), log, err, fields)
		}

		return err
	}
}
