package http_chi

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"
)

type API struct {
	ctx    context.Context
	jsonpb protojson.MarshalOptions
}
