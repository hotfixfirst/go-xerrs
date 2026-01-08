# Chaining Example

This example demonstrates the `xerrs` error type conversion and fluent chaining functionality.

## Run

```bash
cd _examples/chaining
go run main.go
```

## Features Demonstrated

| # | Feature | Function |
| - | ------- | -------- |
| 1 | Validation errors | `AsInvalidFormat()`, `AsRequiredField()`, `AsInvalidRange()` |
| 2 | Authentication errors | `AsInvalidCredentials()`, `AsTokenExpired()`, `AsLoginRequired()` |
| 3 | Authorization errors | `AsAccessDenied()`, `AsInsufficientPermissions()`, `AsResourceForbidden()` |
| 4 | Not found errors | `AsResourceNotFound()` |
| 5 | Conflict errors | `AsResourceExists()` |
| 6 | Database errors | `AsDatabaseError()`, `AsDatabaseConnection()`, `AsDatabaseConstraint()` |
| 7 | External service errors | `AsServiceTimeout()`, `AsExternalServiceUnavailable()` |
| 8 | Rate limit errors | `AsTooManyRequests()` |
| 9 | Custom code with type | `AsValidationWithCode()`, `AsAuthenticationWithCode()`, etc. |
| 10 | Chaining with details | Method chaining with `WithDetails()` |

## Sample Output

```text
=== Error Chaining Examples ===

1. Validation Errors
--------------------
Error: [VALIDATION] INVALID_FORMAT: invalid email format
HTTP Status: 400
Error: [VALIDATION] REQUIRED_FIELD: username is required
Error: [VALIDATION] INVALID_RANGE: age must be between 18 and 100

2. Authentication Errors
------------------------
Error: [AUTHENTICATION] INVALID_CREDENTIALS: wrong password
HTTP Status: 401
Error: [AUTHENTICATION] TOKEN_EXPIRED: JWT token has expired
Error: [AUTHENTICATION] LOGIN_REQUIRED: please log in to continue

3. Authorization Errors
-----------------------
Error: [AUTHORIZATION] ACCESS_DENIED: cannot access this resource
HTTP Status: 403
Error: [AUTHORIZATION] INSUFFICIENT_PERMISSIONS: admin role required
Error: [AUTHORIZATION] RESOURCE_FORBIDDEN: resource is forbidden

4. Not Found Errors
-------------------
Error: [NOT_FOUND] RESOURCE_NOT_FOUND: user not found
HTTP Status: 404

5. Conflict Errors
------------------
Error: [CONFLICT] RESOURCE_EXISTS: user already exists
HTTP Status: 409

6. Database Errors
------------------
Error: [INTERNAL] DATABASE_ERROR: failed to execute query
HTTP Status: 500
Error: [INTERNAL] DATABASE_CONNECTION: cannot connect to database
Error: [INTERNAL] DATABASE_CONSTRAINT: foreign key violation

7. External Service Errors
--------------------------
Error: [EXTERNAL] EXTERNAL_TIMEOUT: payment gateway timeout
HTTP Status: 502
Error: [EXTERNAL] EXTERNAL_UNAVAILABLE: email service unavailable

8. Rate Limit Errors
--------------------
Error: [RATE_LIMIT] RATE_LIMIT_EXCEEDED: too many requests
HTTP Status: 429

9. Custom Code with Type Conversion
------------------------------------
Error: [VALIDATION] CUSTOM_VALIDATION_CODE: custom validation error
Code: CUSTOM_VALIDATION_CODE

10. Chaining with Details
-------------------------
Error: [CONFLICT] USER_EXISTS: user creation failed - email 'test@example.com' is already registered
Type: CONFLICT
Code: USER_EXISTS
Details: email 'test@example.com' is already registered
HTTP Status: 409

=== End of Examples ===
```
