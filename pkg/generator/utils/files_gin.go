package utils

import (
	"fmt"
	"strings"
)

func NewGinAppFile() string {
	lines := []string{
		"package app",
		"",
		`import "github.com/gin-gonic/gin"`,
		"",
		"func Run() {",
		"\tapp := gin.Default()",
		"\t",
		"\t// define your routes here",
		"",
		"\tapp.Run()",
		"}",
		"",
	}

	return strings.Join(lines, "\n")
}

func NewGinValidatorFile(pkg string) string {
	lines := []string{
		"package validators",
		"",
		"import (",
		fmt.Sprintf("\t\"%s/pkg/responses\"", pkg),
		"\t\"reflect\"",
		"\t\"strings\"",
		"\t\"github.com/go-playground/validator/v10\"",
		")",
		"",
		"type AppValidator struct {",
		"\tvalidate *validator.Validate `validate:\"required\"`",
		"}",
		"",
		"func NewAppValidator() AppValidator {",
		"\treturn AppValidator{",
		"\t\tvalidate: validator.New(),",
		"\t}",
		"}",
		"",
		"func (f *AppValidator) Validate(err error, body any) *responses.ValidationError {",
		"\tvar violations []responses.Violation",
		"",
		"\tif err == nil {",
		"\t\treturn nil",
		"\t}",
		"",
		"\tfor _, violation := range err.(validator.ValidationErrors) {",
		"\t\tviolations = append(violations, responses.NewViolation(f.getJSONFieldName(violation, body), f.formatErrorMessage(violation)))",
		"\t}",
		"",
		"\treturn responses.NewValidationError(violations, nil)",
		"}",
		"",
		"func (f *AppValidator) getJSONFieldName(fieldErr validator.FieldError, body any) string {",
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
		"func (f *AppValidator) formatErrorMessage(fieldErr validator.FieldError) string {",
		"\treturn fieldErr.Error()",
		"}",
		"",
		"var AppValidatorInstance AppValidator = NewAppValidator()",
	}

	return strings.Join(lines, "\n")
}
