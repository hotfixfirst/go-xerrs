// Package main demonstrates error wrapping and automatic detection in xerrs.
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/hotfixfirst/go-xerrs"
)

func main() {
	fmt.Println("=== Error Wrapping Examples ===")
	fmt.Println()

	// Example 1: Wrap a standard error
	fmt.Println("1. Wrapping Standard Errors")
	fmt.Println("----------------------------")
	stdErr := errors.New("original error")
	wrappedErr := xerrs.Wrap(stdErr, "failed to process request")
	fmt.Printf("Error: %s\n", wrappedErr.Error())
	fmt.Printf("Type: %s\n", wrappedErr.Type)
	fmt.Printf("Code: %s\n", wrappedErr.Code)
	fmt.Println()

	// Example 2: Wrap GORM record not found error
	fmt.Println("2. Auto-Detection: GORM Record Not Found")
	fmt.Println("-----------------------------------------")
	gormErr := xerrs.Wrap(gorm.ErrRecordNotFound, "user lookup failed")
	fmt.Printf("Error: %s\n", gormErr.Error())
	fmt.Printf("Type: %s (auto-detected)\n", gormErr.Type)
	fmt.Printf("Code: %s (auto-detected)\n", gormErr.Code)
	fmt.Printf("HTTP Status: %d\n", gormErr.GetHTTPStatus())
	fmt.Println()

	// Example 3: Wrap SQL no rows error
	fmt.Println("3. Auto-Detection: SQL No Rows")
	fmt.Println("-------------------------------")
	sqlErr := xerrs.Wrap(sql.ErrNoRows, "record not found in database")
	fmt.Printf("Error: %s\n", sqlErr.Error())
	fmt.Printf("Type: %s (auto-detected)\n", sqlErr.Type)
	fmt.Printf("Code: %s (auto-detected)\n", sqlErr.Code)
	fmt.Println()

	// Example 4: Wrap context deadline exceeded
	fmt.Println("4. Auto-Detection: Context Deadline")
	fmt.Println("------------------------------------")
	ctxErr := xerrs.Wrap(context.DeadlineExceeded, "operation timed out")
	fmt.Printf("Error: %s\n", ctxErr.Error())
	fmt.Printf("Type: %s (auto-detected)\n", ctxErr.Type)
	fmt.Printf("Code: %s (auto-detected)\n", ctxErr.Code)
	fmt.Println()

	// Example 5: Wrap context canceled
	fmt.Println("5. Auto-Detection: Context Canceled")
	fmt.Println("------------------------------------")
	cancelErr := xerrs.Wrap(context.Canceled, "request was canceled")
	fmt.Printf("Error: %s\n", cancelErr.Error())
	fmt.Printf("Type: %s (auto-detected)\n", cancelErr.Type)
	fmt.Printf("Code: %s (auto-detected)\n", cancelErr.Code)
	fmt.Println()

	// Example 6: Wrap JSON unmarshal error
	fmt.Println("6. Auto-Detection: JSON Parse Error")
	fmt.Println("------------------------------------")
	var data map[string]any
	jsonErr := json.Unmarshal([]byte("invalid json"), &data)
	parsedErr := xerrs.Wrap(jsonErr, "failed to parse request body")
	fmt.Printf("Error: %s\n", parsedErr.Error())
	fmt.Printf("Type: %s (auto-detected)\n", parsedErr.Type)
	fmt.Printf("Code: %s (auto-detected)\n", parsedErr.Code)
	fmt.Printf("HTTP Status: %d\n", parsedErr.GetHTTPStatus())
	fmt.Println()

	// Example 7: Wrap and convert type
	fmt.Println("7. Wrap and Convert Type")
	fmt.Println("------------------------")
	originalErr := errors.New("invalid input data")
	convertedErr := xerrs.Wrap(originalErr, "validation failed").AsInvalidInput()
	fmt.Printf("Error: %s\n", convertedErr.Error())
	fmt.Printf("Type: %s\n", convertedErr.Type)
	fmt.Printf("Code: %s\n", convertedErr.Code)
	fmt.Println()

	// Example 8: Error unwrapping
	fmt.Println("8. Error Unwrapping")
	fmt.Println("-------------------")
	rootErr := errors.New("root cause")
	layeredErr := xerrs.Wrap(rootErr, "layer one")
	fmt.Printf("Original error: %s\n", layeredErr.Error())
	fmt.Printf("Unwrapped (immediate cause): %v\n", layeredErr.Unwrap())
	fmt.Printf("UnwrapAll (root cause): %v\n", layeredErr.UnwrapAll())
	fmt.Println()

	// Example 9: Check if error is AppError
	fmt.Println("9. Error Type Checking")
	fmt.Println("----------------------")
	appErr := xerrs.New("app error")
	stdError := errors.New("standard error")
	fmt.Printf("Is AppError (xerrs.New): %v\n", xerrs.IsAppError(appErr))
	fmt.Printf("Is AppError (errors.New): %v\n", xerrs.IsAppError(stdError))
	fmt.Println()

	// Example 10: Convert to AppError
	fmt.Println("10. Convert to AppError")
	fmt.Println("-----------------------")
	someErr := xerrs.Wrap(errors.New("original"), "wrapped")
	if converted, ok := xerrs.AsAppError(someErr); ok {
		fmt.Printf("Successfully converted to AppError\n")
		fmt.Printf("Type: %s\n", converted.Type)
		fmt.Printf("Code: %s\n", converted.Code)
	}
	fmt.Println()

	// Example 11: Wrapping nil error
	fmt.Println("11. Wrapping Nil Error")
	fmt.Println("----------------------")
	nilWrapped := xerrs.Wrap(nil, "no error occurred")
	fmt.Printf("Error: %s\n", nilWrapped.Error())
	fmt.Printf("Type: %s (default)\n", nilWrapped.Type)
	fmt.Println()

	// Example 12: Re-wrapping an AppError preserves type
	fmt.Println("12. Re-wrapping AppError (Preserves Type)")
	fmt.Println("------------------------------------------")
	firstErr := xerrs.New("first error").AsResourceNotFound()
	reWrapped := xerrs.Wrap(firstErr, "could not find user")
	fmt.Printf("Original Type: %s\n", firstErr.Type)
	fmt.Printf("Re-wrapped Type: %s (preserved)\n", reWrapped.Type)
	fmt.Printf("Original Code: %s\n", firstErr.Code)
	fmt.Printf("Re-wrapped Code: %s (preserved)\n", reWrapped.Code)
	fmt.Println()

	fmt.Println("=== End of Examples ===")
}
