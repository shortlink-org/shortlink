package session_interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/shortlink-org/shortlink/internal/pkg/auth/session"
)

// SessionUnaryClientInterceptor - set user-id to gRPC metadata for each request
func SessionUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req any,
		resp any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		sess := session.GetSession(ctx)
		if sess != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, "user-id", sess.GetId())
		}

		return invoker(ctx, method, req, resp, cc, opts...)
	}
}

// SessionStreamClientInterceptor - set user-id to gRPC metadata for each request
func SessionStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		sess := session.GetSession(ctx)
		if sess != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, "user-id", sess.GetId())
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}
