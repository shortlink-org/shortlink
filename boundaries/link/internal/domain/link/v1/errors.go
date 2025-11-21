package v1

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

// DomainErrorCode represents a domain error code
type DomainErrorCode string

const (
	// CodeNotFound - resource not found
	CodeNotFound DomainErrorCode = "not_found"
	// CodeInvalidInput - invalid input/validation error
	CodeInvalidInput DomainErrorCode = "invalid_input"
	// CodePermissionDenied - permission denied
	CodePermissionDenied DomainErrorCode = "permission_denied"
	// CodeConflict - business logic conflict (e.g., duplicate, constraint violation)
	CodeConflict DomainErrorCode = "conflict"
	// CodeInternal - internal domain error
	CodeInternal DomainErrorCode = "internal"
)

// DomainError is the interface that all domain errors must implement
// Domain owns the semantics of error mapping to transport codes
type DomainError interface {
	error
	Code() DomainErrorCode
	GRPCCode() codes.Code
}

// baseError provides common error functionality
// All domain errors should embed this anonymously
type baseError struct {
	code DomainErrorCode
	grpc codes.Code
}

func (e *baseError) Code() DomainErrorCode {
	return e.code
}

func (e *baseError) GRPCCode() codes.Code {
	return e.grpc
}

// NotFoundError - not found link
type NotFoundError struct {
	baseError
	Hash string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("link not found: hash=%s", e.Hash)
}

// NewNotFoundError creates a new NotFoundError
func NewNotFoundError(hash string) *NotFoundError {
	return &NotFoundError{
		baseError: baseError{
			code: CodeNotFound,
			grpc: codes.NotFound,
		},
		Hash: hash,
	}
}

// NotFoundByHashError - not found link by hash
type NotFoundByHashError struct {
	baseError
	Hash string
}

func (e *NotFoundByHashError) Error() string {
	return fmt.Sprintf("link not found: hash=%s", e.Hash)
}

// NewNotFoundByHashError creates a new NotFoundByHashError
func NewNotFoundByHashError(hash string) *NotFoundByHashError {
	return &NotFoundByHashError{
		baseError: baseError{
			code: CodeNotFound,
			grpc: codes.NotFound,
		},
		Hash: hash,
	}
}

// InvalidInputError - invalid input/validation error
type InvalidInputError struct {
	baseError
	Message string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.Message)
}

// NewInvalidInputError creates a new InvalidInputError
func NewInvalidInputError(message string) *InvalidInputError {
	return &InvalidInputError{
		baseError: baseError{
			code: CodeInvalidInput,
			grpc: codes.InvalidArgument,
		},
		Message: message,
	}
}

// ConflictError - business logic conflict (e.g., duplicate, constraint violation)
type ConflictError struct {
	baseError
	Reason string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("link conflict: %s", e.Reason)
}

// NewConflictError creates a new ConflictError
func NewConflictError(reason string) *ConflictError {
	return &ConflictError{
		baseError: baseError{
			code: CodeConflict,
			grpc: codes.FailedPrecondition,
		},
		Reason: reason,
	}
}

// InternalError - internal domain/infrastructure error
type InternalError struct {
	baseError
	Message string
	Err     error
}

func (e *InternalError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("internal error: %s", e.Message)
	}
	if e.Err != nil {
		return fmt.Sprintf("internal error: %s", e.Err.Error())
	}
	return "internal error"
}

// NewInternalError creates a new InternalError
func NewInternalError(message string) *InternalError {
	return &InternalError{
		baseError: baseError{
			code: CodeInternal,
			grpc: codes.Internal,
		},
		Message: message,
	}
}

// NewInternalErrorWithErr creates a new InternalError with wrapped error
func NewInternalErrorWithErr(err error) *InternalError {
	return &InternalError{
		baseError: baseError{
			code: CodeInternal,
			grpc: codes.Internal,
		},
		Err: err,
	}
}

// Unwrap returns the underlying error for error unwrapping support
func (e *InternalError) Unwrap() error {
	return e.Err
}

// PermissionDeniedError - permission denied
type PermissionDeniedError struct {
	baseError
	Err error
}

func (e *PermissionDeniedError) Error() string {
	return "permission denied"
}

// NewPermissionDeniedError creates a new PermissionDeniedError
func NewPermissionDeniedError(err error) *PermissionDeniedError {
	return &PermissionDeniedError{
		baseError: baseError{
			code: CodePermissionDenied,
			grpc: codes.PermissionDenied,
		},
		Err: err,
	}
}

// Unwrap returns the underlying error for error unwrapping support
// This allows errors.Is() and errors.As() to work correctly with wrapped errors
// and enables OTEL to log the original error
func (e *PermissionDeniedError) Unwrap() error {
	return e.Err
}
