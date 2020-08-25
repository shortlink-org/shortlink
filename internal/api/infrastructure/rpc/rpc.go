package rpc

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	metadata "github.com/batazor/shortlink/internal/metadata/domain"
	"github.com/batazor/shortlink/internal/notify"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

type rpc struct {
	client         *grpc.ClientConn
	MetadataClient metadata.MetadataClient
}

func Use(_ context.Context, rpcClient *grpc.ClientConn) (*rpc, error) {
	r := &rpc{
		client: rpcClient,

		// Register clients
		MetadataClient: metadata.NewMetadataClient(rpcClient),
	}

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, r)
	//notify.Subscribe(api_type.METHOD_GET, r)
	//notify.Subscribe(api_type.METHOD_LIST, r)
	//notify.Subscribe(api_type.METHOD_UPDATE, r)
	//notify.Subscribe(api_type.METHOD_DELETE, r)

	return r, nil
}

// Notify ...
func (r *rpc) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		// TODO: do it!!!
		_, err := r.MetadataClient.Set(ctx, &metadata.SetMetaRequest{})
		if err != nil {
		}

		return notify.Response{
			Name:    "RESPONSE_RPC_ADD",
			Payload: payload,
			Error:   errors.New("failed assert type"),
		}
	}

	return notify.Response{}
}
