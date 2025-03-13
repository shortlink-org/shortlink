package session_interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/shortlink-org/shortlink/pkg/auth/session"
)

// SessionUnaryClientInterceptor - set user-id to gRPC metadata for each request
func SessionUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req any,
		resp any,
		clientConn *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		sess, err := session.GetSession(ctx)
		if err != nil {
			return fmt.Errorf("failed to get session: %w", err)
		}

		if sess != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, "user-id", sess.GetId())
		} else {
			userId, err := session.GetUserID(ctx)
			if err != nil {
				return fmt.Errorf("failed to get user id: %w", err)
			}

			ctx = metadata.AppendToOutgoingContext(ctx, "user-id", userId)
		}

		return invoker(ctx, method, req, resp, clientConn, opts...)
	}
}

// SessionStreamClientInterceptor - set user-id to gRPC metadata for each request
func SessionStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		clientConnect *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		sess, err := session.GetSession(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get session: %w", err)
		}

		if sess != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, "user-id", sess.GetId())
		}

		return streamer(ctx, desc, clientConnect, method, opts...)
	}
}
