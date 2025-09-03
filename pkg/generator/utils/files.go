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

func NewGinAppFile() string {
	lines := []string{
		"package app",
		"",
		`import "github.com/gin-gonic/gin"`,
		"",
		"func RunApp() {",
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

func NewMainFile(name string) string {
	lines := []string{
		"package main",
		"",
		fmt.Sprintf(`import "%s/internal/app"`, name),
		"",
		"func main() {",
		"\tapp.RunApp()",
		"}",
	}

	return strings.Join(lines, "\n")
}
