package utils

import (
	"fmt"
	"strings"
)

func NewFiberAppFile() string {
	lines := []string{
		"package app",
		"",
		`import "github.com/gofiber/fiber/v2"`,
		"",
		"func RunApp() {",
		"\tapp := fiber.New()",
		"",
		"\t// define your routes here",
		"",
		"\tapp.Listen(\":8080\")",
		"}",
		"",
	}

	return strings.Join(lines, "\n")
}

func NewFiberValidatorFile(pkg string) string {
	lines := []string{
		"package validators",
		"",
		"import (",
		fmt.Sprintf("\t\"%s/internal/responses\"", pkg),
		"\t\"reflect\"",
		"\t\"strings\"",
		"\t\"github.com/go-playground/validator/v10\"",
		")",
		"",
		"type AppValidator interface {",
		"\tValidate(body any) *responses.ValidationError",
		"}",
		"",
		"type validationFormatter struct {",
		"\tvalidate *validator.Validate",
		"}",
		"",
		"func NewAppValidator() AppValidator {",
		"\treturn &validationFormatter{",
		"\t\tvalidate: validator.New(),",
		"\t}",
		"}",
		"",
		"func (f *validationFormatter) Validate(body any) *responses.ValidationError {",
		"\terr := f.validate.Struct(body)",
		"\tvar violations []responses.Violation",
		"",
		"\tif err == nil {",
		"\t\treturn nil",
		"\t}",
		"",
		"\tfor _, violation := range err.(validator.ValidationErrors) {",
		"\t\tviolations = append(violations, *responses.NewViolation(f.formatErrorMessage(violation), f.getJSONFieldName(violation, body)))",
		"\t}",
		"",
		"\treturn responses.NewValidationError(\"Invalid payload\", violations)",
		"}",
		"",
		"func (f *validationFormatter) getJSONFieldName(fieldErr validator.FieldError, body any) string {",
		"\tt := reflect.TypeOf(body)",
		"",
		"\tif t.Kind() == reflect.Ptr {",
		"\t\tt = t.Elem()",
		"\t}",
		"",
		"\tfield, found := t.FieldByName(fieldErr.Field())",
		"\tif !found {",
		"\t\treturn fieldErr.Field()",
		"\t}",
		"",
		"\tjsonTag := field.Tag.Get(\"json\")",
		"\tif jsonTag == \"\" {",
		"\t\treturn fieldErr.Field()",
		"\t}",
		"",
		"\treturn strings.Split(jsonTag, \",\")[0]",
		"}",
		"",
		"func (f *validationFormatter) formatErrorMessage(fieldErr validator.FieldError) string {",
		"\treturn fieldErr.Error()",
		"}",
		"",
		"var AppValidatorInstance AppValidator = NewAppValidator()",
	}

	return strings.Join(lines, "\n")
}
