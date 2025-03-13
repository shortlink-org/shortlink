package pprof_interceptor

import (
	"context"
	"runtime/pprof"

	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a new unary server interceptors that adds pprof labels.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// Extract method information and create labels
		labels := pprof.Labels("method", info.FullMethod)
		ctx = pprof.WithLabels(ctx, labels)

		pprof.SetGoroutineLabels(ctx)

		return handler(ctx, req)
	}
}

// pprofWrappedServerStream overrides the Context() method to return the labeled ctx
type pprofWrappedServerStream struct {
	grpc.ServerStream
	ctxFunc func() context.Context
}

// StreamServerInterceptor returns a new streaming server interceptor that adds pprof labels.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := stream.Context()

		// Create labels and add them to the context
		labels := pprof.Labels("method", info.FullMethod)
		ctx = pprof.WithLabels(ctx, labels)
		pprof.SetGoroutineLabels(ctx)

		// Wrap the server stream in a custom type that overrides the Context() method
		wrapped := &pprofWrappedServerStream{
			ServerStream: stream,
			ctxFunc:      func() context.Context { return ctx },
		}

		// Proceed with handler and catch any panic
		return handler(srv, wrapped)
	}
}
