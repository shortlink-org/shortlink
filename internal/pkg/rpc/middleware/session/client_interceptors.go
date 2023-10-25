package session_interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/shortlink-org/shortlink/internal/pkg/auth/session"
)

// SessionUnaryInterceptor - set user-id to gRPC metadata for each request
func SessionUnaryInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req any,
		resp any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) (err error) {

		sess := session.GetSession(ctx)

		ctx = metadata.AppendToOutgoingContext(ctx, "user-id", sess.GetId())

		return invoker(ctx, method, req, resp, cc, opts...)
	}
}

// SessionStreamInterceptor - set user-id to gRPC metadata for each request
func SessionStreamInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {

		sess := session.GetSession(ctx)

		ctx = metadata.AppendToOutgoingContext(ctx, "user-id", sess.GetId())

		return streamer(ctx, desc, cc, method, opts...)
	}
}
