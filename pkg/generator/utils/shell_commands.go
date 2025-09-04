package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func RemoveFile(path string) error {
	return os.RemoveAll(path)
}

func MoveFiles(from, to string) error {
	if err := os.MkdirAll(to, 0755); err != nil {
        return err
    }

	entries, err := os.ReadDir(from)
    if err != nil {
        return err
    }
    for _, entry := range entries {
        srcPath := filepath.Join(from, entry.Name())
        dstPath := filepath.Join(to, entry.Name())
        if err := os.Rename(srcPath, dstPath); err != nil {
            return err
        }
    }
    return nil
}

func WriteFile(path string, input string) {
	if err := os.WriteFile(path, []byte(input), 0644); err != nil {
		fmt.Printf("Не удалось записать файл %v", err)
		os.Exit(1)
	}
}
