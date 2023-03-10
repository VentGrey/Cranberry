package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var root string
	flag.StringVar(&root, "dir", ".", "Directory to examine")
	flag.StringVar(&root, "d", ".", "Directory to examine (shorthand)")
	flag.Parse()

	if flag.NFlag() == 0 {
		help()
		os.Exit(0)
	}

	incidents := 0
	skippedDirs := []string{".git", "vendor", "node_modules", "bower_components",
		"tmp", "log", "logs", "cache", "coverage", "bin", "build", "dist", "out",
		"target", "obj", "docs", "doc", "documentation"}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && contains(skippedDirs, info.Name()) {
			return filepath.SkipDir
		}

		if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") &&
			!strings.HasSuffix(info.Name(), ".d.ts") &&
			strings.HasSuffix(info.Name(), ".ts") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			lines := strings.Split(string(content), "\n")

			for i, line := range lines {
				if strings.Contains(line, "console.log(") ||
					strings.Contains(line, "console.table(") ||
					strings.Contains(line, "console.warn(") ||
					strings.Contains(line, "console.info(") ||
					strings.Contains(line, "console.debug(") &&
						!strings.Contains(line, "console.error") {

					incidentType := ""

					if strings.Contains(line, "console.log(") {
						incidentType = "console.log()"
					} else if strings.Contains(line, "console.table(") {
						incidentType = "console.table()"
					} else if strings.Contains(line, "console.warn(") {
						incidentType = "console.warn()"
					} else if strings.Contains(line, "console.info(") {
						incidentType = "console.info()"
					} else if strings.Contains(line, "console.debug(") {
						incidentType = "console.debug()"
					}

					red := color.New(color.FgRed).SprintFunc()
					yellow := color.New(color.FgYellow).SprintFunc()
					blue := color.New(color.FgBlue).SprintFunc()

					fmt.Printf("%s in %s, line %s: %s\n", red(incidentType), yellow(path), blue(i+1), line)
					incidents += 1
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if incidents > 0 {
		fmt.Print(strings.Repeat("-", 80))
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("\n%s %d %s\n", red("We found"), incidents, red("incidents!"))
		os.Exit(1)
	}

	fmt.Println("No console logging was found!")
}

func help() {
	fmt.Println("Usage: cranberry [options]")
	fmt.Println("Options:")
	fmt.Println("  -d, --dir <path>  Directory to examine")
	fmt.Println("  -h, --help        Display this help message")
}

func contains(s []string, val string) bool {
	for _, item := range s {
		if item == val {
			return true
		}
	}
	return false
}
