package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func RunInPath(path string, cmd *exec.Cmd) error {
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func GoModInit(name string) *exec.Cmd {
	return exec.Command("go", "mod", "init", name)
}

func AddPackage(pkg string) *exec.Cmd {
	return exec.Command("go", "get", "-u", pkg)
}

func GoModVendor() *exec.Cmd {
	return exec.Command("go", "mod", "vendor")
}

func GoModTidy() *exec.Cmd {
	return exec.Command("go", "mod", "tidy")
}

func CreateDir(dir string) *exec.Cmd {
	return exec.Command("mkdir", "-p", dir)
}

func TouchFile(file string) *exec.Cmd {
	return exec.Command("touch", file)
}

func WriteFile(path string, input string) {
	if err := os.WriteFile(path, []byte(input), 0644); err != nil {
		fmt.Printf("Не удалось записать файл %v", err)
		os.Exit(1)
	}
}