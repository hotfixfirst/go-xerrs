// Package main demonstrates error type conversion and chaining in xerrs.
package main

import (
	"fmt"

	"github.com/hotfixfirst/go-xerrs"
)

func main() {
	fmt.Println("=== Error Chaining Examples ===")
	fmt.Println()

	// Example 1: Validation errors
	fmt.Println("1. Validation Errors")
	fmt.Println("--------------------")
	validationErr := xerrs.New("invalid email format").AsInvalidFormat()
	fmt.Printf("Error: %s\n", validationErr.Error())
	fmt.Printf("HTTP Status: %d\n", validationErr.GetHTTPStatus())

	requiredErr := xerrs.New("username is required").AsRequiredField()
	fmt.Printf("Error: %s\n", requiredErr.Error())

	rangeErr := xerrs.New("age must be between 18 and 100").AsInvalidRange()
	fmt.Printf("Error: %s\n", rangeErr.Error())
	fmt.Println()

	// Example 2: Authentication errors
	fmt.Println("2. Authentication Errors")
	fmt.Println("------------------------")
	credErr := xerrs.New("wrong password").AsInvalidCredentials()
	fmt.Printf("Error: %s\n", credErr.Error())
	fmt.Printf("HTTP Status: %d\n", credErr.GetHTTPStatus())

	tokenErr := xerrs.New("JWT token has expired").AsTokenExpired()
	fmt.Printf("Error: %s\n", tokenErr.Error())

	loginErr := xerrs.New("please log in to continue").AsLoginRequired()
	fmt.Printf("Error: %s\n", loginErr.Error())
	fmt.Println()

	// Example 3: Authorization errors
	fmt.Println("3. Authorization Errors")
	fmt.Println("-----------------------")
	accessErr := xerrs.New("cannot access this resource").AsAccessDenied()
	fmt.Printf("Error: %s\n", accessErr.Error())
	fmt.Printf("HTTP Status: %d\n", accessErr.GetHTTPStatus())

	permErr := xerrs.New("admin role required").AsInsufficientPermissions()
	fmt.Printf("Error: %s\n", permErr.Error())

	forbiddenErr := xerrs.New("resource is forbidden").AsResourceForbidden()
	fmt.Printf("Error: %s\n", forbiddenErr.Error())
	fmt.Println()

	// Example 4: Not Found errors
	fmt.Println("4. Not Found Errors")
	fmt.Println("-------------------")
	notFoundErr := xerrs.New("user not found").AsResourceNotFound()
	fmt.Printf("Error: %s\n", notFoundErr.Error())
	fmt.Printf("HTTP Status: %d\n", notFoundErr.GetHTTPStatus())
	fmt.Println()

	// Example 5: Conflict errors
	fmt.Println("5. Conflict Errors")
	fmt.Println("------------------")
	existsErr := xerrs.New("user already exists").AsResourceExists()
	fmt.Printf("Error: %s\n", existsErr.Error())
	fmt.Printf("HTTP Status: %d\n", existsErr.GetHTTPStatus())
	fmt.Println()

	// Example 6: Database errors
	fmt.Println("6. Database Errors")
	fmt.Println("------------------")
	dbErr := xerrs.New("failed to execute query").AsDatabaseError()
	fmt.Printf("Error: %s\n", dbErr.Error())
	fmt.Printf("HTTP Status: %d\n", dbErr.GetHTTPStatus())

	connErr := xerrs.New("cannot connect to database").AsDatabaseConnection()
	fmt.Printf("Error: %s\n", connErr.Error())

	constraintErr := xerrs.New("foreign key violation").AsDatabaseConstraint()
	fmt.Printf("Error: %s\n", constraintErr.Error())
	fmt.Println()

	// Example 7: External service errors
	fmt.Println("7. External Service Errors")
	fmt.Println("--------------------------")
	timeoutErr := xerrs.New("payment gateway timeout").AsServiceTimeout()
	fmt.Printf("Error: %s\n", timeoutErr.Error())
	fmt.Printf("HTTP Status: %d\n", timeoutErr.GetHTTPStatus())

	unavailErr := xerrs.New("email service unavailable").AsExternalServiceUnavailable()
	fmt.Printf("Error: %s\n", unavailErr.Error())
	fmt.Println()

	// Example 8: Rate limit errors
	fmt.Println("8. Rate Limit Errors")
	fmt.Println("--------------------")
	rateLimitErr := xerrs.New("too many requests").AsTooManyRequests()
	fmt.Printf("Error: %s\n", rateLimitErr.Error())
	fmt.Printf("HTTP Status: %d\n", rateLimitErr.GetHTTPStatus())
	fmt.Println()

	// Example 9: Custom code with type conversion
	fmt.Println("9. Custom Code with Type Conversion")
	fmt.Println("------------------------------------")
	customErr := xerrs.New("custom validation error").
		AsValidationWithCode("CUSTOM_VALIDATION_CODE")
	fmt.Printf("Error: %s\n", customErr.Error())
	fmt.Printf("Code: %s\n", customErr.Code)
	fmt.Println()

	// Example 10: Chaining with additional details
	fmt.Println("10. Chaining with Details")
	fmt.Println("-------------------------")
	chainedErr := xerrs.New("user creation failed").
		AsConflictWithCode("USER_EXISTS").
		WithDetails("email 'test@example.com' is already registered").
		WithHTTPStatus(409)
	fmt.Printf("Error: %s\n", chainedErr.Error())
	fmt.Printf("Type: %s\n", chainedErr.Type)
	fmt.Printf("Code: %s\n", chainedErr.Code)
	fmt.Printf("Details: %s\n", chainedErr.Details)
	fmt.Printf("HTTP Status: %d\n", chainedErr.GetHTTPStatus())
	fmt.Println()

	fmt.Println("=== End of Examples ===")
}
