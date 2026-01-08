package xerrs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppError(t *testing.T) {
	err := NewAppError(ErrorTypeValidation, "VALIDATION_ERROR", "Validation failed")
	assert.Equal(t, ErrorTypeValidation, err.Type)
	assert.Equal(t, "VALIDATION_ERROR", err.Code)
	assert.Equal(t, "Validation failed", err.Message)
	assert.Equal(t, ErrorTypeValidation.DefaultHTTPStatus(), err.HTTPStatus)
}

func TestNew(t *testing.T) {
	err := New("An internal error occurred")
	assert.Equal(t, ErrorTypeInternal, err.Type)
	assert.Equal(t, CodeInternalError, err.Code)
	assert.Equal(t, "An internal error occurred", err.Message)
	assert.Equal(t, StatusInternalServerError, err.HTTPStatus)
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	err := Wrap(originalErr, "wrapped error")
	assert.Equal(t, ErrorTypeInternal, err.Type)
	assert.Equal(t, CodeInternalError, err.Code)
	assert.Equal(t, "wrapped error", err.Message)
	assert.Equal(t, StatusInternalServerError, err.HTTPStatus)
	assert.Contains(t, errors.Unwrap(err).Error(), "original error")
}

func TestWithType(t *testing.T) {
	err := New("An error occurred").WithType(ErrorTypeValidation)
	assert.Equal(t, ErrorTypeValidation, err.Type)
}

func TestWithCode(t *testing.T) {
	err := New("An error occurred").WithCode("NEW_CODE")
	assert.Equal(t, "NEW_CODE", err.Code)
}

func TestWithMessage(t *testing.T) {
	err := New("An error occurred").WithMessage("Updated message")
	assert.Equal(t, "Updated message", err.Message)
}

func TestWithDetails(t *testing.T) {
	err := New("An error occurred").WithDetails("Additional details")
	assert.Equal(t, "Additional details", err.Details)
}

func TestWithHTTPStatus(t *testing.T) {
	err := New("An error occurred").WithHTTPStatus(418)
	assert.Equal(t, 418, err.HTTPStatus)
}

func TestWithCause(t *testing.T) {
	cause := errors.New("root cause")
	err := New("An error occurred").WithCause(cause)
	assert.Contains(t, errors.Unwrap(err).Error(), "root cause")
}

func TestWithCodeAndMessage(t *testing.T) {
	err := New("An error occurred").WithCodeAndMessage("NEW_CODE", "New message")
	assert.Equal(t, "NEW_CODE", err.Code)
	assert.Equal(t, "New message", err.Message)
}

func TestGetHTTPStatus(t *testing.T) {
	err := New("An error occurred")
	assert.Equal(t, StatusInternalServerError, err.GetHTTPStatus())

	err.WithHTTPStatus(418)
	assert.Equal(t, 418, err.GetHTTPStatus())
}

func TestIsType(t *testing.T) {
	err := New("An error occurred").WithType(ErrorTypeValidation)
	assert.True(t, err.IsType(ErrorTypeValidation))
	assert.False(t, err.IsType(ErrorTypeInternal))
}

func TestHasCode(t *testing.T) {
	err := New("An error occurred").WithCode("NEW_CODE")
	assert.True(t, err.HasCode("NEW_CODE"))
	assert.False(t, err.HasCode("OTHER_CODE"))
}

func TestGetStackTrace(t *testing.T) {
	err := New("An error occurred")
	stackTrace := err.GetStackTrace()
	assert.Contains(t, stackTrace, "An error occurred")
}

func TestGetStackTraceLines(t *testing.T) {
	err := New("An error occurred")
	stackLines := err.GetStackTraceLines()
	assert.Greater(t, len(stackLines), 0)
	assert.Contains(t, stackLines[0], "An error occurred")
}

func TestUnwrap(t *testing.T) {
	cause := errors.New("root cause")
	err := New("An error occurred").WithCause(cause)
	assert.Contains(t, err.Unwrap().Error(), "root cause")
}

func TestUnwrapAll(t *testing.T) {
	cause := errors.New("root cause")
	err := New("An error occurred").WithCause(cause)
	assert.Equal(t, cause, err.UnwrapAll())
}

func TestCause(t *testing.T) {
	cause := errors.New("root cause")
	err := New("An error occurred").WithCause(cause)
	assert.Contains(t, err.Unwrap().Error(), "root cause")
}
