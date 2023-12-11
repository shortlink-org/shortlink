package http

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"
)

// API ...
type Server struct {
	ctx    context.Context
	jsonpb protojson.MarshalOptions
}
