package s3

import (
	"errors"
)

var ErrConnectionFailed = errors.New("connection to S3 failed")
