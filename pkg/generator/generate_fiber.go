package generator

import (
	"fmt"
	"gb/pkg/logger"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type GenerateFiberProject struct{}

func NewFiberProjectGenerator() ProjectGenerator {
	return &GenerateFiberProject{}
}

func (g *GenerateFiberProject) Generate() {
	var path string
	var name string

	logger.Log.Println("Введите путь вашего проекта: ")
	fmt.Scan(&path)
	logger.Log.Println("Теперь введите название вашего проекта: ")
	fmt.Scan(&name)

	_, err := os.Stat(path)

	if err != nil {
		logger.Log.Fatal("Путь не существует")
		return
	}

	name = strings.Join(strings.Split(name, " "), "")
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)

	{
		RunInPath(path, GoModInit(name))
		Init(path)
	}
}

func RunInPath(path string, cmd *exec.Cmd) {
	cmd.Dir = path
}

func AddPackage(pkg string) *exec.Cmd {
	return exec.Command(fmt.Sprintf("go get -u %s", pkg))
}

func GoModInit(name string) *exec.Cmd {
	return exec.Command(fmt.Sprintf("go mod init %s", name))
}

func GoModVendor() *exec.Cmd {
	return exec.Command("go mod vendor")
}

func Init(path string) {
	var wg sync.WaitGroup
	initers := []func(path string){
		InitPackages,
		InitDirectoryStructure,
		InitMainFiles,
	}

	for _, initer := range initers {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			initer(path)
		}(path)
	}

	wg.Wait()
}

func InitPackages(path string) {
	var wg sync.WaitGroup
	defer GoModVendor()
	packages := []string{
		"github.com/gofiber/fiber/v2",
		"github.com/go-playground/validator/v10",
	}

	for _, pkg := range packages {
		wg.Add(1)
		go func(pkg string) {
			defer wg.Done()
			RunInPath(path, AddPackage(pkg))
		}(pkg)
	}

	wg.Wait()
}

func InitDirectoryStructure(path string) {
	var wg sync.WaitGroup

	directories := []string{
		"cmd/app",
		"internal/app",
		"internal/controllers",
		"internal/models",
		"internal/services",
	}

	for _, dir := range directories {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			RunInPath(path, CreateDir(dir))
		}(dir)
	}

	wg.Wait()
}

func InitMainFiles(path string) {
	var wg sync.WaitGroup
	files := []string{
		"cmd/app/main.go",
		"internal/app/main.go",
	}

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			RunInPath(path, CreateDir(file))
		}(file)
	}

	wg.Wait()
}

func CreateDir(dir string) *exec.Cmd {
	return exec.Command(fmt.Sprintf("mkdir -p \"%s\"", dir))
}

func TouchFile(file string) *exec.Cmd {
	return exec.Command(fmt.Sprintf("mkdir -p \"%s\"", file))
}
