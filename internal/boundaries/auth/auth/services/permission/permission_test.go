//go:build unit || auth

package permission

import (
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"testing/fstest"
	"time"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/internal/pkg/auth"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// TestGetPermissions tests the GetPermissions function.
func TestGetPermissions(t *testing.T) {
	// Create a mock file system
	mockFS := fstest.MapFS{
		"test1.zed.yaml": &fstest.MapFile{Data: []byte(`
schema: |-
  text: 123
`)},
		"test2.zed.yaml": &fstest.MapFile{Data: []byte(`
schema:
`)},
		"test3.txt": &fstest.MapFile{Data: []byte("content3")},
	}

	permissionsData, err := GetPermissions(mockFS)
	require.NoError(t, err)

	// Expecting 2 files with .zed extension
	require.Len(t, permissionsData, 2)

	// Check the content of the first file
	require.Equal(t, "text: 123", permissionsData[0].Schema)

	// Check the content of the second file
	require.Equal(t, "", permissionsData[1].Schema)
}

func TestSpiceDB(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	client := &Service{}

	// get logger
	log, err := logger.New(logger.Zap, config.Configuration{})

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "authzed/spicedb",
		Tag:          "v1.28.0",
		Cmd:          []string{"serve-testing"},
		ExposedPorts: []string{"50051/tcp"},
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if errRetry := pool.Retry(func() error {
		errSetenv := os.Setenv("GRPC_CLIENT_PORT", resource.GetPort("50051/tcp"))
		require.NoError(t, errSetenv, "Cannot set ENV")

		client.permission, err = auth.New(log, nil, nil)
		require.NoError(t, err, "Cannot create client")

		// wait for the connection to be ready
		time.Sleep(2 * time.Second)

		return nil
	}); errRetry != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		require.NoError(t, errRetry, "Could not connect to docker")
	}

	// test migrations
	t.Run("Migrations", func(t *testing.T) {
		errMigrations := client.Migrations(ctx, permissions)
		require.NoError(t, errMigrations, "Cannot migrate")
	})

	// mock data
	emilia := &pb.SubjectReference{Object: &pb.ObjectReference{
		ObjectType: "user",
		ObjectId:   "emilia",
	}}

	beatrice := &pb.SubjectReference{Object: &pb.ObjectReference{
		ObjectType: "user",
		ObjectId:   "beatrice",
	}}

	firstItem := &pb.ObjectReference{
		ObjectType: "link",
		ObjectId:   "1",
	}

	// test write
	t.Run("Write", func(t *testing.T) {
		request := &pb.WriteRelationshipsRequest{
			Updates: []*pb.RelationshipUpdate{
				{
					// Emilia is a Writer on Post 1
					Operation: pb.RelationshipUpdate_OPERATION_CREATE,
					Relationship: &pb.Relationship{
						Resource: firstItem,
						Relation: "writer",
						Subject:  emilia,
					},
				},
				{
					// Beatrice is a Reader on Post 1
					Operation: pb.RelationshipUpdate_OPERATION_CREATE,
					Relationship: &pb.Relationship{
						Resource: firstItem,
						Relation: "reader",
						Subject:  beatrice,
					},
				},
			},
		}

		_, errWrite := client.permission.WriteRelationships(context.Background(), request)
		require.NoError(t, errWrite)
	})

	// check permissions
	t.Run("CheckPermissions", func(t *testing.T) {
		resp, err := client.permission.CheckPermission(ctx, &pb.CheckPermissionRequest{
			Resource:   firstItem,
			Permission: "view",
			Subject:    emilia,
		})
		require.NoError(t, err, "Cannot check permission")
		require.Equal(t, pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, resp.Permissionship, "Emilia should have view permission")

		resp, err = client.permission.CheckPermission(ctx, &pb.CheckPermissionRequest{
			Resource:   firstItem,
			Permission: "edit",
			Subject:    emilia,
		})
		require.NoError(t, err, "Cannot check permission")
		require.Equal(t, pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, resp.Permissionship, "Emilia should have write permission")

		resp, err = client.permission.CheckPermission(ctx, &pb.CheckPermissionRequest{
			Resource:   firstItem,
			Permission: "view",
			Subject:    beatrice,
		})
		require.NoError(t, err, "Cannot check permission")
		require.Equal(t, pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, resp.Permissionship, "Beatrice should have view permission")

		resp, err = client.permission.CheckPermission(ctx, &pb.CheckPermissionRequest{
			Resource:   firstItem,
			Permission: "edit",
			Subject:    beatrice,
		})
		require.NoError(t, err, "Cannot check permission")
		require.Equal(t, pb.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION, resp.Permissionship, "Beatrice should have write permission")
	})

	// check lookup
	t.Run("Lookup", func(t *testing.T) {
		relationship := &pb.LookupResourcesRequest{
			ResourceObjectType: "link",
			Permission:         "view",
			Subject: &pb.SubjectReference{
				Object: &pb.ObjectReference{
					ObjectType: "user",
					ObjectId:   "emilia",
				},
			},
		}

		stream, err := client.permission.PermissionsServiceClient.LookupResources(ctx, relationship)
		if err != nil {
			t.Fatal(err)
		}

		resources := []*pb.LookupResourcesResponse{}
		for {
			resp, errRead := stream.Recv()
			if errRead != nil {
				if errors.Is(errRead, io.EOF) {
					break
				}

				t.Fatal(errRead)
			}

			resources = append(resources, resp)
		}
		require.Equal(t, 1, len(resources), "Should have 1 resource")
	})

	t.Cleanup(func() {
		// delete all data
		_, errDelete := client.permission.DeleteRelationships(ctx, &pb.DeleteRelationshipsRequest{
			RelationshipFilter: &pb.RelationshipFilter{
				ResourceType: "link",
			},
		})
		require.NoError(t, errDelete, "Cannot delete relationships")

		cancel()

		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}
	})
}
