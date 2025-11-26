package errors

import (
	"errors"
	"fmt"
)

type Code string

const (
	CodeInvalidURL            Code = "METADATA_INVALID_URL"
	CodeMetadataNotFound      Code = "METADATA_NOT_FOUND"
	CodeMetadataExtraction    Code = "METADATA_EXTRACTION_FAILED"
	CodeScreenshotUnavailable Code = "METADATA_SCREENSHOT_UNAVAILABLE"
	CodeProcessingFailed      Code = "METADATA_PROCESSING_FAILED"
	CodeInternal              Code = "METADATA_INTERNAL"
)

var (
	ErrInvalidURL            = &Error{code: CodeInvalidURL, message: "invalid metadata URL"}
	ErrMetadataNotFound      = &Error{code: CodeMetadataNotFound, message: "metadata not found"}
	ErrMetadataExtraction    = &Error{code: CodeMetadataExtraction, message: "metadata extraction failed"}
	ErrScreenshotUnavailable = &Error{code: CodeScreenshotUnavailable, message: "screenshot unavailable"}
	ErrProcessingFailed      = &Error{code: CodeProcessingFailed, message: "metadata processing failed"}
	ErrInternal              = &Error{code: CodeInternal, message: "internal metadata error"}
)

type Error struct {
	code    Code
	message string
	cause   error
}

func newError(code Code, message string, cause error) *Error {
	return &Error{
		code:    code,
		message: message,
		cause:   cause,
	}
}

func NewInvalidURLError(url string, cause error) *Error {
	return newError(CodeInvalidURL, "invalid metadata URL: "+url, cause)
}

func NewMetadataNotFoundError(id string, cause error) *Error {
	return newError(CodeMetadataNotFound, "metadata not found: "+id, cause)
}

func NewMetadataExtractionError(target string, cause error) *Error {
	return newError(CodeMetadataExtraction, "failed to extract metadata for "+target, cause)
}

func NewScreenshotUnavailableError(target string, cause error) *Error {
	return newError(CodeScreenshotUnavailable, "screenshot unavailable for "+target, cause)
}

func NewInternalError(message string, cause error) *Error {
	return newError(CodeInternal, message, cause)
}

func newProcessingError(step string, cause error) *Error {
	return newError(CodeProcessingFailed, "processing failed at step "+step, cause)
}

func ProcessingFailed(step string, cause error) *Error {
	return newProcessingError(step, cause)
}

func (e *Error) Error() string {
	if e.message != "" {
		return e.message
	}

	return fmt.Sprintf("metadata error: %s", e.code)
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) Unwrap() error {
	return e.cause
}

// Is compares domain errors by code. Different instances with the same code
// are considered equal when using errors.Is.
func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return e.code == t.code
}

// Normalize converts arbitrary errors to domain errors, preserving existing domain errors.
func Normalize(step string, err error) *Error {
	if err == nil {
		return nil
	}

	var derr *Error
	if errors.As(err, &derr) {
		return derr
	}

	return ProcessingFailed(step, err)
}
