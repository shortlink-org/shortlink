package auth

import (
	"context"
)

// insecureMetadataCreds is a gRPC credentials implementation that returns a static set of metadata.
type insecureMetadataCreds map[string]string

// RequireTransportSecurity returns false to indicate that transport security is not required.
func (insecureMetadataCreds) RequireTransportSecurity() bool { return false }

// GetRequestMetadata returns the static metadata.
func (c insecureMetadataCreds) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return c, nil
}
