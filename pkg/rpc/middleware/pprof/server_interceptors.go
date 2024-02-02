package pprof_interceptor

import (
	"context"
	"runtime/pprof"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a new unary server interceptors that adds pprof labels.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		// Extract method information and create labels
		labels := pprof.Labels("method", info.FullMethod)
		ctx = pprof.WithLabels(ctx, labels)

		pprof.SetGoroutineLabels(ctx)

		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds pprof labels.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Wrap the server stream and extract the context
		wrapped := grpc_middleware.WrapServerStream(stream)
		ctx := wrapped.Context()

		// Create labels and add them to the context
		labels := pprof.Labels("method", info.FullMethod)
		ctx = pprof.WithLabels(ctx, labels)

		// Update the context of the wrapped stream
		wrapped.WrappedContext = ctx

		pprof.SetGoroutineLabels(ctx)

		// Proceed with handler and catch any panic
		return handler(srv, wrapped)
	}
}
