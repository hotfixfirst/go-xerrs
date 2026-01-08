# Wrapping Example

This example demonstrates the `xerrs` error wrapping and automatic detection functionality.

## Run

```bash
cd _examples/wrapping
go run main.go
```

## Features Demonstrated

| # | Feature | Function |
| - | ------- | -------- |
| 1 | Wrap standard errors | `Wrap()` |
| 2 | Auto-detect GORM errors | `Wrap()` with `gorm.ErrRecordNotFound` |
| 3 | Auto-detect SQL errors | `Wrap()` with `sql.ErrNoRows` |
| 4 | Auto-detect context deadline | `Wrap()` with `context.DeadlineExceeded` |
| 5 | Auto-detect context canceled | `Wrap()` with `context.Canceled` |
| 6 | Auto-detect JSON errors | `Wrap()` with JSON unmarshal error |
| 7 | Wrap and convert type | `Wrap()` + `AsInvalidInput()` |
| 8 | Error unwrapping | `Unwrap()`, `UnwrapAll()` |
| 9 | Type checking | `IsAppError()` |
| 10 | Convert to AppError | `AsAppError()` |
| 11 | Wrapping nil error | `Wrap(nil, ...)` |
| 12 | Re-wrap preserves type | `Wrap()` on existing `AppError` |

## Auto-Detection Support

The `Wrap()` function automatically detects error types from:

| Error Source | Detected Type | Detected Code |
| ------------ | ------------- | ------------- |
| `gorm.ErrRecordNotFound` | NOT_FOUND | RESOURCE_NOT_FOUND |
| `sql.ErrNoRows` | NOT_FOUND | RESOURCE_NOT_FOUND |
| `context.DeadlineExceeded` | INTERNAL | INTERNAL_TIMEOUT |
| `context.Canceled` | INTERNAL | OPERATION_CANCELED |
| JSON unmarshal errors | VALIDATION | INVALID_FORMAT |
| "duplicate key" messages | CONFLICT | RESOURCE_EXISTS |
| "required" messages | VALIDATION | REQUIRED_FIELD |
| And many more... | ... | ... |

## Sample Output

```text
=== Error Wrapping Examples ===

1. Wrapping Standard Errors
----------------------------
Error: [INTERNAL] INTERNAL_ERROR: failed to process request
Type: INTERNAL
Code: INTERNAL_ERROR

2. Auto-Detection: GORM Record Not Found
-----------------------------------------
Error: [NOT_FOUND] RESOURCE_NOT_FOUND: user lookup failed
Type: NOT_FOUND (auto-detected)
Code: RESOURCE_NOT_FOUND (auto-detected)
HTTP Status: 404

3. Auto-Detection: SQL No Rows
-------------------------------
Error: [NOT_FOUND] RESOURCE_NOT_FOUND: record not found in database
Type: NOT_FOUND (auto-detected)
Code: RESOURCE_NOT_FOUND (auto-detected)

4. Auto-Detection: Context Deadline
------------------------------------
Error: [INTERNAL] INTERNAL_TIMEOUT: operation timed out
Type: INTERNAL (auto-detected)
Code: INTERNAL_TIMEOUT (auto-detected)

5. Auto-Detection: Context Canceled
------------------------------------
Error: [INTERNAL] OPERATION_CANCELED: request was canceled
Type: INTERNAL (auto-detected)
Code: OPERATION_CANCELED (auto-detected)

6. Auto-Detection: JSON Parse Error
------------------------------------
Error: [VALIDATION] INVALID_FORMAT: failed to parse request body
Type: VALIDATION (auto-detected)
Code: INVALID_FORMAT (auto-detected)
HTTP Status: 400

7. Wrap and Convert Type
------------------------
Error: [VALIDATION] INVALID_INPUT: validation failed
Type: VALIDATION
Code: INVALID_INPUT

8. Error Unwrapping
-------------------
Original error: [INTERNAL] INTERNAL_ERROR: layer one
Unwrapped (immediate cause): layer one: root cause
UnwrapAll (root cause): root cause

9. Error Type Checking
----------------------
Is AppError (xerrs.New): true
Is AppError (errors.New): false

10. Convert to AppError
-----------------------
Successfully converted to AppError
Type: INTERNAL
Code: INTERNAL_ERROR

11. Wrapping Nil Error
----------------------
Error: [INTERNAL] INTERNAL_ERROR: no error occurred
Type: INTERNAL (default)

12. Re-wrapping AppError (Preserves Type)
------------------------------------------
Original Type: NOT_FOUND
Re-wrapped Type: NOT_FOUND (preserved)
Original Code: RESOURCE_NOT_FOUND
Re-wrapped Code: RESOURCE_NOT_FOUND (preserved)

=== End of Examples ===
```
