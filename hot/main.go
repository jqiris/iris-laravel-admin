package main

import (
	"github.com/kataras/rizla/rizla"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error loading project rootPath")
	}
	project := rizla.NewProject(rootPath + "/../server.go")
	project.Watcher = func(absolutePath string) bool {
		base := filepath.Base(absolutePath)
		return !(base == ".git" ||
			base == ".idea" ||
			base == "public" ||
			base == "resources" ||
			base == "storage" ||
			base == "tests")
	}

	project.Matcher = func(filename string) bool {
		isWindows := runtime.GOOS == "windows"
		goExt := ".go"
		envExt := ".env"
		return (filepath.Ext(filename) == goExt || filepath.Ext(filename) == envExt) ||
			(!isWindows && strings.Contains(filename, goExt)) ||
			(!isWindows && strings.Contains(filename, envExt))
	}
	rizla.Add(project)
	rizla.Run(nil)
}
