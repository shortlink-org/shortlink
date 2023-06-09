package auth

import (
	"context"

	"google.golang.org/grpc"
)

// insecureMetadataCreds is a gRPC credentials implementation that returns a static set of metadata.
type insecureMetadataCreds map[string]string

// RequireTransportSecurity returns false to indicate that transport security is not required.
func (c insecureMetadataCreds) RequireTransportSecurity() bool { return false }

// GetRequestMetadata returns the static metadata.
func (c insecureMetadataCreds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return c, nil
}

// WithInsecureBearerToken returns a grpc.DialOption that adds a standard HTTP
// Bearer token to all requests sent from an insecure client.
//
// Must be used in conjunction with `insecure.NewCredentials()`.
func WithInsecureBearerToken(token string) grpc.DialOption {
	return grpc.WithPerRPCCredentials(insecureMetadataCreds{"authorization": "Bearer " + token})
}
