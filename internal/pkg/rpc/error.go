package rpc

import (
	"errors"
)

var ErrGetMetadataFromContext = errors.New("error get metadata from context")
var ErrGetSessionFromMetadata = errors.New("error get session from metadata")
