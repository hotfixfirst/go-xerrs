# go-xerrs

A structured error handling package for Go applications with HTTP status mapping, error chaining, and stack trace support.

## Installation

```bash
go get github.com/hotfixfirst/go-xerrs
```

Or with a specific version:

```bash
go get github.com/hotfixfirst/go-xerrs@v1.0.0
```

## Quick Start

```go
import "github.com/hotfixfirst/go-xerrs"

// Create a simple error
err := xerrs.New("something went wrong")

// Create a validation error with chaining
err := xerrs.New("invalid email format").AsInvalidFormat()

// Wrap an existing error with auto-detection
err := xerrs.Wrap(gorm.ErrRecordNotFound, "user not found")
// Automatically detects as NOT_FOUND with HTTP 404
```

## Features

| Feature | Description | Documentation |
| ------- | ----------- | ------------- |
| [Error Creation](#error-creation) | Create structured errors with type, code, and message | [Examples](./_examples/basic/) |
| [Error Chaining](#error-chaining) | Fluent API for error type conversion | [Examples](./_examples/chaining/) |
| [Error Wrapping](#error-wrapping) | Wrap existing errors with auto-detection | [Examples](./_examples/wrapping/) |
| [HTTP Status Mapping](#http-status-mapping) | Automatic HTTP status codes based on error type | - |
| [Stack Traces](#stack-traces) | Built-in stack trace support via cockroachdb/errors | - |

## Error Creation

### Functions

| Function | Description |
| -------- | ----------- |
| `New(message)` | Create a new error with default internal type |
| `NewAppError(type, code, message)` | Create a structured error with specific type and code |
| `Wrap(err, message)` | Wrap an existing error with auto-detection |

### Examples

```go
// Simple error
err := xerrs.New("database connection failed")
// [INTERNAL] INTERNAL_ERROR: database connection failed

// Structured error
err := xerrs.NewAppError(
    xerrs.ErrorTypeValidation,
    xerrs.CodeInvalidInput,
    "email is required",
)
// [VALIDATION] INVALID_INPUT: email is required

// Error with details
err := xerrs.New("validation failed").
    WithDetails("field 'email' must be a valid email address")
```

## Error Chaining

Convert errors to specific types using fluent API methods.

### Validation Errors (400 Bad Request)

| Method | Code |
| ------ | ---- |
| `AsValidationError()` | VALIDATION_ERROR |
| `AsInvalidInput()` | INVALID_INPUT |
| `AsRequiredField()` | REQUIRED_FIELD |
| `AsInvalidFormat()` | INVALID_FORMAT |
| `AsInvalidRange()` | INVALID_RANGE |
| `AsValidationWithCode(code)` | Custom code |

### Authentication Errors (401 Unauthorized)

| Method | Code |
| ------ | ---- |
| `AsAuthentication()` | (preserves code) |
| `AsInvalidCredentials()` | INVALID_CREDENTIALS |
| `AsTokenExpired()` | TOKEN_EXPIRED |
| `AsTokenInvalid()` | TOKEN_INVALID |
| `AsLoginRequired()` | LOGIN_REQUIRED |
| `AsAuthenticationWithCode(code)` | Custom code |

### Authorization Errors (403 Forbidden)

| Method | Code |
| ------ | ---- |
| `AsAccessDenied()` | ACCESS_DENIED |
| `AsInsufficientPermissions()` | INSUFFICIENT_PERMISSIONS |
| `AsResourceForbidden()` | RESOURCE_FORBIDDEN |
| `AsAuthorizationWithCode(code)` | Custom code |

### Not Found Errors (404 Not Found)

| Method | Code |
| ------ | ---- |
| `AsResourceNotFound()` | RESOURCE_NOT_FOUND |
| `AsNotFoundWithCode(code)` | Custom code |

### Conflict Errors (409 Conflict)

| Method | Code |
| ------ | ---- |
| `AsResourceExists()` | RESOURCE_EXISTS |
| `AsConflictWithCode(code)` | Custom code |

### Rate Limit Errors (429 Too Many Requests)

| Method | Code |
| ------ | ---- |
| `AsTooManyRequests()` | RATE_LIMIT_EXCEEDED |
| `AsRateLimitWithCode(code)` | Custom code |

### Internal Errors (500 Internal Server Error)

| Method | Code |
| ------ | ---- |
| `AsDatabaseError()` | DATABASE_ERROR |
| `AsDatabaseConnection()` | DATABASE_CONNECTION |
| `AsDatabaseTimeout()` | INTERNAL_TIMEOUT |
| `AsDatabaseConstraint()` | DATABASE_CONSTRAINT |
| `AsConfiguration()` | CONFIGURATION_ERROR |
| `AsTimeout()` | INTERNAL_TIMEOUT |
| `AsInternalWithCode(code)` | Custom code |

### External Service Errors (502 Bad Gateway)

| Method | Code |
| ------ | ---- |
| `AsServiceTimeout()` | EXTERNAL_TIMEOUT |
| `AsExternalServiceUnavailable()` | EXTERNAL_UNAVAILABLE |
| `AsExternalWithCode(code)` | Custom code |

### Service Unavailable Errors (503 Service Unavailable)

| Method | Code |
| ------ | ---- |
| `AsServiceUnavailable()` | SERVICE_UNAVAILABLE |
| `AsUnavailableWithCode(code)` | Custom code |

### Chaining Examples

```go
// Validation error
err := xerrs.New("email is invalid").AsInvalidFormat()
// [VALIDATION] INVALID_FORMAT: email is invalid (HTTP 400)

// Authentication error
err := xerrs.New("session expired").AsTokenExpired()
// [AUTHENTICATION] TOKEN_EXPIRED: session expired (HTTP 401)

// Not found error
err := xerrs.New("user not found").AsResourceNotFound()
// [NOT_FOUND] RESOURCE_NOT_FOUND: user not found (HTTP 404)

// Custom code
err := xerrs.New("custom error").AsValidationWithCode("CUSTOM_CODE")
// [VALIDATION] CUSTOM_CODE: custom error (HTTP 400)
```

## Error Wrapping

Wrap existing errors with automatic type and code detection.

### Auto-Detection Support

| Error Source | Detected Type | Detected Code | HTTP Status |
| ------------ | ------------- | ------------- | ----------- |
| `gorm.ErrRecordNotFound` | NOT_FOUND | RESOURCE_NOT_FOUND | 404 |
| `sql.ErrNoRows` | NOT_FOUND | RESOURCE_NOT_FOUND | 404 |
| `context.DeadlineExceeded` | INTERNAL | INTERNAL_TIMEOUT | 500 |
| `context.Canceled` | INTERNAL | OPERATION_CANCELED | 500 |
| JSON unmarshal errors | VALIDATION | INVALID_FORMAT | 400 |
| "duplicate key" errors | CONFLICT | RESOURCE_EXISTS | 409 |
| "required" errors | VALIDATION | REQUIRED_FIELD | 400 |
| "unauthorized" errors | AUTHENTICATION | AUTH_REQUIRED | 401 |
| "forbidden" errors | AUTHORIZATION | ACCESS_DENIED | 403 |
| "timeout" errors | INTERNAL | INTERNAL_TIMEOUT | 500 |

### Wrapping Examples

```go
// Wrap with auto-detection
err := xerrs.Wrap(gorm.ErrRecordNotFound, "user not found")
// Type: NOT_FOUND, Code: RESOURCE_NOT_FOUND, HTTP: 404

// Wrap and convert
err := xerrs.Wrap(originalErr, "validation failed").AsInvalidInput()

// Re-wrapping preserves type
firstErr := xerrs.New("not found").AsResourceNotFound()
reWrapped := xerrs.Wrap(firstErr, "user lookup failed")
// Type is preserved as NOT_FOUND
```

## HTTP Status Mapping

Error types automatically map to HTTP status codes.

| Error Type | HTTP Status |
| ---------- | ----------- |
| VALIDATION | 400 Bad Request |
| AUTHENTICATION | 401 Unauthorized |
| AUTHORIZATION | 403 Forbidden |
| NOT_FOUND | 404 Not Found |
| CONFLICT | 409 Conflict |
| RATE_LIMIT | 429 Too Many Requests |
| INTERNAL | 500 Internal Server Error |
| EXTERNAL | 502 Bad Gateway |
| UNAVAILABLE | 503 Service Unavailable |

```go
err := xerrs.New("not found").AsResourceNotFound()
status := err.GetHTTPStatus() // 404

// Override status
err := xerrs.New("custom").WithHTTPStatus(422)
status := err.GetHTTPStatus() // 422
```

## Configuration Methods

| Method | Description |
| ------ | ----------- |
| `WithType(type)` | Set error type |
| `WithCode(code)` | Set error code |
| `WithMessage(message)` | Set error message |
| `WithDetails(details)` | Add detailed information |
| `WithHTTPStatus(status)` | Override HTTP status |
| `WithCause(err)` | Set underlying cause |
| `WithCodeAndMessage(code, message)` | Set both code and message |

## Inspection Methods

| Method | Description |
| ------ | ----------- |
| `Error()` | Get formatted error string |
| `GetHTTPStatus()` | Get HTTP status code |
| `IsType(type)` | Check if error is of specific type |
| `HasCode(code)` | Check if error has specific code |
| `Unwrap()` | Get immediate underlying cause |
| `UnwrapAll()` | Get root cause |
| `Cause()` | Get direct cause |
| `GetStackTrace()` | Get full stack trace string |
| `GetStackTraceLines()` | Get stack trace as lines |

## Helper Functions

| Function | Description |
| -------- | ----------- |
| `IsAppError(err)` | Check if error is an AppError |
| `AsAppError(err)` | Convert error to AppError if possible |

## Stack Traces

Built-in stack trace support via `cockroachdb/errors`.

```go
err := xerrs.New("something failed")

// Get full stack trace
trace := err.GetStackTrace()

// Get stack trace as lines
lines := err.GetStackTraceLines()
for _, line := range lines {
    fmt.Println(line)
}
```

## Error Codes

### Validation Codes

- `VALIDATION_ERROR`, `INVALID_INPUT`, `REQUIRED_FIELD`, `INVALID_FORMAT`, `INVALID_RANGE`

### Authentication Codes

- `INVALID_CREDENTIALS`, `TOKEN_EXPIRED`, `TOKEN_INVALID`, `LOGIN_REQUIRED`, `AUTH_REQUIRED`

### Authorization Codes

- `ACCESS_DENIED`, `INSUFFICIENT_PERMISSIONS`, `RESOURCE_FORBIDDEN`, `INSUFFICIENT_ROLE`

### Resource Codes

- `RESOURCE_NOT_FOUND`, `RESOURCE_EXISTS`

### Internal Codes

- `INTERNAL_ERROR`, `DATABASE_ERROR`, `DATABASE_CONNECTION`, `DATABASE_CONSTRAINT`, `INTERNAL_TIMEOUT`, `CONFIGURATION_ERROR`, `OPERATION_CANCELED`

### External Codes

- `EXTERNAL_ERROR`, `EXTERNAL_TIMEOUT`, `EXTERNAL_UNAVAILABLE`, `SERVICE_UNAVAILABLE`

### Rate Limit Codes

- `RATE_LIMIT_EXCEEDED`

## Runnable Examples

See the [_examples](./_examples/) directory for runnable examples:

- [basic](./_examples/basic/) - Basic error creation and configuration
- [chaining](./_examples/chaining/) - Fluent error type conversion
- [wrapping](./_examples/wrapping/) - Error wrapping and auto-detection

## License

MIT License - see [LICENSE](LICENSE) for details.
