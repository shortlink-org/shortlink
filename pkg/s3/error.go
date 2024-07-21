package s3

import (
	"errors"
	"fmt"
)

var ErrConnectionFailed = errors.New("connection to S3 failed")

type CreateBucketError struct {
	Err  error
	Name string
}

func (e CreateBucketError) Error() string {
	return fmt.Sprintf("failed to create bucket %s: %v", e.Name, e.Err)
}
