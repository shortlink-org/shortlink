package http_chi

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"
)

// API ...
type API struct {
	ctx    context.Context
	jsonpb protojson.MarshalOptions
}
