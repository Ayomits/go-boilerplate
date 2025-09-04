package utils

const (
	GinType   = "gin"
	FiberType = "fiber"
)

type MainFilesContent struct {
	Filename string
	Content  string
}

func NewMainFileContent(f string, c string) MainFilesContent {
	return MainFilesContent{
		Filename: f,
		Content:  c,
	}
}

func getCommonFilesContent(pkg string) []MainFilesContent {
	return []MainFilesContent{
		NewMainFileContent("cmd/app/main.go", NewMainFile(pkg)),
		NewMainFileContent("internal/responses/internal_error_response.go", NewInternalResponseFile()),
		NewMainFileContent("internal/responses/validation_error_response.go", NewValidationErrorResponseFile()),
		NewMainFileContent("internal/responses/not_found_error_response.go", NewNotFoundErrorFile()),
		NewMainFileContent("internal/responses/bad_request_error_response.go", NewBadRequestErrorFile()),
	}
}

func NewFiberContents(pkg string) []MainFilesContent {
	contents := []MainFilesContent{
		NewMainFileContent("internal/app/app.go", NewFiberAppFile()),
		NewMainFileContent("internal/services/validators/validator.go", NewFiberValidatorFile(pkg)),
	}

	contents = append(contents, getCommonFilesContent(pkg)...)

	return contents
}

func NewGinContents(pkg string) []MainFilesContent {
	contents := []MainFilesContent{
		NewMainFileContent("internal/app/app.go", NewGinAppFile()),
		NewMainFileContent("internal/services/validators/validator.go", NewGinValidatorFile(pkg)),
	}

	contents = append(contents, getCommonFilesContent(pkg)...)

	return contents
}

func GetContentsByType(pkg string, type_ string) []MainFilesContent {
	switch type_ {
	case GinType:
		return NewGinContents(pkg)
	case FiberType:
		return NewFiberContents(pkg)
	}
	return []MainFilesContent{}
}
