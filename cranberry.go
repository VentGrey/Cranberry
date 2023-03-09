package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func help() {
	fmt.Println("Usage: cranberry [options]")
	fmt.Println("Options:")
	fmt.Println("  -d, --dir <path>  Directory to examine")
	fmt.Println("  -h, --help        Display this help message")
}

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

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") && !strings.HasSuffix(info.Name(), ".d.ts") && strings.HasSuffix(info.Name(), ".ts") {
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
					fmt.Printf("Se encontró una aparición de consola en el archivo %s, línea %d, columna %d\n", path, i+1, strings.Index(line, "console.log")+1)
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
		fmt.Printf("Se encontraron %d apariciones de impresiones de consola\n", incidents)
		os.Exit(1)
	} else {
		fmt.Println("No se encontraron apariciones de impresiones de consola")
		os.Exit(0)
	}
}
