package auth

import (
	"context"
)

// insecureMetadataCreds is a gRPC credentials implementation that returns a static set of metadata.
type insecureMetadataCreds map[string]string

// RequireTransportSecurity returns false to indicate that transport security is not required.
func (c insecureMetadataCreds) RequireTransportSecurity() bool { return false }

// GetRequestMetadata returns the static metadata.
func (c insecureMetadataCreds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return c, nil
}
