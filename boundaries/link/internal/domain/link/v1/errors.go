package v1

import "fmt"

type Code string

const (
	CodeNotFound         Code = "LINK_NOT_FOUND"
	CodeInvalidInput     Code = "LINK_INVALID_INPUT"
	CodeConflict         Code = "LINK_CONFLICT"
	CodePermissionDenied Code = "LINK_PERMISSION_DENIED"
	CodeInternal         Code = "LINK_INTERNAL"
)

type LinkError struct {
	code    Code
	message string
	cause   error
}

// NewLinkError creates a new LinkError with the given code, message, and cause.
func NewLinkError(code Code, message string, cause error) *LinkError {
	return &LinkError{
		code:    code,
		message: message,
		cause:   cause,
	}
}

func newLinkError(code Code, message string, cause error) *LinkError {
	return NewLinkError(code, message, cause)
}

func (e *LinkError) Error() string {
	if e == nil {
		return "link error"
	}

	if e.message != "" {
		return e.message
	}

	return fmt.Sprintf("link error: %s", e.code)
}

func (e *LinkError) Code() Code {
	if e == nil {
		return ""
	}

	return e.code
}

func (e *LinkError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.cause
}

func (e *LinkError) Is(target error) bool {
	if e == nil {
		return false
	}

	t, ok := target.(*LinkError)
	if !ok || t == nil {
		return false
	}

	return e.code == t.code
}

// ErrNotFound indicates the link with the provided hash does not exist.
func ErrNotFound(hash string) *LinkError {
	message := "link not found"
	if hash != "" {
		message = fmt.Sprintf("link not found: hash=%s", hash)
	}

	return newLinkError(CodeNotFound, message, nil)
}

// ErrInvalidInput represents validation errors in the domain layer.
func ErrInvalidInput(message string) *LinkError {
	if message == "" {
		message = "invalid input"
	}

	return newLinkError(CodeInvalidInput, "invalid input: "+message, nil)
}

// ErrConflict represents business logic conflicts (duplicates, constraints, etc.).
func ErrConflict(reason string) *LinkError {
	if reason == "" {
		reason = "unknown reason"
	}

	return newLinkError(CodeConflict, "link conflict: "+reason, nil)
}

// ErrPermissionDenied indicates the caller lacks required permissions.
func ErrPermissionDenied(cause error) *LinkError {
	return newLinkError(CodePermissionDenied, "permission denied", cause)
}

// ErrInternal wraps unexpected domain or infrastructure errors.
func ErrInternal(message string, cause error) *LinkError {
	if message == "" {
		message = "internal error"
	} else {
		message = "internal error: " + message
	}

	return newLinkError(CodeInternal, message, cause)
}

// NotFoundError indicates the link with the provided hash does not exist.
type NotFoundError struct {
	Hash string
}

func (e *NotFoundError) Error() string {
	if e == nil || e.Hash == "" {
		return "link not found"
	}

	return "link not found: hash=" + e.Hash
}

// NewNotFoundError creates a new NotFoundError.
func NewNotFoundError(hash string) *NotFoundError {
	return &NotFoundError{Hash: hash}
}

// InvalidInputError represents validation errors in the domain layer.
type InvalidInputError struct {
	Message string
}

func (e *InvalidInputError) Error() string {
	if e == nil || e.Message == "" {
		return "invalid input"
	}

	return "invalid input: " + e.Message
}

// NewInvalidInputError creates a new InvalidInputError.
func NewInvalidInputError(message string) *InvalidInputError {
	return &InvalidInputError{Message: message}
}

// ConflictError represents business logic conflicts (duplicates, constraints, etc.).
type ConflictError struct {
	Reason string
}

func (e *ConflictError) Error() string {
	if e == nil || e.Reason == "" {
		return "link conflict"
	}

	return "link conflict: " + e.Reason
}

// NewConflictError creates a new ConflictError.
func NewConflictError(reason string) *ConflictError {
	return &ConflictError{Reason: reason}
}

// InternalError wraps unexpected domain or infrastructure errors.
type InternalError struct {
	Message string
	Err     error
}

func (e *InternalError) Error() string {
	switch {
	case e == nil:
		return "internal error"
	case e.Message != "":
		return "internal error: " + e.Message
	case e.Err != nil:
		return "internal error: " + e.Err.Error()
	default:
		return "internal error"
	}
}

// Unwrap returns the underlying error for error unwrapping support.
func (e *InternalError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Err
}

// NewInternalError creates a new InternalError.
func NewInternalError(message string) *InternalError {
	return &InternalError{Message: message}
}

// NewInternalErrorWithErr creates a new InternalError with wrapped error.
func NewInternalErrorWithErr(err error) *InternalError {
	return &InternalError{Err: err}
}

// PermissionDeniedError indicates the caller lacks required permissions.
type PermissionDeniedError struct {
	Err error
}

func (e *PermissionDeniedError) Error() string {
	return "permission denied"
}

// Unwrap returns the underlying error for error unwrapping support.
func (e *PermissionDeniedError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Err
}

// NewPermissionDeniedError creates a new PermissionDeniedError.
func NewPermissionDeniedError(err error) *PermissionDeniedError {
	return &PermissionDeniedError{Err: err}
}

