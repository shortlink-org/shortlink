package http_chi

import (
	"context"

	"github.com/golang/protobuf/jsonpb"
)

// API ...
type API struct { // nolint unused
	ctx    context.Context
	jsonpb jsonpb.Marshaler
}
