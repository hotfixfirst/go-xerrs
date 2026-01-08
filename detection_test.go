package xerrs

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDetectErrorTypeAndCode(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedType ErrorType
		expectedCode string
	}{
		{"AppError", &AppError{Type: ErrorTypeValidation, Code: CodeInvalidInput}, ErrorTypeValidation, CodeInvalidInput},
		{"Context Deadline Exceeded", context.DeadlineExceeded, ErrorTypeInternal, CodeInternalTimeout},
		{"Context Canceled", context.Canceled, ErrorTypeInternal, CodeOperationCanceled},
		{"GORM Record Not Found", gorm.ErrRecordNotFound, ErrorTypeNotFound, CodeResourceNotFound},
		{"GORM Invalid Transaction", gorm.ErrInvalidTransaction, ErrorTypeInternal, CodeDatabaseError},
		{"GORM Missing Where Clause", gorm.ErrMissingWhereClause, ErrorTypeValidation, CodeInvalidInput},
		{"SQL No Rows", sql.ErrNoRows, ErrorTypeNotFound, CodeResourceNotFound},
		{"SQL Tx Done", sql.ErrTxDone, ErrorTypeInternal, CodeDatabaseError},
		{"Default Case", errors.New("unknown error"), ErrorTypeInternal, CodeInternalError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeResult, codeResult := detectErrorTypeAndCode(tt.err)
			assert.Equal(t, tt.expectedType, typeResult)
			assert.Equal(t, tt.expectedCode, codeResult)
		})
	}
}

func TestDetectErrorTypeAndCode_GORMAndSQLCases(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedType ErrorType
		expectedCode string
	}{
		{"GORM Unsupported Relation", gorm.ErrUnsupportedRelation, ErrorTypeInternal, CodeDatabaseError},
		{"GORM Primary Key Required", gorm.ErrPrimaryKeyRequired, ErrorTypeValidation, CodeRequiredField},
		{"GORM Not Implemented", gorm.ErrNotImplemented, ErrorTypeInternal, CodeDatabaseError},
		{"SQL Connection Done", sql.ErrConnDone, ErrorTypeInternal, CodeDatabaseConnection},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeResult, codeResult := detectErrorTypeAndCode(tt.err)
			assert.Equal(t, tt.expectedType, typeResult)
			assert.Equal(t, tt.expectedCode, codeResult)
		})
	}
}

func TestDetectFromErrorMessage(t *testing.T) {
	tests := []struct {
		name         string
		errMsg       string
		expectedType ErrorType
		expectedCode string
	}{
		{"Validation Error - JSON", "invalid character in JSON", ErrorTypeValidation, CodeInvalidFormat},
		{"Database Constraint - Duplicate Key", "duplicate key value violates unique constraint", ErrorTypeConflict, CodeResourceExists},
		{"Authentication Error - Token Expired", "token expired", ErrorTypeAuthentication, CodeTokenExpired},
		{"Authorization Error - Access Denied", "access denied", ErrorTypeAuthorization, CodeAccessDenied},
		{"Rate Limit Error", "rate limit exceeded", ErrorTypeRateLimit, CodeRateLimitExceeded},
		{"Default Case", "unknown error", ErrorTypeInternal, CodeInternalError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeResult, codeResult := detectFromErrorMessage(tt.errMsg)
			assert.Equal(t, tt.expectedType, typeResult)
			assert.Equal(t, tt.expectedCode, codeResult)
		})
	}
}

func TestMatchesAnyPattern(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		patterns []string
		expected bool
	}{
		{"Match Found", "rate limit exceeded", []string{"rate limit", "quota exceeded"}, true},
		{"No Match Found", "unknown error", []string{"rate limit", "quota exceeded"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchesAnyPattern(tt.msg, tt.patterns)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsRecordNotFound(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "GORM ErrRecordNotFound",
			err:      gorm.ErrRecordNotFound,
			expected: true,
		},
		{
			name:     "SQL ErrNoRows",
			err:      sql.ErrNoRows,
			expected: true,
		},
		{
			name:     "wrapped GORM ErrRecordNotFound",
			err:      errors.Join(gorm.ErrRecordNotFound, errors.New("additional context")),
			expected: true,
		},
		{
			name:     "wrapped SQL ErrNoRows",
			err:      errors.Join(sql.ErrNoRows, errors.New("additional context")),
			expected: true,
		},
		{
			name:     "other GORM error",
			err:      gorm.ErrInvalidTransaction,
			expected: false,
		},
		{
			name:     "other SQL error",
			err:      sql.ErrTxDone,
			expected: false,
		},
		{
			name:     "generic error",
			err:      errors.New("some generic error"),
			expected: false,
		},
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "context error",
			err:      context.DeadlineExceeded,
			expected: false,
		},
		{
			name:     "custom error message with 'not found'",
			err:      errors.New("record not found in database"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsRecordNotFound(tt.err)
			assert.Equal(t, tt.expected, result, "IsRecordNotFound(%v) should return %v", tt.err, tt.expected)
		})
	}
}
