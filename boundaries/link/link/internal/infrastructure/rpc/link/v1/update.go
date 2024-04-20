package v1

import (
	"context"
)

func (l *LinkRPC) Update(ctx context.Context, in *UpdateRequest) (*UpdateResponse, error) {
	// resp, err := l.service.Update(ctx, in.GetLink())
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }

	return &UpdateResponse{
		// Link: resp,
	}, nil
}
