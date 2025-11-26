package link

import (
	"net/http"
	"strings"

	"github.com/segmentio/encoding/json"
	"github.com/shortlink-org/go-sdk/auth/session"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	domainerrors "github.com/shortlink-org/shortlink/boundaries/link/bff/internal/domain/errors"
)

// ============================================================
// Response Models
// ============================================================

// ErrorDetail represents a structured error message for clients (JSON response)
type ErrorDetail struct {
	Code     string         `json:"code"`
	Title    string         `json:"title"`
	Detail   string         `json:"detail"`
	Action   string         `json:"action,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// ErrorResponse is used to send a detailed structured error payload
type ErrorResponse struct {
	Messages []ErrorDetail `json:"messages"`
}

// ============================================================
// Conversion Logic
// ============================================================

// ErrMessages builds an ErrorResponse from any given error (domain or gRPC)
func ErrMessages(err error) *ErrorResponse {
	st, ok := status.FromError(err)
	if !ok {
		// Not a gRPC status — wrap as unknown domain error
		domainErr := domainerrors.NewUnknown(err.Error())

		return &ErrorResponse{
			Messages: []ErrorDetail{{
				Code:   domainErr.Code,
				Title:  domainErr.Title,
				Detail: domainErr.Detail,
				Action: domainErr.Action,
			}},
		}
	}

	// Map gRPC status to domain error
	domainErr := mapStatusToResponse(st)
	metadata := map[string]any{
		"grpc_status_code": st.Code().String(),
	}

	// Include any additional gRPC error details (for diagnostics)
	var grpcDetails []string

	for _, d := range st.Details() {
		raw, err := json.Marshal(d)
		if err != nil {
			continue
		}

		grpcDetails = append(grpcDetails, string(raw))
	}

	if len(grpcDetails) > 0 {
		metadata["grpc_details"] = grpcDetails
	}

	return &ErrorResponse{
		Messages: []ErrorDetail{{
			Code:     domainErr.Code,
			Title:    domainErr.Title,
			Detail:   domainErr.Detail,
			Action:   domainErr.Action,
			Metadata: metadata,
		}},
	}
}

// ============================================================
// Mapper from gRPC status to Domain Error
// ============================================================

func mapStatusToResponse(st *status.Status) *domainerrors.Error {
	// Check structured gRPC details (errdetails.ErrorInfo)
	for _, d := range st.Details() {
		if info, ok := d.(*errdetails.ErrorInfo); ok {
			switch info.GetReason() {
			case domainerrors.CodeSessionNotFound:
				return domainerrors.NewSessionNotFound()
			case domainerrors.CodeUserNotIdentified:
				return domainerrors.NewUserNotIdentified()
			case domainerrors.CodeSessionMetadataMissing:
				return domainerrors.NewSessionMetadataMissing()
			}
		}
	}

	message := st.Message()

	// Fallback to string-based inference (for backward compatibility)
	switch {
	case strings.Contains(message, session.ErrSessionNotFound.Error()):
		return domainerrors.NewSessionNotFound()
	case strings.Contains(message, session.ErrUserIDNotFound.Error()):
		return domainerrors.NewUserNotIdentified()
	case strings.Contains(message, session.ErrMetadataNotFound.Error()):
		return domainerrors.NewSessionMetadataMissing()
	default:
		return domainerrors.NewUnknown(message)
	}
}

// ============================================================
// Mapper from gRPC status codes to HTTP status codes
// ============================================================

// grpcStatusToHTTP maps gRPC status codes to HTTP status codes
func grpcStatusToHTTP(err error) int {
	st, ok := status.FromError(err)
	if !ok {
		// Not a gRPC status error, return 500
		return http.StatusInternalServerError
	}

	switch st.Code() {
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusRequestTimeout
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
