package s3Repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/shortlink-org/go-sdk/s3"
)

type Service struct {
	store *s3.Client
}

func New(ctx context.Context, store *s3.Client) (*Service, error) {
	// create bucket if not exists
	bucketName := "screenshot"

	err := store.CreateBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check if error is because bucket already exists - this is fine
		var minioErr minio.ErrorResponse
		if errors.As(err, &minioErr) {
			if minioErr.Code == "BucketAlreadyExists" || minioErr.Code == "BucketAlreadyOwnedByYou" {
				// Bucket exists, which is fine - continue without error
				err = nil
			}
		}

		if err != nil {
			return nil, fmt.Errorf("create bucket %s: %w", bucketName, err)
		}
	}

	return &Service{
		store: store,
	}, nil
}

func (s *Service) Get(ctx context.Context, linkURL string) (*url.URL, error) {
	// replace characters that are not allowed in the URL
	linkURL = url.PathEscape(linkURL)

	urlValue, err := s.store.GetFileURL(ctx, "screenshot", linkURL)
	if err != nil {
		return nil, fmt.Errorf("get screenshot url for %s: %w", linkURL, err)
	}

	return urlValue, nil
}

func (s *Service) Put(ctx context.Context, linkURL string, screenshot []byte) error {
	// convert byte slice to io.Reader
	reader := bytes.NewReader(screenshot)

	// replace characters that are not allowed in the URL
	linkURL = url.PathEscape(linkURL)

	err := s.store.UploadFile(ctx, "screenshot", linkURL, reader)
	if err != nil {
		return fmt.Errorf("upload screenshot %s: %w", linkURL, err)
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, linkURL string) error {
	err := s.store.RemoveFile(ctx, "screenshot", linkURL)
	if err != nil {
		return fmt.Errorf("delete screenshot %s: %w", linkURL, err)
	}

	return nil
}
