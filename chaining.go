package xerrs

// AsValidationWithCode converts the error to a validation error with a specific code.
func (e *AppError) AsValidationWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeValidation).WithCode(code)
}

// AsValidationError converts the error to a validation error with a default code.
func (e *AppError) AsValidationError() *AppError {
	return e.AsValidationWithCode(CodeValidationError)
}

// AsInvalidInput converts the error to an invalid input validation error.
func (e *AppError) AsInvalidInput() *AppError {
	return e.AsValidationWithCode(CodeInvalidInput)
}

// AsRequiredField converts the error to a required field validation error.
func (e *AppError) AsRequiredField() *AppError {
	return e.AsValidationWithCode(CodeRequiredField)
}

// AsInvalidFormat converts the error to an invalid format validation error.
func (e *AppError) AsInvalidFormat() *AppError {
	return e.AsValidationWithCode(CodeInvalidFormat)
}

// AsInvalidRange converts the error to an invalid range validation error.
func (e *AppError) AsInvalidRange() *AppError {
	return e.AsValidationWithCode(CodeInvalidRange)
}

// AsAuthentication converts the error to an authentication error type.
func (e *AppError) AsAuthentication() *AppError {
	return e.setTypeAndStatus(ErrorTypeAuthentication)
}

// AsAuthenticationWithCode converts the error to an authentication error with a specific code.
func (e *AppError) AsAuthenticationWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeAuthentication).WithCode(code)
}

// AsInvalidCredentials converts the error to an invalid credentials error.
func (e *AppError) AsInvalidCredentials() *AppError {
	return e.AsAuthenticationWithCode(CodeInvalidCredentials)
}

// AsTokenExpired converts the error to a token expired error.
func (e *AppError) AsTokenExpired() *AppError {
	return e.AsAuthenticationWithCode(CodeTokenExpired)
}

// AsTokenInvalid converts the error to a token invalid error.
func (e *AppError) AsTokenInvalid() *AppError {
	return e.AsAuthenticationWithCode(CodeTokenInvalid)
}

// AsLoginRequired converts the error to a login required error.
func (e *AppError) AsLoginRequired() *AppError {
	return e.AsAuthenticationWithCode(CodeLoginRequired)
}

// AsAuthorizationWithCode converts the error to an authorization error with a specific code.
func (e *AppError) AsAuthorizationWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeAuthorization).WithCode(code)
}

// AsAccessDenied converts the error to an access denied error.
func (e *AppError) AsAccessDenied() *AppError {
	return e.AsAuthorizationWithCode(CodeAccessDenied)
}

// AsInsufficientPermissions converts the error to an insufficient permissions error.
func (e *AppError) AsInsufficientPermissions() *AppError {
	return e.AsAuthorizationWithCode(CodeInsufficientPermissions)
}

// AsResourceForbidden converts the error to a resource forbidden error.
func (e *AppError) AsResourceForbidden() *AppError {
	return e.AsAuthorizationWithCode(CodeResourceForbidden)
}

// AsNotFoundWithCode converts the error to a not found error with a specific code.
func (e *AppError) AsNotFoundWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeNotFound).WithCode(code)
}

// AsResourceNotFound converts the error to a resource not found error.
func (e *AppError) AsResourceNotFound() *AppError {
	return e.AsNotFoundWithCode(CodeResourceNotFound)
}

// AsConflictWithCode converts the error to a conflict error with a specific code.
func (e *AppError) AsConflictWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeConflict).WithCode(code)
}

// AsResourceExists converts the error to a resource exists error.
func (e *AppError) AsResourceExists() *AppError {
	return e.AsConflictWithCode(CodeResourceExists)
}

// AsInternalWithCode converts the error to an internal error with a specific code.
func (e *AppError) AsInternalWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeInternal).WithCode(code)
}

// AsDatabaseError converts the error to a database error.
func (e *AppError) AsDatabaseError() *AppError {
	return e.AsInternalWithCode(CodeDatabaseError)
}

// AsDatabaseConnection converts the error to a database connection error.
func (e *AppError) AsDatabaseConnection() *AppError {
	return e.AsInternalWithCode(CodeDatabaseConnection)
}

// AsDatabaseTimeout converts the error to a database timeout error.
func (e *AppError) AsDatabaseTimeout() *AppError {
	return e.AsInternalWithCode(CodeInternalTimeout)
}

// AsDatabaseConstraint converts the error to a database constraint error.
func (e *AppError) AsDatabaseConstraint() *AppError {
	return e.AsInternalWithCode(CodeDatabaseConstraint)
}

// AsUnavailableWithCode converts the error to an unavailable service error with a specific code.
func (e *AppError) AsUnavailableWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeUnavailable).WithCode(code)
}

// AsServiceUnavailable converts the error to a service unavailable error.
func (e *AppError) AsServiceUnavailable() *AppError {
	return e.AsUnavailableWithCode(CodeServiceUnavailable)
}

// AsExternalWithCode converts the error to an external service error with a specific code.
func (e *AppError) AsExternalWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeExternal).WithCode(code)
}

// AsServiceTimeout converts the error to a service timeout error.
func (e *AppError) AsServiceTimeout() *AppError {
	return e.AsExternalWithCode(CodeExternalTimeout)
}

// AsServiceUnavailable converts the error to a service unavailable error.
func (e *AppError) AsExternalServiceUnavailable() *AppError {
	return e.AsExternalWithCode(CodeExternalUnavailable)
}

// AsRateLimitWithCode converts the error to a rate limit error with a specific code.
func (e *AppError) AsRateLimitWithCode(code string) *AppError {
	return e.setTypeAndStatus(ErrorTypeRateLimit).WithCode(code)
}

// AsTooManyRequests converts the error to a too many requests error.
func (e *AppError) AsTooManyRequests() *AppError {
	return e.AsRateLimitWithCode(CodeRateLimitExceeded)
}

// AsConfiguration converts the error to a configuration error.
func (e *AppError) AsConfiguration() *AppError {
	return e.AsInternalWithCode(CodeConfigurationError)
}

// AsTimeout converts the error to a timeout error.
func (e *AppError) AsTimeout() *AppError {
	return e.AsInternalWithCode(CodeInternalTimeout)
}

// setTypeAndStatus sets the error type and updates the HTTP status.
func (e *AppError) setTypeAndStatus(errorType ErrorType) *AppError {
	if e == nil {
		return nil
	}
	e.Type = errorType
	e.HTTPStatus = errorType.DefaultHTTPStatus()
	return e
}
