package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func GeneratTemplateSafe(type_ string) {
	var path string
	var name string

	fmt.Println("Введите путь вашего проекта: ")
	fmt.Scan(&path)
	fmt.Println("Теперь введите название вашего проекта: ")
	fmt.Scan(&name)

	relativePath, err := VerifyPath(path)

	if err != nil {
		fmt.Printf("Неудалось создать шаблон проекта: %v", err)
		os.Exit(1)
		return
	}

	err = generateTemplate(type_, *relativePath, name)

	if err != nil {
		fmt.Printf("Не удалось переместить шаблон: %v\n", err)
		os.Exit(1)
		return
	}

	RunInPath(*relativePath, GoModTidy())
	RunInPath(*relativePath, GoModVendor())
}

func generateTemplate(type_ string, path, name string) error {
	name = resolveName(name)

	err := RunInPath(path, GoModInit(name))
	if err != nil {
		return err
	}
	err = InitDirectoryStructure(path)
	if err != nil {
		return err
	}
	err = InitMainFiles(name, path, GetContentsByType(name, type_))
	if err != nil {
		return err
	}
	return nil
}

func VerifyPath(path string) (*string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Ошибка обработки пути: %v", err)
		os.Exit(1)
		return nil, err
	}

	_, err = os.Stat(path)

	if err != nil {
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Ошибка, не удалось создать директории")
			os.Exit(1)
			return nil, err
		}
	}

	return &path, nil
}

func InitDirectoryStructure(path string) error {
	var wg sync.WaitGroup

	directories := []string{
		"cmd/app",
		"internal/app",
		"internal/controllers",
		"internal/models",
		"internal/services",
		"internal/repositories",
		"internal/services/validators",
		"internal/dtos",
		"internal/responses",
	}

	errors_ := []error{}

	for _, dir := range directories {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			err := RunInPath(path, CreateDir(dir))
			errors_ = append(errors_, err)
		}(dir)
	}

	if len(errors_) > 0 {
		return errors.New("something went wrong")
	}

	wg.Wait()
	return nil
}

func InitMainFiles(pkgName string, path string, content []MainFilesContent) error {
	var wg sync.WaitGroup

	errors_ := []error{}

	for _, c := range content {
		wg.Add(1)
		go func(c MainFilesContent) {
			fullPath := filepath.Join(path, c.Filename)

			if err := os.WriteFile(fullPath, []byte(c.Content), 0655); err != nil {
				fmt.Printf("Ошибка во время создания файла %s: \n %v", fullPath, err)
				errors_ = append(errors_, err)
				return
			}

			wg.Done()
		}(c)
	}

	if len(errors_) > 0 {
		return errors.New("something went wrong")
	}

	wg.Wait()
	return nil
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
	common := []string{
		"github.com/go-playground/validator/v10",
	}
	switch type_ {
	case GinType:
		gin := []string{
			"github.com/gin-gonic/gin",
		}
		common = append(common, gin...)
	case FiberType:
		fiber := []string{
			"github.com/gofiber/fiber/v2",
		}
		common = append(common, fiber...)
	}

	return common
}

func resolveName(name string) string {
	name = strings.Join(strings.Split(name, " "), "")
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	return name
}
