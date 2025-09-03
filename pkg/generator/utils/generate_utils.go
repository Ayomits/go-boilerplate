package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func GenerateTemplate(type_ string) {
	var path string
	var name string

	fmt.Println("Введите путь вашего проекта: ")
	fmt.Scan(&path)
	fmt.Println("Теперь введите название вашего проекта: ")
	fmt.Scan(&name)

	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Ошибка обработки пути: %v", err)
		os.Exit(1)
		return
	}

	_, err = os.Stat(path)

	if err != nil {
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Ошибка, не удалось создать директории")
			os.Exit(1)
			return
		}
	}

	name = resolveName(name)

	RunInPath(path, GoModInit(name))
	Init(name, path, type_)
}

func Init(pkg string, path string, type_ string) {
	InitDirectoryStructure(path)
	InitMainFiles(pkg, path, GetContentsByType(pkg, type_))
	InitPackages(path, type_)
}

func InitPackages(path string, type_ string) {
	var wg sync.WaitGroup

	packages := GetPackagesListByType(type_)

	for _, pkg := range packages {
		wg.Add(1)
		go func(pkg string) {
			defer wg.Done()
			RunInPath(path, AddPackage(pkg))
		}(pkg)
	}

	wg.Wait()

	RunInPath(path, GoModVendor())
	RunInPath(path, GoModTidy())
}

func GetPackagesListByType(type_ string) []string {
	switch type_ {
	case GinType:
		return []string{
			"github.com/gin-gonic/gin",
		}
	case FiberType:
		return []string{
			"github.com/gofiber/fiber/v2",
			"github.com/go-playground/validator/v10",
		}
	}

	return []string{}
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

func InitMainFiles(pkgName string, path string, content []MainFilesContent) {
	var wg sync.WaitGroup

	for _, c := range content {
		wg.Add(1)
		go func(c MainFilesContent) {
			fullPath := filepath.Join(path, c.Filename)

			if err := os.WriteFile(fullPath, []byte(c.Content), 0655); err != nil {
				fmt.Printf("Ошибка во время создания файла %s: \n %v", fullPath, err)
				os.Exit(1)
			}

			wg.Done()
		}(c)
	}

	wg.Wait()
}

func resolveName(name string) string {
	name = strings.Join(strings.Split(name, " "), "")
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	return name
}
