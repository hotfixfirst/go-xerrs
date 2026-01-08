// Package main demonstrates basic usage of the xerrs package.
package main

import (
	"fmt"

	"github.com/hotfixfirst/go-xerrs"
)

func main() {
	fmt.Println("=== Basic Error Examples ===")
	fmt.Println()

	// Example 1: Create a simple error with New()
	fmt.Println("1. Simple Error with New()")
	fmt.Println("---------------------------")
	err := xerrs.New("something went wrong")
	fmt.Printf("Error: %s\n", err.Error())
	fmt.Printf("Type: %s\n", err.Type)
	fmt.Printf("Code: %s\n", err.Code)
	fmt.Printf("HTTP Status: %d\n", err.GetHTTPStatus())
	fmt.Println()

	// Example 2: Create an error with NewAppError()
	fmt.Println("2. Structured Error with NewAppError()")
	fmt.Println("---------------------------------------")
	appErr := xerrs.NewAppError(
		xerrs.ErrorTypeValidation,
		xerrs.CodeInvalidInput,
		"email address is invalid",
	)
	fmt.Printf("Error: %s\n", appErr.Error())
	fmt.Printf("Type: %s\n", appErr.Type)
	fmt.Printf("Code: %s\n", appErr.Code)
	fmt.Printf("HTTP Status: %d\n", appErr.GetHTTPStatus())
	fmt.Println()

	// Example 3: Add details to an error
	fmt.Println("3. Error with Details")
	fmt.Println("---------------------")
	detailedErr := xerrs.New("validation failed").
		WithDetails("field 'email' must be a valid email address")
	fmt.Printf("Error: %s\n", detailedErr.Error())
	fmt.Printf("Details: %s\n", detailedErr.Details)
	fmt.Println()

	// Example 4: Override HTTP status
	fmt.Println("4. Custom HTTP Status")
	fmt.Println("---------------------")
	customErr := xerrs.New("custom error").
		WithHTTPStatus(422)
	fmt.Printf("Error: %s\n", customErr.Error())
	fmt.Printf("HTTP Status: %d\n", customErr.GetHTTPStatus())
	fmt.Println()

	// Example 5: Chain multiple configuration methods
	fmt.Println("5. Fluent Configuration")
	fmt.Println("-----------------------")
	fluentErr := xerrs.New("user creation failed").
		WithCode("USER_CREATE_FAILED").
		WithDetails("email already registered").
		WithHTTPStatus(409)
	fmt.Printf("Error: %s\n", fluentErr.Error())
	fmt.Printf("Code: %s\n", fluentErr.Code)
	fmt.Printf("Details: %s\n", fluentErr.Details)
	fmt.Printf("HTTP Status: %d\n", fluentErr.GetHTTPStatus())
	fmt.Println()

	// Example 6: Check error type and code
	fmt.Println("6. Error Type and Code Checking")
	fmt.Println("--------------------------------")
	validationErr := xerrs.NewAppError(
		xerrs.ErrorTypeValidation,
		xerrs.CodeRequiredField,
		"name is required",
	)
	fmt.Printf("Is Validation Type: %v\n", validationErr.IsType(xerrs.ErrorTypeValidation))
	fmt.Printf("Is NotFound Type: %v\n", validationErr.IsType(xerrs.ErrorTypeNotFound))
	fmt.Printf("Has REQUIRED_FIELD code: %v\n", validationErr.HasCode(xerrs.CodeRequiredField))
	fmt.Printf("Has INVALID_INPUT code: %v\n", validationErr.HasCode(xerrs.CodeInvalidInput))
	fmt.Println()

	// Example 7: Get stack trace
	fmt.Println("7. Stack Trace")
	fmt.Println("--------------")
	stackErr := xerrs.New("error with stack trace")
	lines := stackErr.GetStackTraceLines()
	fmt.Printf("Stack trace has %d lines\n", len(lines))
	if len(lines) > 0 {
		fmt.Printf("First line: %s\n", lines[0])
	}
	fmt.Println()

	fmt.Println("=== End of Examples ===")
}
