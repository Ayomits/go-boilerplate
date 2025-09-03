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

func NewFiberContents(pkg string) []MainFilesContent {
	return []MainFilesContent{
		NewMainFileContent("cmd/app/main.go", NewMainFile(pkg)),
		NewMainFileContent("internal/app/app.go", NewFiberAppFile()),
	}
}

func NewGinContents(pkg string) []MainFilesContent {
	return []MainFilesContent{
		NewMainFileContent("cmd/app/main.go", NewMainFile(pkg)),
		NewMainFileContent("internal/app/app.go", NewGinAppFile()),
	}
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
