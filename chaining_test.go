package xerrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsValidationWithCode(t *testing.T) {
	err := &AppError{}
	result := err.AsValidationWithCode("VALIDATION_CODE")

	assert.Equal(t, ErrorTypeValidation, result.Type)
	assert.Equal(t, "VALIDATION_CODE", result.Code)
}

func TestAsValidationError(t *testing.T) {
	err := &AppError{}
	result := err.AsValidationError()

	assert.Equal(t, ErrorTypeValidation, result.Type)
	assert.Equal(t, CodeValidationError, result.Code)
}

func TestAsInvalidInput(t *testing.T) {
	err := &AppError{}
	result := err.AsInvalidInput()

	assert.Equal(t, ErrorTypeValidation, result.Type)
	assert.Equal(t, CodeInvalidInput, result.Code)
}

func TestAsRequiredField(t *testing.T) {
	err := &AppError{}
	result := err.AsRequiredField()

	assert.Equal(t, ErrorTypeValidation, result.Type)
	assert.Equal(t, CodeRequiredField, result.Code)
}

func TestAsInvalidFormat(t *testing.T) {
	err := &AppError{}
	result := err.AsInvalidFormat()

	assert.Equal(t, ErrorTypeValidation, result.Type)
	assert.Equal(t, CodeInvalidFormat, result.Code)
}

func TestAsInvalidRange(t *testing.T) {
	err := &AppError{}
	result := err.AsInvalidRange()

	assert.Equal(t, ErrorTypeValidation, result.Type)
	assert.Equal(t, CodeInvalidRange, result.Code)
}

func TestAsAuthentication(t *testing.T) {
	err := &AppError{}
	result := err.AsAuthentication()

	assert.Equal(t, ErrorTypeAuthentication, result.Type)
}

func TestAsAuthenticationWithCode(t *testing.T) {
	err := &AppError{}
	result := err.AsAuthenticationWithCode("AUTH_CODE")

	assert.Equal(t, ErrorTypeAuthentication, result.Type)
	assert.Equal(t, "AUTH_CODE", result.Code)
}

func TestAsInvalidCredentials(t *testing.T) {
	err := &AppError{}
	result := err.AsInvalidCredentials()

	assert.Equal(t, ErrorTypeAuthentication, result.Type)
	assert.Equal(t, CodeInvalidCredentials, result.Code)
}

func TestAsTokenExpired(t *testing.T) {
	err := &AppError{}
	result := err.AsTokenExpired()

	assert.Equal(t, ErrorTypeAuthentication, result.Type)
	assert.Equal(t, CodeTokenExpired, result.Code)
}

func TestAsTokenInvalid(t *testing.T) {
	err := &AppError{}
	result := err.AsTokenInvalid()

	assert.Equal(t, ErrorTypeAuthentication, result.Type)
	assert.Equal(t, CodeTokenInvalid, result.Code)
}

func TestAsLoginRequired(t *testing.T) {
	err := &AppError{}
	result := err.AsLoginRequired()

	assert.Equal(t, ErrorTypeAuthentication, result.Type)
	assert.Equal(t, CodeLoginRequired, result.Code)
}

func TestAsAuthorizationWithCode(t *testing.T) {
	err := &AppError{}
	result := err.AsAuthorizationWithCode("AUTHZ_CODE")

	assert.Equal(t, ErrorTypeAuthorization, result.Type)
	assert.Equal(t, "AUTHZ_CODE", result.Code)
}

func TestAsAccessDenied(t *testing.T) {
	err := &AppError{}
	result := err.AsAccessDenied()

	assert.Equal(t, ErrorTypeAuthorization, result.Type)
	assert.Equal(t, CodeAccessDenied, result.Code)
}

func TestAsInsufficientPermissions(t *testing.T) {
	err := &AppError{}
	result := err.AsInsufficientPermissions()

	assert.Equal(t, ErrorTypeAuthorization, result.Type)
	assert.Equal(t, CodeInsufficientPermissions, result.Code)
}

func TestAsResourceForbidden(t *testing.T) {
	err := &AppError{}
	result := err.AsResourceForbidden()

	assert.Equal(t, ErrorTypeAuthorization, result.Type)
	assert.Equal(t, CodeResourceForbidden, result.Code)
}

func TestAsNotFoundWithCode(t *testing.T) {
	err := &AppError{}
	result := err.AsNotFoundWithCode("NOT_FOUND_CODE")

	assert.Equal(t, ErrorTypeNotFound, result.Type)
	assert.Equal(t, "NOT_FOUND_CODE", result.Code)
}

func TestAsResourceNotFound(t *testing.T) {
	err := &AppError{}
	result := err.AsResourceNotFound()

	assert.Equal(t, ErrorTypeNotFound, result.Type)
	assert.Equal(t, CodeResourceNotFound, result.Code)
}

func TestAsConflictWithCode(t *testing.T) {
	err := &AppError{}
	result := err.AsConflictWithCode("CONFLICT_CODE")

	assert.Equal(t, ErrorTypeConflict, result.Type)
	assert.Equal(t, "CONFLICT_CODE", result.Code)
}

func TestAsResourceExists(t *testing.T) {
	err := &AppError{}
	result := err.AsResourceExists()

	assert.Equal(t, ErrorTypeConflict, result.Type)
	assert.Equal(t, CodeResourceExists, result.Code)
}

func TestAsInternalWithCode(t *testing.T) {
	err := &AppError{}
	result := err.AsInternalWithCode("INTERNAL_CODE")

	assert.Equal(t, ErrorTypeInternal, result.Type)
	assert.Equal(t, "INTERNAL_CODE", result.Code)
}

func TestAsDatabaseConnection(t *testing.T) {
	err := &AppError{}
	result := err.AsDatabaseConnection()

	assert.Equal(t, ErrorTypeInternal, result.Type)
	assert.Equal(t, CodeDatabaseConnection, result.Code)
}

func TestAsDatabaseError(t *testing.T) {
	err := &AppError{}
	result := err.AsDatabaseError()

	assert.Equal(t, ErrorTypeInternal, result.Type)
	assert.Equal(t, CodeDatabaseError, result.Code)
}

func TestAsDatabaseTimeout(t *testing.T) {
	err := &AppError{}
	result := err.AsDatabaseTimeout()

	assert.Equal(t, ErrorTypeInternal, result.Type)
	assert.Equal(t, CodeInternalTimeout, result.Code)
}

func TestAsDatabaseConstraint(t *testing.T) {
	err := &AppError{}
	result := err.AsDatabaseConstraint()

	assert.Equal(t, ErrorTypeInternal, result.Type)
	assert.Equal(t, CodeDatabaseConstraint, result.Code)
}

func TestAsExternalWithCode(t *testing.T) {
	err := &AppError{}
	result := err.AsExternalWithCode("EXTERNAL_CODE")

	assert.Equal(t, ErrorTypeExternal, result.Type)
	assert.Equal(t, "EXTERNAL_CODE", result.Code)
}

func TestAsServiceTimeout(t *testing.T) {
	err := &AppError{}
	result := err.AsServiceTimeout()

	assert.Equal(t, ErrorTypeExternal, result.Type)
	assert.Equal(t, CodeExternalTimeout, result.Code)
}

func TestAsServiceUnavailable(t *testing.T) {
	err := &AppError{}
	result := err.AsServiceUnavailable()

	assert.Equal(t, ErrorTypeUnavailable, result.Type)
	assert.Equal(t, CodeServiceUnavailable, result.Code)
}

func TestAsRateLimitWithCode(t *testing.T) {
	err := &AppError{}
	result := err.AsRateLimitWithCode("RATE_LIMIT_CODE")

	assert.Equal(t, ErrorTypeRateLimit, result.Type)
	assert.Equal(t, "RATE_LIMIT_CODE", result.Code)
}

func TestAsTooManyRequests(t *testing.T) {
	err := &AppError{}
	result := err.AsTooManyRequests()

	assert.Equal(t, ErrorTypeRateLimit, result.Type)
	assert.Equal(t, CodeRateLimitExceeded, result.Code)
}

func TestAsConfiguration(t *testing.T) {
	err := &AppError{}
	result := err.AsConfiguration()

	assert.Equal(t, ErrorTypeInternal, result.Type)
	assert.Equal(t, CodeConfigurationError, result.Code)
}

func TestAsTimeout(t *testing.T) {
	err := &AppError{}
	result := err.AsTimeout()

	assert.Equal(t, ErrorTypeInternal, result.Type)
	assert.Equal(t, CodeInternalTimeout, result.Code)
}

func TestSetTypeAndStatus_NilError(t *testing.T) {
	var err *AppError
	result := err.setTypeAndStatus(ErrorTypeValidation)

	assert.Nil(t, result)
}
