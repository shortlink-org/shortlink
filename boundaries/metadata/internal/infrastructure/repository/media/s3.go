package s3Repository

import (
	"bytes"
	"context"
	"errors"
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
			return nil, err
		}
	}

	return &Service{
		store: store,
	}, nil
}

func (s *Service) Get(ctx context.Context, linkURL string) (*url.URL, error) {
	// replace characters that are not allowed in the URL
	linkURL = url.PathEscape(linkURL)

	return s.store.GetFileURL(ctx, "screenshot", linkURL)
}

func (s *Service) Put(ctx context.Context, linkURL string, screenshot []byte) error {
	// convert byte slice to io.Reader
	reader := bytes.NewReader(screenshot)

	// replace characters that are not allowed in the URL
	linkURL = url.PathEscape(linkURL)

	err := s.store.UploadFile(ctx, "screenshot", linkURL, reader)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, linkURL string) error {
	err := s.store.RemoveFile(ctx, "screenshot", linkURL)
	if err != nil {
		return err
	}

	return nil
}
