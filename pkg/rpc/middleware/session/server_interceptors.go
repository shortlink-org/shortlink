package session_interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/shortlink-org/shortlink/pkg/auth/session"
)

// SessionUnaryServerInterceptor - extracts user-id from gRPC metadata and adds it to context
func SessionUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		// Skip user-id extraction for ServerReflectionInfo method
		if info.FullMethod == "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo" {
			return handler(ctx, req)
		}

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

// sessionWrappedServerStream overrides Context() to return the derived context
type sessionWrappedServerStream struct {
	grpc.ServerStream
	ctxFunc func() context.Context
}

// SessionStreamServerInterceptor - extracts user-id from gRPC metadata and adds it to context for streams
func SessionStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv any,
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		// Skip user-id extraction for ServerReflectionInfo method
		if info.FullMethod == "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo" {
			return handler(srv, stream)
		}

		ctx := stream.Context()

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return session.ErrMetadataNotFound
		}

		userIDs := md.Get(session.ContextUserIDKey.String())
		if len(userIDs) == 0 {
			return session.ErrUserIDNotFound
		}

		// Create a derived context with user-id
		newCtx := session.WithUserID(ctx, userIDs[0])

		// Wrap the stream but store only a function that returns newCtx
		wrapped := &sessionWrappedServerStream{
			ServerStream: stream,
			ctxFunc:      func() context.Context { return newCtx },
		}

		// Invoke the handler with our wrapped stream
		return handler(srv, wrapped)
	}
}
