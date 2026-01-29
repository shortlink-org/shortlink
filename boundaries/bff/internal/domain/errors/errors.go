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
	CodePermissionDenied       = "PERMISSION_DENIED"
	CodeInvalidToken           = "INVALID_TOKEN"
	CodeServiceUnavailable     = "SERVICE_UNAVAILABLE"
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

var (
	// ErrSessionNotFound is a sentinel error for missing user sessions.
	ErrSessionNotFound = &Error{
		Code:   CodeSessionNotFound,
		Title:  "Session expired",
		Detail: "Session expired. Please sign in again.",
		Action: "LOGIN",
	}
	// ErrUserNotIdentified is a sentinel error for anonymous requests.
	ErrUserNotIdentified = &Error{
		Code:   CodeUserNotIdentified,
		Title:  "User not identified",
		Detail: "Unable to resolve your account. Please sign in again.",
		Action: "LOGIN",
	}
	// ErrSessionMetadataMissing is a sentinel error for missing auth metadata.
	ErrSessionMetadataMissing = &Error{
		Code:   CodeSessionMetadataMissing,
		Title:  "Authentication metadata missing",
		Detail: "Request is missing authentication metadata.",
		Action: "LOGIN",
	}
	// ErrPermissionDenied is a sentinel error for access denied.
	ErrPermissionDenied = &Error{
		Code:   CodePermissionDenied,
		Title:  "Access denied",
		Detail: "You don't have permission to perform this action.",
		Action: "NONE",
	}
	// ErrInvalidToken is a sentinel error for invalid authentication tokens.
	ErrInvalidToken = &Error{
		Code:   CodeInvalidToken,
		Title:  "Authentication failed",
		Detail: "Your session has expired or is invalid. Please sign in again.",
		Action: "LOGIN",
	}
	// ErrServiceUnavailable is a sentinel error for service connectivity issues.
	ErrServiceUnavailable = &Error{
		Code:   CodeServiceUnavailable,
		Title:  "Service temporarily unavailable",
		Detail: "We're experiencing technical difficulties. Please try again later.",
		Action: "RETRY",
	}
	// ErrUnknown is a sentinel error for unexpected states.
	ErrUnknown = &Error{
		Code:   CodeUnknown,
		Title:  "Unexpected error",
		Detail: "Something went wrong. Please try again later.",
		Action: "RETRY",
	}
)

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

func NewPermissionDenied() *Error {
	return &Error{
		Code:   CodePermissionDenied,
		Title:  "Access denied",
		Detail: "You don't have permission to perform this action.",
		Action: "NONE",
	}
}

func NewInvalidToken() *Error {
	return &Error{
		Code:   CodeInvalidToken,
		Title:  "Authentication failed",
		Detail: "Your session has expired or is invalid. Please sign in again.",
		Action: "LOGIN",
	}
}

func NewServiceUnavailable() *Error {
	return &Error{
		Code:   CodeServiceUnavailable,
		Title:  "Service temporarily unavailable",
		Detail: "We're experiencing technical difficulties. Please try again later.",
		Action: "RETRY",
	}
}

func NewUnknown(detail string) *Error {
	return &Error{
		Code:   CodeUnknown,
		Title:  "Something went wrong",
		Detail: "An unexpected error occurred. Please try again later.",
		Action: "RETRY",
	}
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

// Is allows errors.Is to compare domain errors by code only.
func (e *Error) Is(target error) bool {
	if e == nil {
		return target == nil
	}

	t, ok := target.(*Error)
	if !ok || t == nil {
		return false
	}

	return e.Code == t.Code
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

// WithDetail returns a shallow copy with an overridden detail message.
func (e *Error) WithDetail(detail string) *Error {
	if e == nil {
		return nil
	}

	cp := *e
	cp.Detail = detail

	return &cp
}

// WithTitle returns a shallow copy with an overridden title.
func (e *Error) WithTitle(title string) *Error {
	if e == nil {
		return nil
	}

	cp := *e
	cp.Title = title

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
