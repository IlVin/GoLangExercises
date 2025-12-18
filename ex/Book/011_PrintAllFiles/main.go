package main

import (
	"fmt"
	"os"
	"strings"
)

func PrintAllFiles(path string) (ok bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		return false
	}
	for _, f := range files {
		file := strings.Join([]string{path, f.Name()}, "")
		fmt.Println("File: ", file)
		if f.IsDir() {
			PrintAllFiles(strings.Join([]string{path, f.Name(), "/"}, ""))
		}
	}

	return true
}

func PrintAllFilesWithFilter(path string, filter string) (ok bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		return false
	}
	for _, f := range files {
		file := strings.Join([]string{path, f.Name()}, "")
		if strings.Contains(f.Name(), filter) {
			fmt.Println("File: ", file)
		}
		if f.IsDir() {
			PrintAllFilesWithFilter(strings.Join([]string{path, f.Name(), "/"}, ""), filter)
		}
	}

	return true
}

func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	var fn func(string)
	fn = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			return
		}
		for _, f := range files {
			file := strings.Join([]string{path, f.Name()}, "")
			if predicate(f.Name()) {
				fmt.Println("File: ", file)
			}
			if f.IsDir() {
				fn(strings.Join([]string{path, f.Name(), "/"}, ""))
			}
		}
	}
	fn(path)
}

func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}

func main() {
	//	PrintAllFilesWithFilter("/home/ilvin/GitHub/IlVin/", "ex")
	//PrintFilesWithFuncFilter("/home/ilvin/GitHub/IlVin/", func(name string) bool { return strings.Contains(name, "_") })

	fmt.Println(unintuitive())
}
