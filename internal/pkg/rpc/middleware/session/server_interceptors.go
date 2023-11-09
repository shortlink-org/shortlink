package session_interceptor

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/shortlink-org/shortlink/internal/pkg/auth/session"
)

// SessionUnaryServerInterceptor - extracts user-id from gRPC metadata and adds it to context
func SessionUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, session.ErrMetadataNotFound
		}

		userIDs := md.Get(session.ContextUserIDKey.String())
		if len(userIDs) == 0 {
			return nil, session.ErrUserIDNotFound
		}

		ctx = session.WithUserID(ctx, userIDs[0])

		return handler(ctx, req)
	}
}

// SessionStreamServerInterceptor - extracts user-id from gRPC metadata and adds it to context for streams
func SessionStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv any,
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		wrapped := grpc_middleware.WrapServerStream(stream)
		ctx := wrapped.Context()

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return session.ErrMetadataNotFound
		}

		userIDs := md.Get(session.ContextUserIDKey.String())
		if len(userIDs) == 0 {
			return session.ErrUserIDNotFound
		}

		wrapped.WrappedContext = session.WithUserID(ctx, userIDs[0])

		return handler(srv, wrapped)
	}
}
