package link

import (
	"strings"

	"github.com/segmentio/encoding/json"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	"github.com/shortlink-org/go-sdk/auth/session"

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
