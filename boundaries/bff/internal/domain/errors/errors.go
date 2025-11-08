package errors

import (
	"errors"
	"fmt"
)

// ============================================================
// Codes (ubiquitous language for this domain)
// ============================================================

const (
	CodeSessionNotFound        = "SESSION_NOT_FOUND"
	CodeUserNotIdentified      = "USER_NOT_IDENTIFIED"
	CodeSessionMetadataMissing = "SESSION_METADATA_MISSING"
	CodeUnknown                = "UNKNOWN"
)

// ============================================================
// Error Aggregate Root
// ============================================================

// Error represents a rich domain error that carries meaning
// within the ubiquitous language of the domain.
// It is designed for layering (wrapping causes) and structured reporting.
type Error struct {
	Code   string // domain-specific error code
	Title  string // human-readable short summary
	Detail string // developer-facing detailed message
	Action string // suggested remediation or next step (LOGIN, RETRY, etc.)
	Cause  error  // underlying technical cause (if any)
}

// ============================================================
// Behavior
// ============================================================

// Error implements the standard error interface.
func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Cause == nil {
		return fmt.Sprintf("[%s] %s: %s", e.Code, e.Title, e.Detail)
	}
	return fmt.Sprintf("[%s] %s: %s → cause: %v", e.Code, e.Title, e.Detail, e.Cause)
}

// Unwrap allows integration with errors.Is / errors.As.
func (e *Error) Unwrap() error {
	return e.Cause
}

// WithCause returns a shallow copy with a new cause attached.
// Keeps the domain semantics intact while layering infra or app details.
func (e *Error) WithCause(cause error) *Error {
	if e == nil {
		return nil
	}
	cp := *e
	cp.Cause = cause
	return &cp
}

// IsCode checks if an error (possibly wrapped) matches a given domain code.
func IsCode(err error, code string) bool {
	var de *Error
	if errors.As(err, &de) {
		return de.Code == code
	}
	return false
}

// ============================================================
// Factory Functions (ubiquitous, intention-revealing)
// ============================================================

func NewSessionNotFound() *Error {
	return &Error{
		Code:   CodeSessionNotFound,
		Title:  "Session expired",
		Detail: "Session expired. Please sign in again.",
		Action: "LOGIN",
	}
}

func NewUserNotIdentified() *Error {
	return &Error{
		Code:   CodeUserNotIdentified,
		Title:  "User not identified",
		Detail: "Unable to resolve your account. Please sign in again.",
		Action: "LOGIN",
	}
}

func NewSessionMetadataMissing() *Error {
	return &Error{
		Code:   CodeSessionMetadataMissing,
		Title:  "Authentication metadata missing",
		Detail: "Request is missing authentication metadata.",
		Action: "LOGIN",
	}
}

func NewUnknown(detail string) *Error {
	return &Error{
		Code:   CodeUnknown,
		Title:  "Unexpected error",
		Detail: detail,
		Action: "NONE",
	}
}
