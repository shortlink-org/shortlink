package grpc_logger

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

func printLog(ctx context.Context, log logger.Logger, err error, fields field.Fields) {
	switch status.Code(err) {
	case codes.OK, codes.Canceled, codes.InvalidArgument, codes.NotFound, codes.AlreadyExists, codes.ResourceExhausted, codes.FailedPrecondition, codes.Aborted, codes.OutOfRange: //nolint:lll
		log.DebugWithContext(ctx, err.Error(), fields)
	case codes.Unknown, codes.DeadlineExceeded, codes.PermissionDenied, codes.Unauthenticated:
		log.InfoWithContext(ctx, err.Error(), fields)
	case codes.Unimplemented, codes.Internal, codes.Unavailable, codes.DataLoss:
		log.WarnWithContext(ctx, err.Error(), fields)
	default:
		log.InfoWithContext(ctx, err.Error(), fields)
	}
}
