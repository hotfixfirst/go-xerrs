package xerrs

import (
	"context"
	"database/sql"
	"strings"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

// detectErrorTypeAndCode analyzes an error and returns appropriate ErrorType and Code.
func detectErrorTypeAndCode(err error) (ErrorType, string) {
	// Check if it's already an AppError
	if appErr, ok := AsAppError(err); ok {
		return appErr.Type, appErr.Code
	}

	// Fast path: Check specific error types first (no string operations)
	switch {
	// Context-related errors
	case errors.Is(err, context.DeadlineExceeded):
		return ErrorTypeInternal, CodeInternalTimeout
	case errors.Is(err, context.Canceled):
		return ErrorTypeInternal, CodeOperationCanceled

	// Database-related errors (GORM)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrorTypeNotFound, CodeResourceNotFound
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return ErrorTypeInternal, CodeDatabaseError
	case errors.Is(err, gorm.ErrNotImplemented):
		return ErrorTypeInternal, CodeDatabaseError
	case errors.Is(err, gorm.ErrMissingWhereClause):
		return ErrorTypeValidation, CodeInvalidInput
	case errors.Is(err, gorm.ErrUnsupportedRelation):
		return ErrorTypeInternal, CodeDatabaseError
	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return ErrorTypeValidation, CodeRequiredField

	// SQL-related errors
	case errors.Is(err, sql.ErrNoRows):
		return ErrorTypeNotFound, CodeResourceNotFound
	case errors.Is(err, sql.ErrTxDone):
		return ErrorTypeInternal, CodeDatabaseError
	case errors.Is(err, sql.ErrConnDone):
		return ErrorTypeInternal, CodeDatabaseConnection
	}

	// Slow path: Pattern matching (only when necessary)
	return detectFromErrorMessage(err.Error())
}

// Pre-compiled pattern groups for better performance
type errorPattern struct {
	patterns  []string
	errorType ErrorType
	code      string
}

// errorPatterns is pre-compiled for better performance
var errorPatterns = []errorPattern{
	// Validation errors (highest priority - check JSON/parsing first)
	{
		patterns:  []string{"json", "unmarshal", "parse", "invalid character", "looking for beginning", "unexpected end of JSON input"},
		errorType: ErrorTypeValidation,
		code:      CodeInvalidFormat,
	},
	{
		patterns:  []string{"validation failed", "invalid format", "malformed"},
		errorType: ErrorTypeValidation,
		code:      CodeInvalidFormat,
	},
	{
		patterns:  []string{"required", "missing"},
		errorType: ErrorTypeValidation,
		code:      CodeRequiredField,
	},
	{
		patterns:  []string{"out of range", "too large", "too small"},
		errorType: ErrorTypeValidation,
		code:      CodeInvalidRange,
	},

	// Database constraint errors (second priority - most common)
	{
		patterns:  []string{"duplicate key", "unique constraint", "already exists"},
		errorType: ErrorTypeConflict,
		code:      CodeResourceExists,
	},
	{
		patterns:  []string{"foreign key constraint", "violates foreign key"},
		errorType: ErrorTypeValidation,
		code:      CodeInvalidInput,
	},
	{
		patterns:  []string{"not null constraint", "violates not-null"},
		errorType: ErrorTypeValidation,
		code:      CodeRequiredField,
	},
	{
		patterns:  []string{"check constraint"},
		errorType: ErrorTypeValidation,
		code:      CodeInvalidRange,
	},

	// Authentication errors (third priority)
	{
		patterns:  []string{"unauthorized", "invalid credentials", "authentication failed"},
		errorType: ErrorTypeAuthentication,
		code:      CodeInvalidCredentials,
	},
	{
		patterns:  []string{"token expired", "jwt expired"},
		errorType: ErrorTypeAuthentication,
		code:      CodeTokenExpired,
	},
	{
		patterns:  []string{"invalid token", "malformed token"},
		errorType: ErrorTypeAuthentication,
		code:      CodeTokenInvalid,
	},

	// Authorization errors
	{
		patterns:  []string{"forbidden", "access denied", "permission denied"},
		errorType: ErrorTypeAuthorization,
		code:      CodeAccessDenied,
	},

	// Rate limiting
	{
		patterns:  []string{"rate limit", "too many requests", "quota exceeded"},
		errorType: ErrorTypeRateLimit,
		code:      CodeRateLimitExceeded,
	},

	// Network errors (lower priority - more generic)
	{
		patterns:  []string{"connection refused", "connection reset", "no such host", "network is unreachable"},
		errorType: ErrorTypeExternal,
		code:      CodeExternalError,
	},
	{
		patterns:  []string{"timeout", "deadline exceeded"},
		errorType: ErrorTypeExternal,
		code:      CodeExternalTimeout,
	},
	{
		patterns:  []string{"service unavailable", "bad gateway", "gateway timeout"},
		errorType: ErrorTypeUnavailable,
		code:      CodeExternalUnavailable,
	},

	// File/IO errors
	{
		patterns:  []string{"file not found", "no such file"},
		errorType: ErrorTypeNotFound,
		code:      CodeResourceNotFound,
	},

	// Configuration errors (lowest priority)
	{
		patterns:  []string{"configuration", "config", "environment"},
		errorType: ErrorTypeInternal,
		code:      CodeConfigurationError,
	},
}

// detectFromErrorMessage performs optimized pattern matching on error message.
func detectFromErrorMessage(errMsg string) (ErrorType, string) {
	// Convert to lowercase once
	lowerMsg := strings.ToLower(errMsg)

	// Check each pattern group
	for _, pattern := range errorPatterns {
		if matchesAnyPattern(lowerMsg, pattern.patterns) {
			return pattern.errorType, pattern.code
		}
	}

	// Default to internal error
	return ErrorTypeInternal, CodeInternalError
}

// IsRecordNotFound checks if an error indicates that a record was not found.
func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, sql.ErrNoRows)
}

// matchesAnyPattern checks if message contains any of the patterns (optimized).
func matchesAnyPattern(msg string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(msg, pattern) {
			return true
		}
	}
	return false
}
