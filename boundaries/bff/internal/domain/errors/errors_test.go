package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorConstructors(t *testing.T) {
	tests := []struct {
		name       string
		fn         func() *Error
		wantCode   string
		wantTitle  string
		wantAction string
	}{
		{
			name:       "SessionNotFound",
			fn:         NewSessionNotFound,
			wantCode:   CodeSessionNotFound,
			wantTitle:  "Session expired",
			wantAction: "LOGIN",
		},
		{
			name:       "UserNotIdentified",
			fn:         NewUserNotIdentified,
			wantCode:   CodeUserNotIdentified,
			wantTitle:  "User not identified",
			wantAction: "LOGIN",
		},
		{
			name:       "SessionMetadataMissing",
			fn:         NewSessionMetadataMissing,
			wantCode:   CodeSessionMetadataMissing,
			wantTitle:  "Authentication metadata missing",
			wantAction: "LOGIN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			require.NotNil(t, err)
			require.Equal(t, tt.wantCode, err.Code)
			require.Equal(t, tt.wantTitle, err.Title)
			require.NotEmpty(t, err.Detail)
			require.Equal(t, tt.wantAction, err.Action)
			require.NoError(t, err.Cause)
		})
	}
}

func TestNewUnknown(t *testing.T) {
	const detail = "unexpected I/O failure"

	err := NewUnknown(detail)

	require.NotNil(t, err)
	require.Equal(t, CodeUnknown, err.Code)
	require.Equal(t, "Unexpected error", err.Title)
	require.Equal(t, detail, err.Detail)
	require.Equal(t, "NONE", err.Action)
	require.NoError(t, err.Cause)
}

func TestWithCause(t *testing.T) {
	cause := errors.New("database timeout")
	base := NewUnknown("operation failed")
	wrapped := base.WithCause(cause)

	require.NotNil(t, wrapped)
	require.Equal(t, base.Code, wrapped.Code)
	require.Equal(t, base.Detail, wrapped.Detail)
	require.Equal(t, base.Title, wrapped.Title)
	require.Equal(t, base.Action, wrapped.Action)
	require.ErrorIs(t, wrapped, cause)
	require.Contains(t, wrapped.Error(), "cause:")
}

func TestIsCode(t *testing.T) {
	tests := []struct {
		name string
		err  *Error
		code string
		want bool
	}{
		{
			name: "Match direct error",
			err:  NewSessionNotFound(),
			code: CodeSessionNotFound,
			want: true,
		},
		{
			name: "Mismatch direct error",
			err:  NewUserNotIdentified(),
			code: CodeSessionNotFound,
			want: false,
		},
		{
			name: "Match wrapped error",
			err:  NewSessionNotFound().WithCause(errors.New("network failure")),
			code: CodeSessionNotFound,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsCode(tt.err, tt.code)
			require.Equal(t, tt.want, got)
		})
	}
}
