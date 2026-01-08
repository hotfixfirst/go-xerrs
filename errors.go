package xerrs

import (
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
)

// AppError represents a structured application error with HTTP mapping capabilities
// and enhanced error details using cockroachdb/errors.
type AppError struct {
	Type       ErrorType `json:"type"`
	Code       string    `json:"code"`
	Message    string    `json:"message"`
	Details    string    `json:"details,omitempty"`
	HTTPStatus int       `json:"http_status,omitempty"`
	cause      error     `json:"-"`
}

// NewAppError creates a new AppError with specified type, code, and message.
func NewAppError(errorType ErrorType, code, message string) *AppError {
	code = strings.TrimSpace(code)
	message = strings.TrimSpace(message)
	if errorType == "" {
		errorType = ErrorTypeInternal
	}
	if code == "" {
		code = CodeInternalError
	}
	if message == "" {
		message = MsgUnknownError
	}
	return &AppError{
		Type:       errorType,
		Code:       code,
		Message:    message,
		HTTPStatus: errorType.DefaultHTTPStatus(),
		cause:      errors.NewWithDepth(1, message),
	}
}

// New creates a new AppError with a default internal error type and message.
func New(message string) *AppError {
	message = strings.TrimSpace(message)
	if message == "" {
		message = MsgUnknownError
	}
	return &AppError{
		Type:       ErrorTypeInternal,
		Code:       CodeInternalError,
		Message:    message,
		HTTPStatus: StatusInternalServerError,
		cause:      errors.NewWithDepth(1, message),
	}
}

// Wrap wraps an existing error into an AppError with a specified message.
func Wrap(err error, message string) *AppError {
	message = strings.TrimSpace(message)
	if message == "" {
		message = MsgUnknownError
	}
	if err == nil {
		return &AppError{
			Type:       ErrorTypeInternal,
			Code:       CodeInternalError,
			Message:    message,
			HTTPStatus: StatusInternalServerError,
			cause:      errors.NewWithDepth(1, message),
		}
	}
	// Check if it's already an AppError - preserve original structure
	if appErr, ok := AsAppError(err); ok {
		return &AppError{
			Type:       appErr.Type,
			Code:       appErr.Code,
			Message:    message,
			Details:    appErr.Details,
			HTTPStatus: appErr.HTTPStatus,
			cause:      errors.WrapWithDepth(1, appErr.cause, message),
		}
	}
	// Auto-detect error type and code from the original error
	errorType, code := detectErrorTypeAndCode(err)
	return &AppError{
		Type:       errorType,
		Code:       code,
		Message:    message,
		HTTPStatus: errorType.DefaultHTTPStatus(),
		cause:      errors.WrapWithDepth(1, err, message),
	}
}

// AsAppError safely converts an error to AppError if possible.
func AsAppError(err error) (*AppError, bool) {
	if err == nil {
		return nil, false
	}
	// Use errors.As for safer type assertion
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	// Don't auto-wrap, just return nil to avoid infinite recursion
	return nil, false
}

// IsAppError checks if an error is an AppError.
func IsAppError(err error) bool {
	_, ok := AsAppError(err)
	return ok
}

// Error implements the error interface.
func (e *AppError) Error() string {
	if details := strings.TrimSpace(e.Details); details != "" {
		return fmt.Sprintf("[%s] %s: %s - %s", e.Type, e.Code, e.Message, details)
	}
	return fmt.Sprintf("[%s] %s: %s", e.Type, e.Code, e.Message)
}

// String returns a human-readable representation of the error.
func (e *AppError) String() string {
	return e.Error()
}

// Unwrap retrieves the immediate underlying cause of the error.
func (e *AppError) Unwrap() error {
	return errors.Unwrap(e.cause)
}

// UnwrapAll retrieves the root cause of the error by traversing the entire chain of causes.
func (e *AppError) UnwrapAll() error {
	if e == nil {
		return nil
	}
	return errors.Cause(e.cause)
}

// Cause returns the direct cause of the error, if any.
func (e *AppError) Cause() error {
	if e == nil || e.cause == nil {
		return nil
	}
	return e.cause
}

// WithType sets the error type for the AppError.
func (e *AppError) WithType(errorType ErrorType) *AppError {
	if e == nil {
		return nil
	}
	e.Type = errorType
	return e
}

// WithCode sets the error code for the AppError.
func (e *AppError) WithCode(code string) *AppError {
	if e == nil {
		return nil
	}
	if code := strings.TrimSpace(code); code != "" {
		e.Code = code
	}
	return e
}

// WithMessage sets the error message for the AppError.
func (e *AppError) WithMessage(message string) *AppError {
	if e == nil {
		return nil
	}
	if message := strings.TrimSpace(message); message != "" {
		e.Message = message
	}
	return e
}

// WithDetails adds detailed information to the error.
func (e *AppError) WithDetails(details string) *AppError {
	if e == nil {
		return nil
	}
	e.Details = strings.TrimSpace(details)
	return e
}

// WithHTTPStatus allows overriding the default HTTP status code.
func (e *AppError) WithHTTPStatus(status int) *AppError {
	if e == nil {
		return nil
	}
	if status >= 100 && status <= 599 {
		e.HTTPStatus = status
	}
	return e
}

// WithCause sets the underlying cause of the error.
func (e *AppError) WithCause(cause error) *AppError {
	if e == nil {
		return nil
	}
	if cause != nil {
		e.cause = errors.WrapWithDepth(1, cause, e.Message)
	}
	return e
}

// WithCodeAndMessage sets both the error code and message for the AppError.
func (e *AppError) WithCodeAndMessage(code, message string) *AppError {
	if e == nil {
		return nil
	}
	if code := strings.TrimSpace(code); code != "" {
		e.Code = code
	}
	if message := strings.TrimSpace(message); message != "" {
		e.Message = message
	}
	return e
}

// GetHTTPStatus returns the HTTP status code for this error.
func (e *AppError) GetHTTPStatus() int {
	if e == nil {
		return StatusInternalServerError
	}
	if e.HTTPStatus != 0 {
		return e.HTTPStatus
	}
	return e.Type.DefaultHTTPStatus()
}

// IsType checks if the error is of a specific type.
func (e *AppError) IsType(errorType ErrorType) bool {
	if e == nil {
		return false
	}
	return e.Type == errorType
}

// HasCode checks if the error has a specific code.
func (e *AppError) HasCode(code string) bool {
	if e == nil {
		return false
	}
	return e.Code == code
}

// GetStackTrace returns the full stack trace if available.
func (e *AppError) GetStackTrace() string {
	if e == nil || e.cause == nil {
		return "no stack trace available"
	}
	return fmt.Sprintf("%+v", e.cause)
}

// GetStackTraceLines returns stack trace as a slice of strings from an error.
// Removes leading whitespace, tabs, and carriage returns from each line.
func (e *AppError) GetStackTraceLines() []string {
	if e == nil || e.cause == nil {
		return []string{"no stack trace available"}
	}
	stack := fmt.Sprintf("%+v", e.cause)
	lines := strings.Split(stack, "\n")

	// Clean each line by removing leading/trailing whitespace and special characters
	cleanedLines := make([]string, 0, len(lines))
	for _, line := range lines {
		// Remove \t, \r, and leading/trailing spaces
		cleaned := strings.TrimSpace(line)
		cleaned = strings.ReplaceAll(cleaned, "\t", "")
		cleaned = strings.ReplaceAll(cleaned, "\r", "")

		// Only add non-empty lines
		if cleaned != "" {
			cleanedLines = append(cleanedLines, cleaned)
		}
	}

	return cleanedLines
}
