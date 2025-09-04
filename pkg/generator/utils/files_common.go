package utils

import (
	"fmt"
	"strings"
)

func NewMainFile(pkg string) string {
	lines := []string{
		"package main",
		"",
		fmt.Sprintf(`import "%s/internal/app"`, pkg),
		"",
		"func main() {",
		"\tapp.RunApp()",
		"}",
	}

	return strings.Join(lines, "\n")
}

func NewInternalResponseFile() string {
	lines := []string{
		"package responses",
		"",
		"type InternalError struct {",
		"\tMessage string `json:\"message\"`",
		"}",
		"",

		"func NewInternalError(err ...string) *InternalError {",
		"\tif len(err) > 0 {",
		"\t\treturn &InternalError{",
		"\t\t\tMessage: err[0],",
		"\t\t}",
		"\t}",
		"\treturn &InternalError{",
		"\t\tMessage: \"Internal server exception\",",
		"\t}",
		"}",
	}

	return strings.Join(lines, "\n")
}

func NewValidationErrorResponseFile() string {
	lines := []string{
		"package responses",
		"",
		"type Violation struct {",
		"\tMessage      string `json:\"message\"`",
		"\tPropertyPath string `json:\"propertyPath\"`",
		"}",
		"",
		"type ValidationError struct {",
		"\tMessage    string      `json:\"message\"`",
		"\tViolations []Violation `json:\"violations\"`",
		"}",
		"",
		"func NewValidationError(message string, violations []Violation) *ValidationError {",
		"\treturn &ValidationError{",
		"\t\tMessage:    message,",
		"\t\tViolations: violations,",
		"\t}",
		"}",
		"",
		"func NewViolation(message string, propertyPath string) *Violation {",
		"\treturn &Violation{",
		"\t\tMessage:      message,",
		"\t\tPropertyPath: propertyPath,",
		"\t}",
		"}",
	}

	return strings.Join(lines, "\n")
}

func NewNotFoundErrorFile() string {
	lines := []string{
		"package responses",
		"",
		"type NotFoundError struct {",
		"\tMessage string `json:\"message\"`",
		"}",
		"",
		"func NewNotFoundError(err ...string) *NotFoundError {",
		"\tif len(err) > 0 {",
		"\t\treturn &NotFoundError{",
		"\t\t\tMessage: err[0],",
		"\t\t}",
		"\t}",
		"\treturn &NotFoundError{",
		"\t\tMessage: \"Not found\",",
		"\t}",
		"}",
	}

	return strings.Join(lines, "\n")
}

func NewBadRequestErrorFile() string {
	lines := []string{
		"package responses",
		"",
		"type BadRequestError struct {",
		"\tMessage string `json:\"message\"`",
		"}",
		"",
		"func NewBadRequestError(err ...string) *BadRequestError {",
		"\tif len(err) > 0 {",
		"\t\treturn &BadRequestError{",
		"\t\t\tMessage: err[0],",
		"\t\t}",
		"\t}",
		"\treturn &BadRequestError{",
		"\t\tMessage: \"Bad Request\",",
		"\t}",
		"}",
	}

	return strings.Join(lines, "\n")
}