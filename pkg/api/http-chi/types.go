package http_chi

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"
)

// API ...
type API struct { // nolint unused
	ctx    context.Context
	jsonpb protojson.MarshalOptions
}
