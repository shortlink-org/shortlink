package auth

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v3"
)

type Auth struct {
	client *authzed.Client
}

func New() (*Auth, error) {
	var err error
	auth := &Auth{}

	viper.SetDefault("SPICE_DB_API", "shortlink.spicedb:50051")
	viper.SetDefault("SPICE_DB_COMMON_KEY", "secret-shortlink-preshared-key")

	auth.client, err = authzed.NewClient(
		viper.GetString("SPICE_DB_API"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(insecureMetadataCreds{"authorization": "Bearer " + viper.GetString("SPICE_DB_COMMON_KEY")}),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

// Migrations run the migrations for the authzed service.
func (a *Auth) Migrations(ctx context.Context, fs embed.FS) error {
	permissionsData, err := GetPermissions(fs)
	if err != nil {
		return err
	}

	for i := range permissionsData {
		_, err = a.client.WriteSchema(ctx, permissionsData[i])
		if err != nil {
			return fmt.Errorf("Failed to write schema: %w", err)
		}
	}

	return nil
}

// GetPermissions returns a list of permissions from the embedded permissions directory.
func GetPermissions(fsys fs.FS) ([]*pb.WriteSchemaRequest, error) {
	var zedFileData [][]byte

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".yaml" {
			fileData, errReadFile := fs.ReadFile(fsys, path)
			if errReadFile != nil {
				return fmt.Errorf("failed to read file: %w", errReadFile)
			}

			zedFileData = append(zedFileData, fileData)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}

	schemas, err := GetSchema(zedFileData)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema: %w", err)
	}

	return schemas, nil
}

// GetSchema returns a list of schema from the embedded schema directory.
func GetSchema(files [][]byte) ([]*pb.WriteSchemaRequest, error) {
	schemaData := make([]*pb.WriteSchemaRequest, 0, len(files))

	for _, file := range files {
		schema := &pb.WriteSchemaRequest{}

		err := yaml.Unmarshal(file, schema)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal schema: %w", err)
		}

		schemaData = append(schemaData, schema)
	}

	return schemaData, nil
}
