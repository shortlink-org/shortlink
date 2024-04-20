package link

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorDetail represents a structured error message
type ErrorDetail struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

// ErrorResponse is used to send a detailed error response
type ErrorResponse struct {
	Messages []ErrorDetail `json:"messages"`
}

// ErrMessages creates an ErrorResponse from a given error
func ErrMessages(err error) *ErrorResponse {
	st, ok := status.FromError(err)
	if !ok {
		// If not a gRPC status error, treat as Internal
		return &ErrorResponse{
			Messages: []ErrorDetail{{
				Code: codes.Internal.String(),
				Desc: err.Error(),
			}},
		}
	}

	// Create an ErrorDetail for each gRPC status error
	var details []ErrorDetail
	for _, d := range st.Details() {
		switch t := d.(type) {
		default:
			// Handle other types or log them
			fmt.Printf("Unhandled error type: %T\n", t)
		}
	}
	details = append(details, ErrorDetail{
		Code: st.Code().String(),
		Desc: st.Message(),
	})

	return &ErrorResponse{
		Messages: details,
	}
}
