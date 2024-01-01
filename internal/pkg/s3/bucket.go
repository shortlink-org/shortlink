package s3

import (
	"context"

	"github.com/minio/minio-go/v7"
)

// CreateBucket creates a bucket with context
func (c *Client) CreateBucket(ctx context.Context, bucketName string) error {
	err := c.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
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
