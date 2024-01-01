package media

import (
	"bytes"
	"context"
	"net/url"

	"github.com/minio/minio-go/v7"

	"github.com/shortlink-org/shortlink/internal/pkg/s3"
)

type Service struct {
	store *s3.Client
}

func New(ctx context.Context, store *s3.Client) (*Service, error) {
	// create bucket if not exists
	err := store.CreateBucket(ctx, "screenshot", minio.MakeBucketOptions{})
	if err != nil {
		return nil, err
	}

	return &Service{
		store: store,
	}, nil
}

func (s *Service) Get(ctx context.Context, linkID string) (*url.URL, error) {
	return s.store.GetFileURL(ctx, "screenshot", linkID)
}

func (s *Service) Put(ctx context.Context, linkID string, screenshot []byte) error {
	// convert byte slice to io.Reader
	reader := bytes.NewReader(screenshot)

	err := s.store.UploadFile(ctx, "screenshot", linkID, reader)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, linkID string) error {
	err := s.store.RemoveFile(ctx, "screenshot", linkID)
	if err != nil {
		return err
	}

	return nil
}
