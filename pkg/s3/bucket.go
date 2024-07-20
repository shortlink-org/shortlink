package s3

import (
	"context"

	"github.com/minio/minio-go/v7"
)

// CreateBucket creates a bucket with context
func (c *Client) CreateBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error {
	err := c.client.MakeBucket(ctx, bucketName, opts)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exist, errBucketExists := c.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exist {
			return nil
		}

		return ErrorCreateBucket{
			Err:  err,
			Name: bucketName,
		}
	}

	return nil
}

// RemoveBucket removes a bucket with context
func (c *Client) RemoveBucket(ctx context.Context, bucketName string) error {
	err := c.client.RemoveBucket(ctx, bucketName)
	if err != nil {
		return err
	}

	return nil
}

// BucketExists checks if a bucket exists with context
func (c *Client) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	exists, err := c.client.BucketExists(ctx, bucketName)
	if err != nil {
		return false, err
	}

	return exists, nil
}
