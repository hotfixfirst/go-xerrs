# Basic Example

This example demonstrates the basic `xerrs` error creation and configuration functionality.

## Run

```bash
cd _examples/basic
go run main.go
```

## Features Demonstrated

| # | Feature | Function |
| - | ------- | -------- |
| 1 | Simple error creation | `New()` |
| 2 | Structured error creation | `NewAppError()` |
| 3 | Adding details | `WithDetails()` |
| 4 | Custom HTTP status | `WithHTTPStatus()` |
| 5 | Fluent configuration | Method chaining |
| 6 | Type and code checking | `IsType()`, `HasCode()` |
| 7 | Stack trace access | `GetStackTraceLines()` |

## Sample Output

```text
=== Basic Error Examples ===

1. Simple Error with New()
---------------------------
Error: [INTERNAL] INTERNAL_ERROR: something went wrong
Type: INTERNAL
Code: INTERNAL_ERROR
HTTP Status: 500

2. Structured Error with NewAppError()
---------------------------------------
Error: [VALIDATION] INVALID_INPUT: email address is invalid
Type: VALIDATION
Code: INVALID_INPUT
HTTP Status: 400

3. Error with Details
---------------------
Error: [INTERNAL] INTERNAL_ERROR: validation failed - field 'email' must be a valid email address
Details: field 'email' must be a valid email address

4. Custom HTTP Status
---------------------
Error: [INTERNAL] INTERNAL_ERROR: custom error
HTTP Status: 422

5. Fluent Configuration
-----------------------
Error: [INTERNAL] USER_CREATE_FAILED: user creation failed - email already registered
Code: USER_CREATE_FAILED
Details: email already registered
HTTP Status: 409

6. Error Type and Code Checking
--------------------------------
Is Validation Type: true
Is NotFound Type: false
Has REQUIRED_FIELD code: true
Has INVALID_INPUT code: false

7. Stack Trace
--------------
Stack trace has X lines
First line: error with stack trace

=== End of Examples ===
```
