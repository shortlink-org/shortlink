package link

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shortlink-org/go-sdk/auth/session"

	domainerrors "github.com/shortlink-org/shortlink/boundaries/link/bff/internal/domain/errors"
)

func TestMapStatusToResponse_UsesErrorInfoReason(t *testing.T) {
	st := status.New(codes.InvalidArgument, "session missing")
	info := &errdetails.ErrorInfo{Reason: domainerrors.CodeSessionNotFound}

	stWithDetails, err := st.WithDetails(info)
	require.NoError(t, err, "failed to attach details")

	result := mapStatusToResponse(stWithDetails)

	require.NotNil(t, result)
	require.Equal(t, domainerrors.CodeSessionNotFound, result.Code)
	require.Equal(t, "Session expired", result.Title)
	require.NotEmpty(t, result.Detail)
	require.Equal(t, "LOGIN", result.Action)
}

func TestMapStatusToResponse_FallbackToMessage(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{
			name:     "session not found",
			message:  session.ErrSessionNotFound.Error(),
			expected: domainerrors.CodeSessionNotFound,
		},
		{
			name:     "user id not found",
			message:  session.ErrUserIDNotFound.Error(),
			expected: domainerrors.CodeUserNotIdentified,
		},
		{
			name:     "metadata missing",
			message:  session.ErrMetadataNotFound.Error(),
			expected: domainerrors.CodeSessionMetadataMissing,
		},
		{
			name:     "unknown message",
			message:  "unexpected",
			expected: domainerrors.CodeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := status.New(codes.InvalidArgument, tt.message)
			result := mapStatusToResponse(st)
			require.Equal(t, tt.expected, result.Code)
		})
	}
}

func TestErrMessages_BuildsUnifiedContract(t *testing.T) {
	st := status.New(codes.InvalidArgument, session.ErrSessionNotFound.Error())

	response := ErrMessages(st.Err())

	require.Len(t, response.Messages, 1)

	msg := response.Messages[0]
	require.Equal(t, domainerrors.CodeSessionNotFound, msg.Code)
	require.NotEmpty(t, msg.Title)
	require.NotEmpty(t, msg.Detail)
	require.Equal(t, "LOGIN", msg.Action)

	require.NotNil(t, msg.Metadata)
	require.Equal(t, codes.InvalidArgument.String(), msg.Metadata["grpc_status_code"])
}

func TestErrMessages_IncludesGrpcDetails(t *testing.T) {
	st := status.New(codes.InvalidArgument, session.ErrSessionNotFound.Error())
	info := &errdetails.ErrorInfo{Reason: domainerrors.CodeSessionNotFound}

	stWithDetails, err := st.WithDetails(info)
	require.NoError(t, err, "failed to attach details")

	response := ErrMessages(stWithDetails.Err())
	require.Len(t, response.Messages, 1)

	msg := response.Messages[0]
	require.Contains(t, msg.Metadata, "grpc_details")

	details, ok := msg.Metadata["grpc_details"].([]string)
	require.True(t, ok, "grpc_details should be a []string")
	require.NotEmpty(t, details, "grpc_details should not be empty")
}
