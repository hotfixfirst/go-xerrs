package xerrs

import "net/http"

// Error codes for application error handling.
const (
	// Validation error codes (400)
	CodeValidationError = "VALIDATION_ERROR"
	CodeInvalidInput    = "INVALID_INPUT"
	CodeRequiredField   = "REQUIRED_FIELD"
	CodeInvalidFormat   = "INVALID_FORMAT"
	CodeInvalidRange    = "INVALID_RANGE"

	// Authentication error codes (401)
	CodeInvalidCredentials = "INVALID_CREDENTIALS"
	CodeTokenExpired       = "TOKEN_EXPIRED"
	CodeTokenInvalid       = "TOKEN_INVALID"
	CodeLoginRequired      = "LOGIN_REQUIRED"
	CodeAuthRequired       = "AUTH_REQUIRED"

	// Authorization error codes (403)
	CodeAccessDenied            = "ACCESS_DENIED"
	CodeInsufficientPermissions = "INSUFFICIENT_PERMISSIONS"
	CodeResourceForbidden       = "RESOURCE_FORBIDDEN"
	CodeInsufficientRole        = "INSUFFICIENT_ROLE"

	// Resource error codes (404, 409)
	CodeResourceNotFound = "RESOURCE_NOT_FOUND"
	CodeResourceExists   = "RESOURCE_EXISTS"

	// Rate limit error codes (429)
	CodeRateLimitExceeded = "RATE_LIMIT_EXCEEDED"

	// Internal system error codes (500)
	CodeInternalError      = "INTERNAL_ERROR"
	CodeDatabaseError      = "DATABASE_ERROR"
	CodeDatabaseConnection = "DATABASE_CONNECTION"
	CodeDatabaseConstraint = "DATABASE_CONSTRAINT"
	CodeInternalTimeout    = "INTERNAL_TIMEOUT"
	CodeConfigurationError = "CONFIGURATION_ERROR"
	CodeOperationCanceled  = "OPERATION_CANCELED"

	// Context and middleware error codes (500)
	CodeInvalidUserContext    = "INVALID_USER_CONTEXT"
	CodeOrgContextMissing     = "ORG_CONTEXT_MISSING"
	CodeInvalidOrgContext     = "INVALID_ORG_CONTEXT"
	CodeUserRoleNotFound      = "USER_ROLE_NOT_FOUND"
	CodePermissionCheckFailed = "PERMISSION_CHECK_FAILED"

	// External service error codes (502)
	CodeExternalError       = "EXTERNAL_ERROR"
	CodeExternalTimeout     = "EXTERNAL_TIMEOUT"
	CodeExternalUnavailable = "EXTERNAL_UNAVAILABLE"

	// Service unavailable error codes (503)
	CodeServiceUnavailable = "SERVICE_UNAVAILABLE"
)

// ErrorType defines the category of application errors for HTTP status mapping.
type ErrorType string

const (
	// Client-side Errors (4xx)
	ErrorTypeValidation     ErrorType = "VALIDATION"     // 400 Bad Request
	ErrorTypeAuthentication ErrorType = "AUTHENTICATION" // 401 Unauthorized
	ErrorTypeAuthorization  ErrorType = "AUTHORIZATION"  // 403 Forbidden
	ErrorTypeNotFound       ErrorType = "NOT_FOUND"      // 404 Not Found
	ErrorTypeConflict       ErrorType = "CONFLICT"       // 409 Conflict
	ErrorTypeRateLimit      ErrorType = "RATE_LIMIT"     // 429 Too Many Requests

	// Server-side Errors (5xx)
	ErrorTypeInternal    ErrorType = "INTERNAL"    // 500 Internal Server Error
	ErrorTypeExternal    ErrorType = "EXTERNAL"    // 502 Bad Gateway
	ErrorTypeUnavailable ErrorType = "UNAVAILABLE" // 503 Service Unavailable
)

// Default HTTP status codes for error types.
const (
	StatusInternalServerError = http.StatusInternalServerError
)

// DefaultHTTPStatus returns the default HTTP status code for each error type.
func (et ErrorType) DefaultHTTPStatus() int {
	switch et {
	// Client-side errors (4xx)
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeAuthentication:
		return http.StatusUnauthorized
	case ErrorTypeAuthorization:
		return http.StatusForbidden
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeConflict:
		return http.StatusConflict
	case ErrorTypeRateLimit:
		return http.StatusTooManyRequests

	// Server-side errors (5xx)
	case ErrorTypeInternal:
		return http.StatusInternalServerError
	case ErrorTypeExternal:
		return http.StatusBadGateway
	case ErrorTypeUnavailable:
		return http.StatusServiceUnavailable

	default:
		return http.StatusInternalServerError
	}
}

// Error messages for consistent error reporting
const (
	// General messages
	MsgUnknownError = "Unknown error occurred"

	// Authentication messages
	MsgAuthRequired       = "Authentication required"
	MsgInvalidCredentials = "Invalid credentials provided"

	// Authorization messages
	MsgInsufficientPermissions = "Access denied - insufficient permissions"
	MsgInsufficientRole        = "Access denied - insufficient role permissions"

	// Context error messages
	MsgInvalidUserContext    = "Invalid user context"
	MsgOrgContextMissing     = "Organization context missing"
	MsgInvalidOrgContext     = "Invalid organization context"
	MsgUserRoleNotFound      = "User role not found"
	MsgPermissionCheckFailed = "Permission check failed"
)
