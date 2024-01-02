package screenshot

import (
	s3Repository "github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/infrastructure/repository/media"
)

const (
	// These are the default options that are used if no options are specified.
	defaultWidth  = 1920
	defaultHeight = 1080
)

// UC is a use case for screenshot
type UC struct {
	media *s3Repository.Service
}
