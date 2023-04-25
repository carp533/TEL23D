package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Datei konnte nicht gelesen werden")
		return
	}

	_, err = f.Write([]byte("Test"))
	if err != nil {
		fmt.Println("Datei konnte nicht geschrieben werden")
		return
	}

	defer f.Close()
	subdirs, err := getSubdirs(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(subdirs)

}

func getSubdirs(dirPath string) ([]string, error) {
	var subdirs []string
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subdirs = append(subdirs, file.Name())
		}
	}

	return subdirs, nil
}
