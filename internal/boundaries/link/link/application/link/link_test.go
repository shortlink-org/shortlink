//go:build unit

//go:generate go run github.com/vektra/mockery/v2

package link

import (
	"context"
	"errors"
	"io"
	"testing"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	mockPermission "github.com/shortlink-org/shortlink/internal/boundaries/link/link/application/link/mocks/authzed"
	"github.com/shortlink-org/shortlink/internal/boundaries/link/link/application/link/mocks/crud"
	"github.com/shortlink-org/shortlink/internal/boundaries/link/link/application/link/mocks/metadata"
	"github.com/shortlink-org/shortlink/internal/boundaries/link/link/application/link/mocks/mq"
	v1 "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	metadata_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/infrastructure/rpc/metadata/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
)

func TestLinkService(t *testing.T) {
	log, err := logger.NewLogger(logger.Zap, config.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	})
	require.NoError(t, err, "Error init a logger")

	// This is the mock repository that we will use to test our service
	mockRepository := new(crud.MockRepository)
	mockMQ := new(mq.MockMQ)
	mockPermissionsServiceClient := new(mockPermission.MockPermissionsServiceClient)
	mockMetadata := new(metadata.MockMetadataServiceClient)

	linkService, err := New(log, mockMQ, mockMetadata, mockRepository, &authzed.Client{
		SchemaServiceClient:      nil,
		PermissionsServiceClient: mockPermissionsServiceClient,
		WatchServiceClient:       nil,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Add", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockRepository.On("Add", mock.Anything, mock.Anything).Return(&v1.Link{}, nil).Once()
			mockMQ.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			mockPermissionsServiceClient.On("WriteRelationships", mock.Anything, mock.Anything).Return(&permission.WriteRelationshipsResponse{}, nil).Once()
			mockMetadata.On("Set", mock.Anything, mock.Anything).Return(&metadata_rpc.MetadataServiceSetResponse{}, nil).Once()

			_, err := linkService.Add(context.Background(), &v1.Link{})
			assert.NoError(t, err)
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockRepository.On("Get", mock.Anything, mock.Anything).Return(&v1.Link{}, nil).Once()
			mockPermissionsServiceClient.On("CheckPermission", mock.Anything, mock.Anything).Return(&permission.CheckPermissionResponse{}, nil).Once()

			_, err := linkService.Get(context.Background(), "somehash")
			assert.NoError(t, err)
		})

		t.Run("Permission Denied", func(t *testing.T) {
			mockPermissionsServiceClient.On("CheckPermission", mock.Anything, mock.Anything).Return(nil, errors.New("permission denied")).Once()

			_, err := linkService.Get(context.Background(), "somehash")
			assert.Error(t, err)
		})

		t.Run("Link Not Found", func(t *testing.T) {
			mockRepository.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New("not found")).Once()
			mockPermissionsServiceClient.On("CheckPermission", mock.Anything, mock.Anything).Return(&permission.CheckPermissionResponse{}, nil).Once()

			_, err := linkService.Get(context.Background(), "somehash")
			assert.Error(t, err)
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockPermissionsService_LookupResourcesClient := new(mockPermission.MockPermissionsService_LookupResourcesClient)
			mockPermissionsService_LookupResourcesClient.On("Recv").Return(&permission.LookupResourcesResponse{}, nil).Times(2)
			mockPermissionsService_LookupResourcesClient.On("Recv").Return(nil, io.EOF).Once()

			mockRepository.On("List", mock.Anything, mock.Anything).Return(&v1.Links{}, nil)
			mockPermissionsServiceClient.On("LookupResources", mock.Anything, mock.Anything).Return(func(context.Context, *permission.LookupResourcesRequest, ...grpc.CallOption) permission.PermissionsService_LookupResourcesClient {
				return mockPermissionsService_LookupResourcesClient
			}, nil).Once()

			resp, err := linkService.List(context.Background(), &v1.FilterLink{})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})

		t.Run("Permission Denied", func(t *testing.T) {
			mockPermissionsServiceClient.On("LookupResources", mock.Anything, mock.Anything).Return(nil, errors.New("permission denied")).Once()

			_, err := linkService.List(context.Background(), &v1.FilterLink{})
			assert.Error(t, err)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
			mockPermissionsServiceClient.On("DeleteRelationships", mock.Anything, mock.Anything).Return(&permission.DeleteRelationshipsResponse{}, nil).Once()

			_, err := linkService.Delete(context.Background(), "somehash")
			assert.NoError(t, err)
		})

		t.Run("Permission Denied", func(t *testing.T) {
			mockPermissionsServiceClient.On("DeleteRelationships", mock.Anything, mock.Anything).Return(nil, errors.New("permission denied")).Once()

			_, err := linkService.Delete(context.Background(), "somehash")
			assert.Error(t, err)
		})
	})

	t.Run("Update", func(t *testing.T) {
		// Assuming you implement the Update method in your service
		t.Run("Success", func(t *testing.T) {
			mockRepository.On("Update", mock.Anything, mock.Anything).Return(&v1.Link{}, nil).Once()

			_, err := linkService.Update(context.Background(), &v1.Link{})
			assert.NoError(t, err)
		})
	})
}
