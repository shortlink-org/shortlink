package v1

import (
	"errors"
)

var ErrParsePayloadAsString = errors.New("error parse payload as string")
var ErrEmptyPayload = errors.New("error empty payload")
