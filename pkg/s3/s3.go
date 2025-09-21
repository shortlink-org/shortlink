package s3

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

type Client struct {
	client *minio.Client
}

// New creates a new S3 client
func New(ctx context.Context, log logger.Logger) (*Client, error) {
	viper.AutomaticEnv()
	viper.SetDefault("S3_ENDPOINT", "localhost:9000")
	viper.SetDefault("S3_ACCESS_KEY_ID", "minio_access_key")
	viper.SetDefault("S3_SECRET_ACCESS_KEY", "minio_secret_key")
	viper.SetDefault("S3_USE_SSL", false)

	client, err := minio.New(viper.GetString("S3_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(viper.GetString("S3_ACCESS_KEY_ID"), viper.GetString("S3_SECRET_ACCESS_KEY"), ""),
		Secure: viper.GetBool("S3_USE_SSL"),
	})
	if err != nil {
		return nil, err
	}

	if client.IsOffline() {
		return nil, ErrConnectionFailed
	}

	log.Info("S3 client created", field.Fields{
		"endpoint": viper.GetString("S3_ENDPOINT"),
	})

	return &Client{
		client: client,
	}, nil
}
